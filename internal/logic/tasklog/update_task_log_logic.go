package tasklog

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-job/job"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogLogic {
	return &UpdateTaskLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskLogLogic) UpdateTaskLog(in *job.TaskLogInfo) (*job.BaseResp, error) {
	err := l.svcCtx.DB.TaskLog.UpdateOneID(in.Id).
		SetFinishedAt(time.Unix(in.FinishedAt, 0)).
		SetNotEmptyResult(uint8(in.Result)).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
