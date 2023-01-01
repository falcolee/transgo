package volcengine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/falcolee/transgo/engines"
	"github.com/pkg/errors"
	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	host            = "open.volcengineapi.com"
	kServiceVersion = "2020-06-01"
)

func init() {
	engines.Register([]string{"volcengine", "huoshan"}, New())
}

type engine struct {
	client *base.Client
}

// New returns a new volc engine
func New() engines.Engine {
	return &engine{}
}

func (e *engine) Translate(text string, options engines.Options) (string, error) {
	if e.client == nil {
		serviceInfo := &base.ServiceInfo{
			Timeout: 5 * time.Second,
			Host:    host,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "translate"},
		}
		apiInfoList := map[string]*base.ApiInfo{
			"TranslateText": {
				Method: http.MethodPost,
				Path:   "/",
				Query: url.Values{
					"Action":  []string{"TranslateText"},
					"Version": []string{kServiceVersion},
				},
			},
		}
		appId := options.AppId
		appSecret := options.AppSecret
		if options.TLConfig != nil && appId == "" && appSecret == "" {
			appId = options.TLConfig.Volcengine.AppId
			appSecret = options.TLConfig.Volcengine.AppSecret
		}
		client := base.NewClient(serviceInfo, apiInfoList)
		client.SetAccessKey(appId)
		client.SetSecretKey(appSecret)
		e.client = client
	}
	sourceLanguage := e.LanguageCode(options.GetFromLanguage())
	targetLanguage := e.LanguageCode(options.GetToLanguage())
	requestJson := fmt.Sprintf(`{"SourceLanguage":"%s","TargetLanguage":"%s","TextList":["%s"]}`, sourceLanguage, targetLanguage, text)
	resp, code, err := e.client.Json("TranslateText", nil, requestJson)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if code != 200 {
		return "", fmt.Errorf("火山翻译错误，返回错误码:%d", code)
	}
	translationData := translationData{}
	err = json.Unmarshal(resp, &translationData)
	if err != nil {
		return "", err
	}
	if translationData.ResponseMetadata.Error.Code != "" {
		return "", fmt.Errorf("火山翻译错误，返回错误码:%s，错误原因:%s", translationData.ResponseMetadata.Error.Code, translationData.ResponseMetadata.Error.Message)
	}
	if len(translationData.TranslationList) == 0 {
		return "", fmt.Errorf("火山翻译错误，返回内容为空")
	}
	return translationData.TranslationList[0].Translation, nil
}

func (e *engine) LanguageCode(code string) string {
	switch code {
	case "auto":
		return ""
	case "zh-TW":
		return "zh-Hant"
	}
	return code
}
