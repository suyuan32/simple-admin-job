package scheduletask

import (
	"github.com/hibiken/asynq"

	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/types/pattern"
)

// Register adds task to cron. | 在此处定义定时任务
func (s *SchedulerTask) Register() {
	// register task to schedule | 注册任务到调度器
	s.svcCtx.AsynqScheduler.Register("@every 5s", asynq.NewTask(pattern.RecordHelloWorld,
		[]byte("{\"name\": \"Jack (Scheduled Task every 5s)\"}")))
}
