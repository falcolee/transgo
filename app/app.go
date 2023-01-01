package app

import (
	"math/rand"
	"os"
	"time"

	"github.com/VividCortex/godaemon"
	"github.com/falcolee/transgo/api"
	"github.com/falcolee/transgo/cache"
	"github.com/falcolee/transgo/common"
	"github.com/falcolee/transgo/runner"
	"github.com/mkideal/cli"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (app *App) Run() {
	rand.Seed(time.Now().Unix())
	os.Exit(cli.Run(new(common.TLOptions), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*common.TLOptions)
		common.Parse(argv)
		cache.InitCache(argv)
		if argv.IsApiMode {
			app.RunWeb(argv)
		} else {
			app.RunCli(argv)
		}
		return nil
	}))
}

func (app *App) RunCli(options *common.TLOptions) {
	runner.RunEnumeration(options)
}

func (app *App) RunWeb(options *common.TLOptions) {
	if options.IsDaemon {
		godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	}
	go api.RunApiWeb(options)
	select {} // 阻塞
}
