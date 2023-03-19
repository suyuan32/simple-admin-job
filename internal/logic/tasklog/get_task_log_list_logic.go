package tasklog

import (
	"context"
	"github.com/suyuan32/simple-admin-job/ent/task"
	"github.com/suyuan32/simple-admin-job/ent/tasklog"

	"github.com/suyuan32/simple-admin-job/ent/predicate"
	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-job/job"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogListLogic {
	return &GetTaskLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskLogListLogic) GetTaskLogList(in *job.TaskLogListReq) (*job.TaskLogListResp, error) {
	var predicates []predicate.TaskLog
	if in.TaskId != 0 {
		predicates = append(predicates, tasklog.HasTasksWith(task.IDEQ(in.TaskId)))
	}

	result, err := l.svcCtx.DB.TaskLog.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &job.TaskLogListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &job.TaskLogInfo{
			Id:         v.ID,
			StartedAt:  v.StartedAt.UnixMilli(),
			FinishedAt: v.FinishedAt.UnixMilli(),
			Result:     uint32(v.Result),
		})
	}

	return resp, nil
}
