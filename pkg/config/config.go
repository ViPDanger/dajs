package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/ini"
)

func NewConfig(name string) Config{
c := config.Default().WithOptions()
c.AddDriver(ini.Driver)
err := c.LoadFiles(name)
	if err != nil {
		panic(err)
	}
	return c
}

type Config interface{
	Set(key string, val any, setByPath ...bool) (err error)
	String(key string, defVal ...string) string
}