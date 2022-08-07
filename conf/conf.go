package conf

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	Port          int
	MySqlUser     string
	MySqlPassword string
)

func Init() {
	cfg, err := ini.Load("./conf/my.ini")
	if err != nil {
		log.Fatal("configuration failed", err)
	}
	Port = cfg.Section("server").Key("port").MustInt(9090)
	MySqlUser = cfg.Section("mysql").Key("user").MustString("root")
	MySqlPassword = cfg.Section("mysql").Key("password").MustString("password")
}
