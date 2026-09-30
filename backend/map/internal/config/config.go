package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Etcd     discov.EtcdConf `json:"Etcd"`
	Amap_Key string          `json:"Amap_Key"`
}
