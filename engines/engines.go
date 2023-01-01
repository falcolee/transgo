package engines

import (
	"strings"
	"sync"

	"github.com/pkg/errors"
)

var lock sync.RWMutex
var engineMap = make(map[string]Engine)

// Register registers an Engine.
func Register(domains []string, e Engine) {
	lock.Lock()
	for _, domain := range domains {
		engineMap[domain] = e
	}
	lock.Unlock()
}

func Translate(domain string, text string, options ...EngineOptions) (string, error) {
	domain = strings.TrimSpace(domain)
	engine := engineMap[domain]
	if engine == nil {
		engine = engineMap[""]
	}
	option := Options{}

	// 遍历可选参数，然后分别调用匿名函数，将连接对象指针传入，进行修改
	for _, op := range options {
		// 遍历调用函数，进行数据修改
		op(&option)
	}
	results, err := engine.Translate(text, option)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return results, nil
}
