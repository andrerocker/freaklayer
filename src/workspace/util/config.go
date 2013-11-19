package util

import (
    "flag"
    "github.com/vaughan0/go-ini"
)

var configuration ini.File

func InitConfig() (err error) {
    configFile := flag.String("config", "/etc/bricklayer/bricklayer.conf", "BrickLayer Configuration File")
    flag.Parse()
    configuration, err = ini.LoadFile(*configFile)
    return err
}

func GetConfig(section string, key string) string {
    value, _ := configuration.Get(section, key)
    return value
}
