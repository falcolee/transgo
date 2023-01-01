package common

import (
	"time"

	"github.com/falcolee/transgo/common/utils"
	"github.com/mkideal/cli"
)

type TLOptions struct {
	cli.Helper
	Engines      []string  `cli:"engine,e" usage:"指定翻译引擎，支持多个，如-e baidu -e qq"`
	Text         string    `cli:"text,w" usage:"要翻译的文本"`
	InputFile    string    `cli:"input,i" usage:"批量查询，文本按行分隔，输入txt路径"`
	FromLanguage string    `cli:"from,f" usage:"源语言，默认为自动" dft:"auto"`
	ToLanguage   string    `cli:"to,t" usage:"翻译语言，默认为英文" dft:"en"`
	AppId        string    `cli:"appid,a" usage:"翻译引擎接口账户，如appid,appkey等"`
	AppSecret    string    `cli:"secret,s" usage:"翻译引擎接口秘钥，如appsecret,secretkey等"`
	Region       string    `cli:"region,r" usage:"部分翻译引擎提供地区(可选)"`
	ConfigFile   string    `cli:"config,c" usage:"配置文件路径，yaml文件"`
	ConfigDir    string    `cli:"config-dir,k" usage:"配置文件生成目录，默认读取目录(可选)"`
	DelayTime    int64     `cli:"delay,l" usage:"每个请求延迟（S）默认无延迟，-1为随机延迟1-5秒"`
	Version      bool      `cli:"version,v" usage:"版本信息"`
	IsDebug      bool      `cli:"debug,u" usage:"是否显示debug详细信息"`
	IsApiMode    bool      `cli:"api,p" usage:"是否API模式"`
	IsDaemon     bool      `cli:"daemon,d" usage:"是否在后台运行，开启守护进程，仅在API模式中生效"`
	UseCache     bool      `cli:"cache,x" usage:"是否启用缓存，默认启用" dft:"true"`
	SplitCache   bool      `cli:"split,y" usage:"开启全局缓存"`
	CacheStorage string    `cli:"storage,m" usage:"缓存引擎，默认使用内存，支持memory/redis/file" dft:"memery"`
	StorageDir   string    `cli:"dir,z" usage:"如开启文件缓存，需指定文件缓存目录"`
	TLConfig     *TLConfig `cli:"-"`
}

func (h *TLOptions) GetDelayRTime() int64 {
	if h.DelayTime == 0 {
		return 0
	}
	return utils.RangeRand(0, h.DelayTime)
}

func (h *TLOptions) GetTLConfig() *TLConfig {
	return h.TLConfig
}

// TLConfig YML配置文件
type TLConfig struct {
	Cache struct {
		UseCache      bool   `yaml:"useCache"`
		CacheStorage  string `yaml:"cacheStorage"`
		StorageDir    string `yaml:"storageDir"`
		RedisAddr     string `yaml:"redisAddr"`
		RedisPassword string `yaml:"redisPassword"`
		RedisPrefix   string `yaml:"redisPrefix"`
	}
	Http struct {
		Server string `yaml:"server"`
	}
	Engine struct {
		Default string `yaml:"default"`
	}
	Baidu struct {
		AppId     string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	}
	Tencent struct {
		AppId     string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	}
	Aliyun struct {
		AppId     string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	}
	Volcengine struct {
		AppId     string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	}
}

type TranslateInfos struct {
	Text         string
	Result       string
	Engine       string
	FromLanguage string
	ToLanguage   string
	UseCache     bool
	Length       int
	InTime       time.Time
}

var (
	BuiltAt   string
	GoVersion string
	GitAuthor string
	GitTag    string
)
var version = "0.1"
var configYaml = `version: 0.1
cache:
  useCache: true       	# 开启缓存
  cacheStorage: 'memory'	# 缓存类型，支持memory/file/redis
  storageDir: ''			# 文件缓存目录
  redisAddr: ''			# redis连接地址，127.0.0.1:6379
  redisPassword: ''		# redis密码
  redisPrefix: ''		# redis前缀
http:
  server: ':32000'      # API启动端口
engine:
  default: 'baidu'      # 不指定时的默认翻译引擎，支持多个用,分割
baidu:
  appId: ''            # 百度接口ID
  appSecret: ''        # 百度接口秘钥
tencent:
  appId: ''         # 腾讯接口ID
  appSecret: ''        # 腾讯接口秘钥
aliyun:
  appId: ''         # 阿里云接口ID
  appSecret: ''     # 阿里云接口秘钥
volcengine:
  appId: ''        # 火山接口ID
  appSecret: ''        # 火山接口秘钥
`
