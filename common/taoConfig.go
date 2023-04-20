package common

import (
	"io/ioutil"

	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
)

type TaoConf struct {
	DbDns          string `yaml:"tao_infra_dns"`
	CoordinatorUrl string `yaml:"tao_coordiantor_url"`
}

func (c *TaoConf) LoadTaoConf(path string) {
	ymlFile, err := ioutil.ReadFile(path)
	if err != nil {
		slog.Error("Can not load file:", path, ", error:", err.Error())
		panic(err)
	}

	err = yaml.Unmarshal(ymlFile, c)
	if err != nil {
		slog.Error("unmarshal failed: ", err.Error())
		panic(err)
	}
	slog.Info("DbDns:", c.DbDns, ",CoordinatorUrl: ", c.CoordinatorUrl)
}
