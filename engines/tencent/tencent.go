package tencent

import (
	"github.com/falcolee/transgo/engines"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

const (
	region = "ap-shanghai"
)

func init() {
	engines.Register([]string{"tencent", "qq"}, New())
}

type engine struct {
	tmtClient *tmt.Client
}

// New returns a new tencent engine
func New() engines.Engine {
	return &engine{tmtClient: nil}
}

func (e *engine) Translate(text string, options engines.Options) (string, error) {
	if e.tmtClient == nil {
		selectRegion := region
		if options.Region != "" {
			selectRegion = options.Region
		}
		appId := options.AppId
		appSecret := options.AppSecret
		if options.TLConfig != nil && appId == "" && appSecret == "" {
			appId = options.TLConfig.Tencent.AppId
			appSecret = options.TLConfig.Tencent.AppSecret
		}
		credential := common.NewCredential(appId, appSecret)
		client, err := tmt.NewClient(credential, selectRegion, profile.NewClientProfile())
		if err != nil {
			return "", errors.WithStack(err)
		}
		e.tmtClient = client
	}

	request := tmt.NewTextTranslateRequest()
	desLanguage := e.LanguageCode(options.GetToLanguage())
	srcLanguage := e.LanguageCode(options.GetFromLanguage())
	request.Source = &srcLanguage
	request.SourceText = &text
	request.Target = &desLanguage
	id := int64(0)
	request.ProjectId = &id
	transResponse, err := e.tmtClient.TextTranslate(request)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return *transResponse.Response.TargetText, nil
}

func (e *engine) LanguageCode(code string) string {
	return code
}
