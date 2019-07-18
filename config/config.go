/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-13
* Time: 上午8:41
* */
package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type myconf struct {
	App struct {
		Host    string `yaml:"host"`
		BaseUrl string `yaml:"baseurl"`
		Debug   bool   `yaml:"debug"`
	}
	Mysql struct {
		Dsn   string `yaml:"dsn"`
		Cache bool   `yaml:"cache"`
	}
	Redis struct {
		Maxidle     int           `yaml:"maxidle"`
		MaxActive   int           `yaml:"max_active"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
		Port        string        `yaml:"port"`
	}
}

var (
	MyConfig *myconf
)

func init() {
	MyConfig = &myconf{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, MyConfig)
	if e != nil {
		panic(e.Error())
	}

	if MyConfig.App.Debug {
		fmt.Println(MyConfig)
	}
}
