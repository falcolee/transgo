package outputfile

import (
	"fmt"

	"github.com/falcolee/transgo/common"
)

var TranslateInfosList = make([][]string, 0)
var Fields = []string{"原本", "译文", "引擎", "原始语言", "转换语言", "长度", "使用缓存"}

// MergeOutPut 数据合并到MAP
func MergeOutPut(transInfos *common.TranslateInfos) [][]string {
	useCache := "否"
	if transInfos.UseCache {
		useCache = "是"
	}
	row := []string{transInfos.Text, transInfos.Result, transInfos.Engine, transInfos.FromLanguage, transInfos.ToLanguage, fmt.Sprintf("%d", transInfos.Length), useCache}
	TranslateInfosList = append(TranslateInfosList, row)
	return TranslateInfosList
}

func OutPutMergeInfo(options *common.TLOptions) {
	common.TableShow(Fields, TranslateInfosList, options)
}
