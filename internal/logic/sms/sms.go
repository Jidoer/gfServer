package sms

import (
	"context"
	"encoding/json"
	"fmt"

	// "io"

	// "encoding/json"
	"math/rand"
	"strconv"
	"time"

	// "fmt"
	"gfAdmin/internal/cache"
	"gfAdmin/internal/service"

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

type SmsData struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

func RandomSmsCode() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(9000) + 1000) // 生成 1000~9999 的随机整数
}

type AliSmsResult struct {
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	BizID     string `json:"BizId"`
	RequestID string `json:"RequestId"`
}

func (s *sSms) SendaSms(ctx context.Context, phone string) (string, error) {
	code := "0000"//RandomSmsCode()
	// r, err := SendAliSms(phone, code)
	// if err != nil {
	// 	return "", err
	// }
	r := "{\"Code\":\"OK\",\"Message\":\"OK\",\"BizId\":\"123123123\",\"RequestId\":\"123\"}"
	var aliSmsResult AliSmsResult
	err := json.Unmarshal([]byte(r), &aliSmsResult)
	if err != nil {
		return "", err
	}
	if aliSmsResult.Code != "OK" || aliSmsResult.Message != "OK" {
		return "", gerror.New(aliSmsResult.Message)
	}
	smsData := SmsData{
		Code:  code,
		Phone: phone,
	}
	err = cache.Instance().Set(ctx, "sms_"+aliSmsResult.BizID, smsData, time.Minute*5)
	if err != nil {
		return "", err
	}

	return aliSmsResult.BizID, nil
}

// //写入redis
// smsData := SmsData{
// 	Code:  code,
// 	Phone: phone,
// }
// err := cache.Instance().Set(ctx, "sms_"+data.(map[string]interface{})["sid"].(string), smsData, time.Minute*5)
// if err != nil {
// 	return "", err
// }
// return data.(map[string]interface{})["sid"].(string), nil

// smsData := SmsData{
// 	Code:  "0000",
// 	Phone: phone,
// }
// err := cache.Instance().Set(ctx, "sms_"+"sid_"+phone, smsData, time.Minute*5)
// if err != nil {
// 	return "", err
// }
// return "sid_" + phone, nil

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
		fmt.Println("验证码错误 --> SCODE:", scode.Code,"CODE:", code)
		return "", gerror.New("验证码错误")
	}
	return scode.Phone, nil
}
