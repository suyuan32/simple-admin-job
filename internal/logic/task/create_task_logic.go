package task

import (
	"context"

	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTaskLogic) CreateTask(in *job.TaskInfo) (*job.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Task.Create().
		SetNotNilStatus(pointy.GetPointer(uint8(*in.Status))).
		SetNotNilName(in.Name).
		SetNotNilTaskGroup(in.TaskGroup).
		SetNotNilCronExpression(in.CronExpression).
		SetNotNilPattern(in.Pattern).
		SetNotNilPayload(in.Payload).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
