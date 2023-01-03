package common

import (
	"io"
	"os"
	"path/filepath"

	"github.com/falcolee/transgo/common/utils"
	"github.com/falcolee/transgo/common/utils/gologger"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func Parse(options *TLOptions) {
	if options.IsApiMode {
		Banner()
	}
	if options.ConfigFile == "" {
		home, err := homedir.Dir()
		if err != nil {
			gologger.Fatalf("配置文件目录查找失败 %s\n", err)
		}
		configDir := utils.GetConfigPath()
		if options.ConfigDir != "" {
			configDir = options.ConfigDir
		}
		viper.AddConfigPath(home)
		viper.AddConfigPath(configDir)
		viper.AddConfigPath(".")
		configFileName := "transgo"
		viper.SetConfigType("yaml")
		viper.SetConfigName(configFileName)
		if err := viper.ReadInConfig(); err != nil {
			gologger.Infof("Current Version: %s\n", GitTag)
			defaultCfgYName := filepath.Join(configDir, configFileName+".yaml")
			f, _ := os.Create(defaultCfgYName) //创建文件
			_, errs := io.WriteString(f, configYaml)
			if errs != nil {
				gologger.Fatalf("配置文件创建失败 %s\n", errs)
			} else {
				gologger.Infof("配置文件生成成功，生成目录：%s\n", configDir)
			}
		}
	} else {
		viper.SetConfigFile(options.ConfigFile)
		err := viper.ReadInConfig()
		// 配置文件检查
		if err != nil {
			gologger.Fatalf("读取配置文件 %s 失败，错误: %s\n", options.ConfigFile, err)
		}
	}
	//加载配置信息~
	conf := new(TLConfig)
	err := viper.Unmarshal(&conf)
	if err != nil {
		gologger.Fatalf("配置文件解析错误 #%v ", err)
	}
	options.TLConfig = conf
	if options.UseCache {
		options.UseCache = conf.Cache.UseCache
	}
	options.StorageDir = conf.Cache.StorageDir
	options.CacheStorage = conf.Cache.CacheStorage

	//DEBUG模式设定
	if options.IsDebug {
		gologger.MaxLevel = gologger.Debug
	}

	// 是否为API模式 加入基本参数判断
	if !options.IsApiMode {
		if options.Text == "" && options.InputFile == "" {
			gologger.Fatalf("请输入要翻译的文本或批量文本文件路径\n")
		}
		if options.InputFile != "" {
			if ok := utils.FileExists(options.InputFile); !ok {
				gologger.Fatalf("没有输入文件 %s\n", options.InputFile)
			}
		}
	}
}
