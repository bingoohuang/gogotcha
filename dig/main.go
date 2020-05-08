package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
	"go.uber.org/dig"
	ini "gopkg.in/ini.v1"
)

type Option struct {
	ConfigFile string `short:"c" long:"config" default:"./my.ini" description:"Name of config file."`
}

func InitOption() (*Option, error) {
	var opt Option
	_, err := flags.Parse(&opt)

	return &opt, err
}

func InitConf(opt *Option) (*ini.File, error) {
	cfg, err := ini.Load(opt.ConfigFile)
	return cfg, err
}

func PrintInfo(cfg *ini.File) {
	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())
}

func main() {
	container := dig.New()

	_ = container.Provide(InitOption)
	_ = container.Provide(InitConf)

	_ = container.Invoke(PrintInfo)
}
