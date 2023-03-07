package scheduletask

import (
	"log"

	"github.com/pkg/errors"

	"github.com/suyuan32/simple-admin-job/internal/svc"
)

type SchedulerTask struct {
	svcCtx *svc.ServiceContext
}

func NewSchedulerTask(svcCtx *svc.ServiceContext) *SchedulerTask {
	return &SchedulerTask{
		svcCtx: svcCtx,
	}
}

// Start starts the server.
func (s *SchedulerTask) Start() {
	s.Register()
	if err := s.svcCtx.AsynqScheduler.Run(); err != nil {
		log.Fatal(errors.Wrapf(err, "failed to start mqtask server, error: %v", err))
	}
}

// Stop stops the server.
func (s *SchedulerTask) Stop() {
	s.svcCtx.AsynqScheduler.Shutdown()
}
