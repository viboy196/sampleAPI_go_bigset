package main

import (
	appconfig "SampleAPI_Bigset/config"
	"SampleAPI_Bigset/models"
	_ "SampleAPI_Bigset/routers"
	"log"
	"strings"

	"github.com/astaxie/beego"
)

func main() {

	config := &appconfig.AppConfig{}
	etcdEndpoints := beego.AppConfig.String("etcd::etcdEndpoints")
	config.EtcdServerEndpoints = strings.Split(etcdEndpoints, ",")
	config.SourceMappingNewSID = beego.AppConfig.String("source_mapping_new::sid")
	config.SourceMappingNewHost = beego.AppConfig.String("source_mapping_new::host")
	config.SourceMappingNewPort = beego.AppConfig.String("source_mapping_new::port")
	config.SourceMappingNewProtocol = beego.AppConfig.String("source_mapping_new::protocol")

	appconfig.Config = config
	log.Println(appconfig.Config, "appconfig.Config")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/v1/swagger"] = "swagger"
	}
	models.InitBigSetIf()
	beego.Run()

}
