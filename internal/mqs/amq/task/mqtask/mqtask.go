package mqtask

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"github.com/suyuan32/simple-admin-job/internal/svc"
)

type MQTask struct {
	svcCtx *svc.ServiceContext
	mux    *asynq.ServeMux
}

func NewMQTask(svcCtx *svc.ServiceContext) *MQTask {
	return &MQTask{
		svcCtx: svcCtx,
	}
}

// Start starts the server.
func (m *MQTask) Start() {
	m.Register()
	if err := m.svcCtx.AsynqServer.Run(m.mux); err != nil {
		log.Fatal(errors.Wrapf(err, "failed to start mqtask server, error: %v", err))
	}
}

// Stop stops the server.
func (m *MQTask) Stop() {
	m.svcCtx.AsynqServer.Stop()
	m.svcCtx.AsynqServer.Shutdown()
}
