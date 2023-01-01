package engines

import (
	"strings"

	"github.com/falcolee/transgo/common"
)

type Options struct {
	Region       string
	FromLanguage string
	ToLanguage   string
	ThreadNumber int
	AppId        string
	AppSecret    string
	TLConfig     *common.TLConfig
}

type EngineOptions func(*Options)

func WithFromLanguage(fromLanguage string) EngineOptions {
	return func(con *Options) {
		con.FromLanguage = fromLanguage
	}
}

func WithToLanguage(toLanguage string) EngineOptions {
	return func(con *Options) {
		con.ToLanguage = toLanguage
	}
}

func WithRegion(region string) EngineOptions {
	return func(con *Options) {
		con.Region = region
	}
}

func WithAppId(appId string) EngineOptions {
	return func(con *Options) {
		con.AppId = appId
	}
}

func WithAppSecret(appSecret string) EngineOptions {
	return func(con *Options) {
		con.AppSecret = appSecret
	}
}

func WithConfig(tlConfig *common.TLConfig) EngineOptions {
	return func(con *Options) {
		con.TLConfig = tlConfig
	}
}

func (options *Options) GetFromLanguage() string {
	if options.FromLanguage == "" {
		return "auto"
	}
	return options.FormatLanguageCode(options.FromLanguage)
}

func (options *Options) GetToLanguage() string {
	if options.ToLanguage == "" {
		return "en"
	}
	return options.FormatLanguageCode(options.ToLanguage)
}

func (options *Options) FormatLanguageCode(code string) string {
	isoCode := strings.Replace(code, "_", "-", -1)
	if strings.HasPrefix(isoCode, "en-") {
		return "en"
	}
	if strings.HasPrefix(isoCode, "fr-") {
		return "fr"
	}
	if strings.HasPrefix(isoCode, "es-") {
		return "es"
	}
	if strings.HasPrefix(isoCode, "pt-") {
		return "pt"
	}
	switch code {
	case "zh-CN":
		return "zh"
	case "zh-HK":
		return "zh-TW"
	}
	return code
}

type Engine interface {
	Translate(text string, options Options) (string, error)
	LanguageCode(code string) string
}
