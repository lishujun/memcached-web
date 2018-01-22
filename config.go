package main

import (
    "github.com/lxmgo/config"
    l4g "github.com/alecthomas/log4go"
)


type AppConfig struct {
    Username string
    Password string
    ConnString string
}

var ConfigReader AppConfig

func InitConfigReader() {

    config, err := config.NewConfig("config.ini")
    if err != nil{
        l4g.Error("load config error, rason: %v", err)
        panic("load config error")
    }

    ConfigReader.Username = config.String("auth::username")
    ConfigReader.Password = config.String("auth::password")
    ConfigReader.ConnString = config.String("memcached::conn_string")

    l4g.Info("init called, the config is : %v", ConfigReader)
}
