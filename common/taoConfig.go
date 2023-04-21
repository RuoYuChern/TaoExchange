package common

import (
	"os"

	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
)

type TaoConf struct {
	DbDns               string `yaml:"tao_infra_dns"`
	CoordinatorUrl      string `yaml:"tao_coordiantor_url"`
	CoordinatorPort     int32  `yaml:"tao_coordiantor_port"`
	CoordinatorRestPort int32  `yaml:"tao_coordiantor_restport"`
	GateRestPort        int32  `yaml:"tao_gate_restport"`
	ExchangePort        int32  `yaml:"tao_exchange_port"`
	StorePort           int32  `yaml:"tao_store_port"`
	AdapterPort         int32  `yaml:"tao_adapter_port"`
	MarketDataPort      int32  `yaml:"tao_market_data_port"`
	BrokerPort          int32  `yaml:"tao_broker_port"`
	BrokerDirPath       string `yaml:"tao_broker_dir"`
}

func (c *TaoConf) LoadTaoConf(path string) {
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
	slog.Info("DbDns:", c.DbDns, ",CoordinatorUrl: ", c.CoordinatorUrl)
}
