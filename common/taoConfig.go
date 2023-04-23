package common

import (
	"os"

	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
)

type taoConf struct {
	CoordinatorUrl      string `yaml:"coordiantor_url"`
	CoordinatorPort     int32  `yaml:"coordiantor_port"`
	CoordinatorRestPort int32  `yaml:"coordiantor_restport"`
	GateRestPort        int32  `yaml:"gate_restport"`
	ExchangePort        int32  `yaml:"exchange_port"`
	StorePort           int32  `yaml:"store_port"`
	AdapterPort         int32  `yaml:"adapter_port"`
	MarketDataPort      int32  `yaml:"market_data_port"`
	BrokerPort          int32  `yaml:"broker_port"`
	BrokerDirPath       string `yaml:"broker_dir"`
}

type infraConf struct {
	DbDns string `yaml:"db_dns"`
}

type TaoAppConf struct {
	Tao   taoConf   `yaml:"tao"`
	Infra infraConf `yaml:"infra"`
}

func (c *TaoAppConf) LoadTaoConf(path string) {
	ymlFile, err := os.ReadFile(path)
	if err != nil {
		slog.Error("Can not load file:", path, ", error:", err.Error())
		panic(err)
	}

	err = yaml.Unmarshal(ymlFile, c)
	if err != nil {
		slog.Error("unmarshal failed: ", err.Error())
		panic(err)
	}
	slog.Info("DbDns:", c.Infra.DbDns)
	slog.Info("CoordinatorUrl: ", c.Tao.CoordinatorUrl)
}
