package runner

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/falcolee/transgo/cache"
	"github.com/falcolee/transgo/common"
	"github.com/falcolee/transgo/common/outputfile"
	"github.com/falcolee/transgo/common/utils"
	"github.com/falcolee/transgo/common/utils/gologger"
	"github.com/falcolee/transgo/engines"
)

// RunEnumeration 普通任务命令行模式，可批量导入文件查询
func RunEnumeration(options *common.TLOptions) {
	if options.InputFile != "" {
		res := utils.ReadFile(options.InputFile)
		time.Sleep(time.Duration(options.GetDelayRTime()) * time.Second)

		for k, v := range res {
			if v == "" {
				gologger.Errorf("【第%d条】文本为空，自动跳过\n", k+1)
				continue
			}
			gologger.Infof("\n====================\n【第%d条】文本 %s 查询中\n====================\n", k+1, v)
			options.Text = v

			RunJob(options)
		}
	} else {
		RunJob(options)
	}
	outputfile.OutPutMergeInfo(options)
}

func RunJobWithRetry(options *common.TLOptions) (infos []common.TranslateInfos, errors []error) {
	if options.Retry > 0 {
		for i := 0; i < options.Retry; i++ {
			infos, errors = RunJob(options)
			if len(errors) == 0 {
				break
			}
			fallbackEngins := strings.Split(options.TLConfig.Engine.Fallback, ",")
			if utils.CheckList(fallbackEngins) {
				for _, fallbackEngine := range fallbackEngins {
					if !utils.InStrArray(fallbackEngine, options.Engines) {
						options.Engines = []string{fallbackEngine}
					}
				}
			}
		}
	} else {
		infos, errors = RunJob(options)
	}
	return
}

func RunJob(options *common.TLOptions) (infos []common.TranslateInfos, errors []error) {
	if len(options.Engines) == 0 {
		options.Engines = strings.Split(options.TLConfig.Engine.Default, ",")
	}
	if !utils.CheckList(options.Engines) {
		options.Engines = []string{"baidu"}
	}
	if len(options.Engines) > 1 {
		options.SplitCache = true
	}
	gologger.Infof("文本:【%s】翻译引擎：%s \n", options.Text, strings.Join(options.Engines, ","))
	infos = make([]common.TranslateInfos, 0)
	errors = make([]error, 0)
	var wg sync.WaitGroup
	for _, v := range options.Engines {
		wg.Add(1)
		var domain = v
		go func() {
			transInfos, err := Translate(domain, options)
			if err == nil {
				infos = append(infos, *transInfos)
				outputfile.MergeOutPut(transInfos)
			} else {
				errors = append(errors, err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return
}

func Translate(domain string, options *common.TLOptions) (*common.TranslateInfos, error) {
	transKey := fmt.Sprintf("TranslateK_%s_%s_%s", options.Text, options.FromLanguage, options.ToLanguage)
	transEngineKey := fmt.Sprintf("TranslateK_%s_%s_%s_%s", options.Text, options.FromLanguage, options.ToLanguage, domain)
	length := utf8.RuneCountInString(options.Text)
	translateInfos := &common.TranslateInfos{
		Text:         options.Text,
		FromLanguage: options.FromLanguage,
		ToLanguage:   options.ToLanguage,
		Engine:       domain,
		UseCache:     false,
		Length:       length,
		InTime:       time.Now(),
	}
	if options.UseCache {
		var result []byte
		if options.SplitCache { //分引擎缓存
			result, _ = cache.Storage.FetchOne(transEngineKey)
		} else {
			result, _ = cache.Storage.FetchOne(transKey)
		}
		if len(result) > 0 {
			translateInfos.UseCache = true
			translateInfos.Result = string(result)
			return translateInfos, nil
		}
	}
	var engineOptions = []engines.EngineOptions{
		engines.WithAppId(options.AppId),
		engines.WithAppSecret(options.AppSecret),
		engines.WithRegion(options.Region),
		engines.WithFromLanguage(options.FromLanguage),
		engines.WithToLanguage(options.ToLanguage),
		engines.WithConfig(options.TLConfig),
	}
	result, err := engines.Translate(domain, options.Text, engineOptions...)
	if err != nil {
		gologger.Errorf("文本:【%s】翻译失败：%s \n", options.Text, err)
		return translateInfos, err
	}

	if options.UseCache && result != "" {
		if options.SplitCache {
			err = cache.Storage.Store(transEngineKey, []byte(result))
		} else {
			exists, _ := cache.Storage.KeyExists(transKey)
			if !exists {
				err = cache.Storage.Store(transKey, []byte(result))
			}
		}
		if err != nil {
			gologger.Errorf("缓存写入失败：%s \n", err)
		}
	}
	translateInfos.InTime = time.Now()
	translateInfos.Result = result
	return translateInfos, nil
}
