package http_client

import (
    "bytes"
    "encoding/json"
    "io"
    // "mime/multipart"
    "net/http"
    "net/url"
    "strings"
    "time"
	"github.com/gogf/gf/frame/g"
)

var logger = g.Log("HttpClient")


// HttpClient 是一个封装的 HTTP 客户端
type HttpClient struct {
    BaseURL string
    Headers map[string]string
    Timeout time.Duration
    Client  *http.Client
}

// NewHttpClient 构造函数
func NewHttpClient(baseURL string, timeout time.Duration) *HttpClient {
    return &HttpClient{
        BaseURL: baseURL,
        Timeout: timeout,
        Headers: make(map[string]string),
        Client: &http.Client{
            Timeout: timeout,
        },
    }
}

// RequestOption 可选参数
type RequestOption struct {
    QueryParams map[string]string
    FormData    map[string]string
    JsonData    interface{}
    Headers     map[string]string
}

// doRequest 通用请求方法
func (hc *HttpClient) doRequest(method, path string, opt *RequestOption) (*http.Response, error) {
    fullURL := hc.BaseURL + path
    logger.Println("doRequest Full URL",fullURL)

    // 添加 query 参数
    if opt != nil && opt.QueryParams != nil {
        u, err := url.Parse(fullURL)
        if err != nil {
            return nil, err
        }
        q := u.Query()
        for k, v := range opt.QueryParams {
            q.Set(k, v)
        }
        u.RawQuery = q.Encode()
        fullURL = u.String()
    }

    var body io.Reader

    if opt != nil {
        if opt.JsonData != nil {
            // JSON 格式
            jsonBytes, err := json.Marshal(opt.JsonData)
            if err != nil {
                return nil, err
            }
            body = bytes.NewBuffer(jsonBytes)
        } else if opt.FormData != nil {
            // 表单格式
            form := url.Values{}
            for k, v := range opt.FormData {
                form.Set(k, v)
            }
            body = strings.NewReader(form.Encode())
        }
    }

    req, err := http.NewRequest(method, fullURL, body)
    if err != nil {
        return nil, err
    }

    // 设置默认 Header
    for k, v := range hc.Headers {
        req.Header.Set(k, v)
    }

    // 设置请求自定义 Header
    if opt != nil && opt.Headers != nil {
        for k, v := range opt.Headers {
            req.Header.Set(k, v)
        }
    }

    // 自动设置 Content-Type
    if opt != nil {
        if opt.JsonData != nil {
            req.Header.Set("Content-Type", "application/json")
        } else if opt.FormData != nil {
            req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
        }
    }

    return hc.Client.Do(req)
}

// GET 请求
func (hc *HttpClient) Get(path string, opt *RequestOption) (*http.Response, error) {
    return hc.doRequest("GET", path, opt)
}

// POST 请求
func (hc *HttpClient) Post(path string, opt *RequestOption) (*http.Response, error) {
    return hc.doRequest("POST", path, opt)
}

// PUT 请求
func (hc *HttpClient) Put(path string, opt *RequestOption) (*http.Response, error) {
    return hc.doRequest("PUT", path, opt)
}

// DELETE 请求
func (hc *HttpClient) Delete(path string, opt *RequestOption) (*http.Response, error) {
    return hc.doRequest("DELETE", path, opt)
}
