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
	serviceGroup.Add(scheduletask.NewSchedulerTask(ctx))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	serviceGroup.Start()
}
