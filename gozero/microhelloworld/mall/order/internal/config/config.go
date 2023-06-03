package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// 去ETCD获取user rpc服务的配置
	UserRpc zrpc.RpcClientConf
}
