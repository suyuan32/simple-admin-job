package mqtask

import (
	"github.com/hibiken/asynq"

	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/handler/base"
	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/types/pattern"
)

// Register adds task to cron.
func (m *MQTask) Register() {
	mux := asynq.NewServeMux()

	mux.Handle(pattern.RecordHelloWorld, base.NewHelloWorldHandler(m.svcCtx))

	m.mux = mux
}
