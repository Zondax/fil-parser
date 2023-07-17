package common

import (
	"github.com/filecoin-project/lotus/api"
	"github.com/zondax/znats/znats"
	"gorm.io/gorm"
)

type DataSourceConfig struct {
	Nats           *znats.ConfigNats
	InputTableName string
	NetworkName    string
}

type DataSource struct {
	Node   api.FullNode
	Db     *gorm.DB
	Config DataSourceConfig
}
