// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package svc

import (
	"github.com/hibiken/asynq"
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
	}
}
