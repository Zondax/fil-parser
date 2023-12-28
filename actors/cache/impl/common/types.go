package common

import (
	"github.com/filecoin-project/lotus/api"
	"github.com/zondax/golem/pkg/zcache"
	"github.com/zondax/znats/znats"
	"gorm.io/gorm"
)

type DataSourceConfig struct {
	Nats           *znats.ConfigNats
	Cache          *zcache.CombinedConfig
	InputTableName string
	NetworkName    string
}

type DataSource struct {
	Node   api.FullNode
	Db     *gorm.DB
	Config DataSourceConfig
}
