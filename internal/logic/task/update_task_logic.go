package task

import (
	"context"

	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-job/job"

	"github.com/suyuan32/simple-admin-common/i18n"

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
	err := l.svcCtx.DB.Task.UpdateOneID(in.Id).
		SetNotEmptyStatus(uint8(in.Status)).
		SetNotEmptyName(in.Name).
		SetNotEmptyTaskGroup(in.TaskGroup).
		SetNotEmptyCronExpression(in.CronExpression).
		SetNotEmptyPattern(in.Pattern).
		SetNotEmptyPayload(in.Payload).
		SetNotEmptyEntryID(int(in.EntryId)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
