package task

import (
	"context"

	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-job/job"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskByIdLogic {
	return &GetTaskByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskByIdLogic) GetTaskById(in *job.IDReq) (*job.TaskInfo, error) {
	result, err := l.svcCtx.DB.Task.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.TaskInfo{
		Id:             result.ID,
		CreatedAt:      result.CreatedAt.UnixMilli(),
		UpdatedAt:      result.UpdatedAt.UnixMilli(),
		Status:         uint32(result.Status),
		Name:           result.Name,
		TaskGroup:      result.TaskGroup,
		CronExpression: result.CronExpression,
		Pattern:        result.Pattern,
		Payload:        result.Payload,
	}, nil
}
