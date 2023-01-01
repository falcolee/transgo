package baidu

import (
	"fmt"

	"github.com/chyroc/baidufanyi"
	"github.com/falcolee/transgo/engines"
	"github.com/pkg/errors"
)

func init() {
	engines.Register([]string{"baidu"}, New())
}

type engine struct{}

// New returns a new baidu engine
func New() engines.Engine {
	return &engine{}
}

func (e *engine) Translate(text string, options engines.Options) (string, error) {
	appId := options.AppId
	appSecret := options.AppSecret
	if options.TLConfig != nil && appId == "" && appSecret == "" {
		appId = options.TLConfig.Baidu.AppId
		appSecret = options.TLConfig.Baidu.AppSecret
	}
	cli := baidufanyi.New(baidufanyi.WithCredential(appId, appSecret))
	res, err := cli.Translate(text, baidufanyi.Language(e.LanguageCode(options.GetFromLanguage())), baidufanyi.Language(e.LanguageCode(options.GetToLanguage())))
	if err != nil {
		return "", errors.WithStack(err)
	}
	if len(res) == 0 {
		return "", fmt.Errorf("百度翻译错误，返回内容为空")
	}
	return res[0].Dst, nil
}

func (e *engine) LanguageCode(code string) string {
	switch code {
	case "es":
		return string(baidufanyi.LanguageSpa)
	case "fr":
		return string(baidufanyi.LanguageFra)
	case "da":
		return string(baidufanyi.LanguageDan)
	case "ro":
		return string(baidufanyi.LanguageRom)
	case "ko":
		return string(baidufanyi.LanguageKor)
	case "ja":
		return string(baidufanyi.LanguageJp)
	case "vi":
		return string(baidufanyi.LanguageVie)
	case "zh-TW":
		return string(baidufanyi.LanguageCht)
	}
	return code
}
