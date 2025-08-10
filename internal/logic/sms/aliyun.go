package sms

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"

	// "os"
	"sort"

	"golang.org/x/exp/maps"

	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	httpMethod   string
	canonicalUri string
	host         string
	xAcsAction   string
	xAcsVersion  string
	headers      map[string]string
	body         []byte
	queryParam   map[string]interface{}
}

func NewRequest(httpMethod, canonicalUri, host, xAcsAction, xAcsVersion string) *Request {
	req := &Request{
		httpMethod:   httpMethod,
		canonicalUri: canonicalUri,
		host:         host,
		xAcsAction:   xAcsAction,
		xAcsVersion:  xAcsVersion,
		headers:      make(map[string]string),
		queryParam:   make(map[string]interface{}),
	}
	req.headers["host"] = host
	req.headers["x-acs-action"] = xAcsAction
	req.headers["x-acs-version"] = xAcsVersion
	req.headers["x-acs-date"] = time.Now().UTC().Format(time.RFC3339)
	req.headers["x-acs-signature-nonce"] = uuid.New().String()
	return req
}

var (
	AccessKeyId     = "LTAI5tD1sGHkvmr3mJWArp3G"       //os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	AccessKeySecret = "4aFfOP0pqGa6l4m9xganblHEDNbSb8" //os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	ALGORITHM       = "ACS3-HMAC-SHA256"
)

func SendAliSms(phone, code string) (string, error) {
	httpMethod := "GET"             // 请求方式，大部分RPC接口同时支持POST和GET，此处以POST为例
	canonicalUri := "/"             // RPC接口无资源路径，故使用正斜杠（/）作为CanonicalURI
	host := "dysmsapi.aliyuncs.com" // 云产品服务接入点
	xAcsAction := "SendSms"         // API名称
	xAcsVersion := "2017-05-25"     // API版本号
	req := NewRequest(httpMethod, canonicalUri, host, xAcsAction, xAcsVersion)
	req.queryParam["PhoneNumbers"] = "0086" + phone
	req.queryParam["SignName"] = "永嘉县艾润科技"
	req.queryParam["TemplateCode"] = "SMS_324515311"
	req.queryParam["TemplateParam"] = "{\"code\":\"" + code + "\"}"
	// 签名过程
	getAuthorization(req)
	// 调用API
	r, error := callAPI(req)
	if error != nil {
		println(error.Error())
		return "", error
	}
	return r, error
}

func callAPI(req *Request) (string, error) {
	urlStr := "https://" + req.host + req.canonicalUri
	q := url.Values{}
	keys := maps.Keys(req.queryParam)
	sort.Strings(keys)
	for _, k := range keys {
		v := req.queryParam[k]
		q.Set(k, fmt.Sprintf("%v", v))
	}
	urlStr += "?" + q.Encode()
	fmt.Println(urlStr)

	httpReq, err := http.NewRequest(req.httpMethod, urlStr, strings.NewReader(string(req.body)))
	if err != nil {
		return "", err
	}

	for key, value := range req.headers {
		httpReq.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	var respBuffer bytes.Buffer
	_, err = io.Copy(&respBuffer, resp.Body)
	if err != nil {
		return "", err
	}
	respBytes := respBuffer.Bytes()
	// fmt.Println(string(respBytes))

	return string(respBytes), nil
}

func getAuthorization(req *Request) {
	// 处理queryParam中参数值为List、Map类型的参数，将参数平铺
	newQueryParams := make(map[string]interface{})
	processObject(newQueryParams, "", req.queryParam)
	req.queryParam = newQueryParams
	// 步骤 1：拼接规范请求串
	canonicalQueryString := ""
	keys := maps.Keys(req.queryParam)
	sort.Strings(keys)
	for _, k := range keys {
		v := req.queryParam[k]
		canonicalQueryString += percentCode(url.QueryEscape(k)) + "=" + percentCode(url.QueryEscape(fmt.Sprintf("%v", v))) + "&"
	}
	canonicalQueryString = strings.TrimSuffix(canonicalQueryString, "&")
	fmt.Printf("canonicalQueryString========>%s\n", canonicalQueryString)

	var bodyContent []byte
	if req.body == nil {
		bodyContent = []byte("")
	} else {
		bodyContent = req.body
	}
	hashedRequestPayload := sha256Hex(bodyContent)
	req.headers["x-acs-content-sha256"] = hashedRequestPayload

	canonicalHeaders := ""
	signedHeaders := ""
	HeadersKeys := maps.Keys(req.headers)
	sort.Strings(HeadersKeys)
	for _, k := range HeadersKeys {
		lowerKey := strings.ToLower(k)
		if lowerKey == "host" || strings.HasPrefix(lowerKey, "x-acs-") || lowerKey == "content-type" {
			canonicalHeaders += lowerKey + ":" + req.headers[k] + "\n"
			signedHeaders += lowerKey + ";"
		}
	}
	signedHeaders = strings.TrimSuffix(signedHeaders, ";")

	canonicalRequest := req.httpMethod + "\n" + req.canonicalUri + "\n" + canonicalQueryString + "\n" + canonicalHeaders + "\n" + signedHeaders + "\n" + hashedRequestPayload
	fmt.Printf("canonicalRequest========>\n%s\n", canonicalRequest)

	// 步骤 2：拼接待签名字符串
	hashedCanonicalRequest := sha256Hex([]byte(canonicalRequest))
	stringToSign := ALGORITHM + "\n" + hashedCanonicalRequest
	fmt.Printf("stringToSign========>\n%s\n", stringToSign)

	// 步骤 3：计算签名
	byteData, err := hmac256([]byte(AccessKeySecret), stringToSign)
	if err != nil {
		fmt.Println(err)
	}
	signature := strings.ToLower(hex.EncodeToString(byteData))

	// 步骤 4：拼接Authorization
	authorization := ALGORITHM + " Credential=" + AccessKeyId + ",SignedHeaders=" + signedHeaders + ",Signature=" + signature
	fmt.Printf("authorization========>%s\n", authorization)
	req.headers["Authorization"] = authorization
}

func hmac256(key []byte, toSignString string) ([]byte, error) {
	// 实例化HMAC-SHA256哈希
	h := hmac.New(sha256.New, key)
	// 写入待签名的字符串
	_, err := h.Write([]byte(toSignString))
	if err != nil {
		return nil, err
	}
	// 计算签名并返回
	return h.Sum(nil), nil
}

func sha256Hex(byteArray []byte) string {
	// 实例化SHA-256哈希函数
	hash := sha256.New()
	// 将字符串写入哈希函数
	_, _ = hash.Write(byteArray)
	// 计算SHA-256哈希值并转换为小写的十六进制字符串
	hexString := hex.EncodeToString(hash.Sum(nil))

	return hexString
}

func percentCode(str string) string {
	// 替换特定的编码字符
	str = strings.ReplaceAll(str, "+", "%20")
	str = strings.ReplaceAll(str, "*", "%2A")
	str = strings.ReplaceAll(str, "%7E", "~")
	return str
}

func formDataToString(formData map[string]interface{}) *string {
	tmp := make(map[string]interface{})
	processObject(tmp, "", formData)
	res := ""
	urlEncoder := url.Values{}
	for key, value := range tmp {
		v := fmt.Sprintf("%v", value)
		urlEncoder.Add(key, v)
	}
	res = urlEncoder.Encode()
	return &res
}

// processObject 递归处理对象，将复杂对象（如Map和List）展开为平面的键值对
func processObject(mapResult map[string]interface{}, key string, value interface{}) {
	if value == nil {
		return
	}

	switch v := value.(type) {
	case []interface{}:
		for i, item := range v {
			processObject(mapResult, fmt.Sprintf("%s.%d", key, i+1), item)
		}
	case map[string]interface{}:
		for subKey, subValue := range v {
			processObject(mapResult, fmt.Sprintf("%s.%s", key, subKey), subValue)
		}
	default:
		if strings.HasPrefix(key, ".") {
			key = key[1:]
		}
		if b, ok := v.([]byte); ok {
			mapResult[key] = string(b)
		} else {
			mapResult[key] = fmt.Sprintf("%v", v)
		}
	}
}
