package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/ini"
)

func NewConfig() Config{
c := config.Default().WithOptions()
c.AddDriver(ini.Driver)
err := c.LoadFiles("..\\..\\pkg\\config\\config.ini")
	if err != nil {
		panic(err)
	}
	return c
}

type Config interface{
	Set(key string, val any, setByPath ...bool) (err error)
	String(key string, defVal ...string) string

}