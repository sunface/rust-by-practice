package misc

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Common struct {
		Version    string
		LogLevel   string
		AdminToken string
	}
	Agent struct {
		VgoAddr          string
		ReportInterval   int
		PinpointCacheLen int
		CmdCacheLen      int
		ReportLen        int
		KeepliveInterval int
	}
}

var Conf *Config

func InitConfig(path string) {
	conf := &Config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("InitConfig:ioutil.ReadFile", err)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal("InitConfig:yaml.Unmarshal", err)
	}
	Conf = conf
	log.Println(Conf.Agent)
}
