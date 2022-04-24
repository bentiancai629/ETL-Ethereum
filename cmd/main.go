package main

import (
	"ETL-Ethereum/cmd/tools"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/beego/beego/v2/core/logs"
	"github.com/urfave/cli"

	"ETL-Ethereum/conf"
	"ETL-Ethereum/listener"
)

var (
	logLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: 1,
	}

	configPathFlag = cli.StringFlag{
		Name:  "config",
		Usage: "Server config file `<path>`",
		Value: "./conf/config.json",
	}

	initFlag = cli.BoolFlag{
		Name:   "init",
		Usage:  "Server will init or update db when true",
		Hidden: false,
	}
)

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "ETL-Ethereum"
	app.Action = StartServer
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		logLevelFlag,
		configPathFlag,
		initFlag,
	}
	app.Commands = []cli.Command{}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func StartServer(ctx *cli.Context) {
	//configFile := ctx.GlobalString("config")
	init := ctx.GlobalBool("init")
	if init {
		config, err := conf.LoadConfig()
		if config == nil || err != nil {
			logs.Error("startServer - read config failed!")
			return
		}

		tools.Deploy(config)
		return
	}

	for true {
		startServer(ctx)
		sig := waitSignal()
		stopServer()
		if sig != syscall.SIGHUP {
			break
		} else {
			continue
		}
	}
}

func startServer(ctx *cli.Context) {
	//configFile := ctx.GlobalString("config")
	config, err := conf.LoadConfig()
	if config == nil || err != nil {
		logs.Error("startServer - read config failed!")
		return
	}
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s", "perm":"0660"}`, "logs/monitor.log"))

	conf, _ := json.Marshal(config)
	logs.Info("%s", string(conf))
	listener.StartLandListen(config.DBConfig, config.ChainListenConfig, config.TokenAddressConfig)
}

func waitSignal() os.Signal {
	exit := make(chan os.Signal, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(sc)
	go func() {
		for sig := range sc {
			logs.Info("cross chain listen received signal:(%s).", sig.String())
			exit <- sig
			close(exit)
			break
		}
	}()
	sig := <-exit
	return sig
}

func stopServer() {
	listener.StopLandListen()
}

func main() {
	if err := setupApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
