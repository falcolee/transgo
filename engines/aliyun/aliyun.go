package aliyun

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
	"github.com/falcolee/transgo/engines"
	"github.com/pkg/errors"
)

const (
	region = "cn-hangzhou"
)

func init() {
	engines.Register([]string{"aliyun"}, New())
}

type engine struct {
	alimtClient *alimt.Client
}

// New returns a new baidu engine
func New() engines.Engine {
	return &engine{alimtClient: nil}
}

func (e *engine) Translate(text string, options engines.Options) (string, error) {
	if e.alimtClient == nil {
		selectRegion := region
		if options.Region != "" {
			selectRegion = options.Region
		}
		appId := options.AppId
		appSecret := options.AppSecret
		if options.TLConfig != nil && appId == "" && appSecret == "" {
			appId = options.TLConfig.Aliyun.AppId
			appSecret = options.TLConfig.Aliyun.AppSecret
		}
		// 创建ecsClient实例
		alimtClient, err := alimt.NewClientWithAccessKey(
			selectRegion, // 地域ID
			appId,        // 您的Access Key ID
			appSecret)    // 您的Access Key Secret
		if err != nil {
			return "", errors.WithStack(err)
		}
		e.alimtClient = alimtClient
	}
	request := alimt.CreateTranslateECommerceRequest()
	// 等价于 request.PageSize = "10"
	request.Method = "POST"                                            //设置请求
	request.FormatType = "text"                                        //翻译文本的格式
	request.SourceLanguage = e.LanguageCode(options.GetFromLanguage()) //源语言
	request.SourceText = text                                          //原文
	request.TargetLanguage = e.LanguageCode(options.GetToLanguage())   //目标语言
	request.Scene = "general"                                          //通用版本
	// 发起请求并处理异常
	response, err := e.alimtClient.TranslateECommerce(request)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if response.Code != 200 {
		return "", fmt.Errorf("aliyun翻译错误，返回错误码:%d", response.Code)
	}
	return response.Data.Translated, nil
}

func (e *engine) LanguageCode(code string) string {
	return code
}
