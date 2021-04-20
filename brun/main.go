package main

import (
	"flag"
	"github.com/gzl-tommy/infra"
	"github.com/gzl-tommy/infra/base"
	"github.com/sirupsen/logrus"
	"github.com/tietang/props/v3/consul"
	"github.com/tietang/props/v3/ini"
	"github.com/tietang/props/v3/kvs"
	_ "gzl-tommy/resk-ui"
)

func main() {
	flag.Parse()

	profile := flag.Arg(0)
	if profile == "" {
		profile = "dev"
	}

	//获取程序运行文件所在的路径
	file := kvs.CurrentFilePath("boot.ini", 1)
	logrus.Info(file)

	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	if _, e := conf.Get("profile"); e != nil {
		conf.Set("profile", profile)
	}

	addr := conf.GetDefault("consul.address", "127.0.0.1:8500")
	contexts := conf.KeyValue("consul.contexts").Strings()
	logrus.Info("consul address:", addr)
	logrus.Info("consul contexts:", contexts)

	consulConf := consul.NewCompositeConsulConfigSourceByType(contexts, addr, kvs.ContentIni)
	consulConf.Add(conf)

	base.InitLog(consulConf)
	app := infra.New(consulConf)
	app.Start()
}
