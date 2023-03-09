package svc

import (
	"github.com/hibiken/asynq"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/suyuan32/simple-admin-job/ent"
	"github.com/suyuan32/simple-admin-job/internal/config"
	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/types/periodicconfig"
)

type ServiceContext struct {
	Config         config.Config
	DB             *ent.Client
	Redis          *redis.Redis
	AsynqServer    *asynq.Server
	AsynqScheduler *asynq.Scheduler
	AsynqPTM       *asynq.PeriodicTaskManager
	CoreRpc        coreclient.Core
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config:         c,
		DB:             db,
		AsynqServer:    c.AsynqConf.NewServer(),
		AsynqScheduler: c.AsynqConf.NewScheduler(),
		AsynqPTM:       c.AsynqConf.NewPeriodicTaskManager(periodicconfig.NewEntConfigProvider(db)),
		Redis:          redis.MustNewRedis(c.RedisConf),
		CoreRpc:        coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
	}
}
