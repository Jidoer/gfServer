package sms

import (
	"context"
	"encoding/json"
	// "encoding/json"
	"math/rand"
	"strconv"
	"time"

	// "fmt"
	"gfAdmin/internal/cache"
	"gfAdmin/internal/service"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sSms struct{}
)

func init() {
	service.RegisterSms(New())
}

var logger = g.Log("sms service")

func New() service.ISms {
	return &sSms{}
}

const apiUrl = "http://v.juhe.cn/sms/send"

type SmsData struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

func RandomSmsCode() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(9000) + 1000) // 生成 1000~9999 的随机整数
}

func (s *sSms) SendaSms(ctx context.Context, phone string) (string, error) {
	code:=RandomSmsCode()
	param := url.Values{}
	param.Set("mobile", phone)             // 接收短信的手机号码
	param.Set("tpl_id", "273084")               // 短信模板ID，请参考个人中心短信模板设置
	param.Set("tpl_value", "#code#=" +code + " #m#=5") // 模板变量，如无则不用填写
	param.Set("key", "c63c06f870950250eda9dfd85d229e60")
	data, err := Post(apiUrl, param)
	if err != nil {
		logger.Errorf("请求异常:\r\n%v", err)
		return "", gerror.New("请求异常")
	} else {
		var netReturn map[string]interface{}
		jsonerr := json.Unmarshal(data, &netReturn)
		if jsonerr != nil {
			logger.Errorf("请求异常:%v", jsonerr)
			return "", gerror.New("请求异常")
		} else {
			errorCode := netReturn["error_code"]
			reason := netReturn["reason"]
			data := netReturn["result"]
			if errorCode.(float64) == 0 {
				logger.Printf("请求成功\n短信ID：%v",
					data.(map[string]interface{})["sid"])
				// sid := data.(map[string]interface{})["sid"]
				//写入redis
				smsData := SmsData{
					Code: code,
					Phone: phone,
				}
				err := cache.Instance().Set(ctx, "sms_"+data.(map[string]interface{})["sid"].(string), smsData, time.Minute*5)
				if err != nil {
					return "", err
				}
				return data.(map[string]interface{})["sid"].(string), nil
			} else {
				logger.Printf("请求失败:%v_%v", errorCode.(float64), reason)
				return "", gerror.New("请求失败")
			}
		}
	}
	// smsData := SmsData{
	// 	Code:  "0000",
	// 	Phone: phone,
	// }
	// err := cache.Instance().Set(ctx, "sms_"+"sid_"+phone, smsData, time.Minute*5)
	// if err != nil {
	// 	return "", err
	// }
	// return "sid_" + phone, nil
}

func (s *sSms) VerifySmsCode(ctx context.Context, sms_id string, code string) (string, error) {
	var scode SmsData
	r, err := cache.Instance().Get(ctx, "sms_"+sms_id)
	if err != nil {
		return "", err
	}
	if err := r.Struct(&scode); err != nil {
		return "", err
	}
	if scode.Code != code {
		return "", gerror.New("验证码错误")
	}
	return scode.Phone, nil
}

// post 方式发起网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
