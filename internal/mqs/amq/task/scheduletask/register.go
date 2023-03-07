package scheduletask

import (
	"github.com/hibiken/asynq"

	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/types/pattern"
)

// Register adds task to cron.
func (s *SchedulerTask) Register() {
	s.svcCtx.AsynqScheduler.Register("@every 5s", asynq.NewTask(pattern.RecordHelloWorld, []byte("{\"name\": \"Jack\"}")))
}
