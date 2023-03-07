package svc

import (
	"github.com/hibiken/asynq"

	"github.com/suyuan32/simple-admin-job/ent"
	"github.com/suyuan32/simple-admin-job/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config         config.Config
	DB             *ent.Client
	Redis          *redis.Redis
	AsynqServer    *asynq.Server
	AsynqScheduler *asynq.Scheduler
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
		Redis:          redis.MustNewRedis(c.RedisConf),
	}
}
