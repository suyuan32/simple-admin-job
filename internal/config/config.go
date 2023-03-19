package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/plugins/mq/asynq"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	RedisConf    redis.RedisConf
	AsynqConf    asynq.AsynqConf
	TaskConf     TaskConf
}

type TaskConf struct {
	EnableScheduledTask bool `json:",default=true"`
	EnableDPTask        bool `json:",default=true"`
}
