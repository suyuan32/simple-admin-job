package tasklog

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogLogic {
	return &CreateTaskLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTaskLogLogic) CreateTaskLog(in *job.TaskLogInfo) (*job.BaseIDResp, error) {
	result, err := l.svcCtx.DB.TaskLog.Create().
		SetNotNilFinishedAt(pointy.GetPointer(time.Unix(*in.FinishedAt, 0))).
		SetNotNilResult(pointy.GetPointer(uint8(*in.Result))).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
