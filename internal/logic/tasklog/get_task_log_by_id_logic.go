package tasklog

import (
	"context"

	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogByIdLogic {
	return &GetTaskLogByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskLogByIdLogic) GetTaskLogById(in *job.IDReq) (*job.TaskLogInfo, error) {
	result, err := l.svcCtx.DB.TaskLog.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.TaskLogInfo{
		Id:         &result.ID,
		StartedAt:  pointy.GetPointer(result.StartedAt.UnixMilli()),
		FinishedAt: pointy.GetPointer(result.FinishedAt.UnixMilli()),
		Result:     pointy.GetPointer(uint32(result.Result)),
	}, nil
}
