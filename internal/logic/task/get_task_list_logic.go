package task

import (
	"context"

	"github.com/suyuan32/simple-admin-job/ent/predicate"
	"github.com/suyuan32/simple-admin-job/ent/task"
	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskListLogic {
	return &GetTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskListLogic) GetTaskList(in *job.TaskListReq) (*job.TaskListResp, error) {
	var predicates []predicate.Task
	if in.Name != nil {
		predicates = append(predicates, task.NameContains(*in.Name))
	}
	if in.TaskGroup != nil {
		predicates = append(predicates, task.TaskGroupContains(*in.TaskGroup))
	}

	result, err := l.svcCtx.DB.Task.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &job.TaskListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &job.TaskInfo{
			Id:             &v.ID,
			CreatedAt:      pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:      pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:         pointy.GetPointer(uint32(v.Status)),
			Name:           &v.Name,
			TaskGroup:      &v.TaskGroup,
			CronExpression: &v.CronExpression,
			Pattern:        &v.Pattern,
			Payload:        &v.Payload,
		})
	}

	return resp, nil
}
