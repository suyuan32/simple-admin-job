package periodicconfig

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-job/ent"
	"github.com/suyuan32/simple-admin-job/ent/task"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
)

type EntConfigProvider struct {
	DB *ent.Client
}

func NewEntConfigProvider(db *ent.Client) *EntConfigProvider {
	return &EntConfigProvider{
		db,
	}
}

func (e *EntConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	configData, err := e.DB.Task.Query().Where(task.StatusEQ(common.StatusNormal)).All(context.Background())
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(logx.WithContext(context.Background()), err, "get task config error")
	}

	var result []*asynq.PeriodicTaskConfig

	for _, v := range configData {
		result = append(result, &asynq.PeriodicTaskConfig{
			Cronspec: v.CronExpression,
			Task:     asynq.NewTask(v.Pattern, []byte(v.Payload)),
			Opts:     nil,
		})
	}

	return result, nil
}
