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

type UpdateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogic {
	return &UpdateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskLogic) UpdateTask(in *job.TaskInfo) (*job.BaseResp, error) {
	err := l.svcCtx.DB.Task.UpdateOneID(*in.Id).
		SetNotNilStatus(pointy.GetPointer(uint8(*in.Status))).
		SetNotNilName(in.Name).
		SetNotNilTaskGroup(in.TaskGroup).
		SetNotNilCronExpression(in.CronExpression).
		SetNotNilPattern(in.Pattern).
		SetNotNilPayload(in.Payload).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.BaseResp{Msg: i18n.CreateSuccess}, nil
}
