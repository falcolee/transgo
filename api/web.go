package api

import (
	"strconv"
	"strings"

	"github.com/falcolee/transgo/common"
	"github.com/falcolee/transgo/common/utils/gologger"
	"github.com/falcolee/transgo/runner"
	"github.com/gin-gonic/gin"
)

func RunApiWeb(options *common.TLOptions) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "OK",
		})
	})
	r.POST("/translate", func(c *gin.Context) {
		text := c.PostForm("text")
		if text == "" {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "text is empty",
			})
			return
		}
		engine := c.PostForm("engine")
		engineList := make([]string, 0)
		if engine != "" {
			engineList = strings.Split(engine, ",")
		}
		options.Text = text
		options.Engines = engineList
		from := c.PostForm("from")
		to := c.PostForm("to")
		retry := c.PostForm("retry")
		if retry != "" {
			retryInt, _ := strconv.Atoi(retry)
			if retryInt > 0 {
				options.Retry = retryInt
			}
		}
		options.FromLanguage = from
		options.ToLanguage = to
		infos, errs := runner.RunJobWithRetry(options)
		c.JSON(200, gin.H{
			"code":    200,
			"message": "OK",
			"data":    infos,
			"error":   errs,
		})
	})
	server := ":32000"
	if options.TLConfig.Http.Server != "" {
		server = options.TLConfig.Http.Server
	}
	gologger.Infof("API 模式已开启 %s\n", server)
	err := r.Run(server)
	if err != nil {
		gologger.Fatalf("web api run error: %v", err)
		return
	} else {
		gologger.Infof("web api run success\n\n")
	}
}
