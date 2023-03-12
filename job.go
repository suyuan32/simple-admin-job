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

package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/suyuan32/simple-admin-job/internal/config"
	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/task/dynamicperiodictask"
	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/task/mqtask"
	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/task/scheduletask"
	"github.com/suyuan32/simple-admin-job/internal/server"
	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/job"
)

var configFile = flag.String("f", "etc/job.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		job.RegisterJobServer(grpcServer, server.NewJobServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	go func() {
		s.Start()
	}()

	serviceGroup := service.NewServiceGroup()
	defer func() {
		serviceGroup.Stop()
		logx.Close()
	}()

	serviceGroup.Add(mqtask.NewMQTask(ctx))
	serviceGroup.Add(dynamicperiodictask.NewDPTask(ctx))
	serviceGroup.Add(scheduletask.NewSchedulerTask(ctx))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	serviceGroup.Start()
}
