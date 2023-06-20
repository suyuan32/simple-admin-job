package tasklog

import (
	"context"

	"github.com/suyuan32/simple-admin-job/ent/predicate"
	"github.com/suyuan32/simple-admin-job/ent/task"
	"github.com/suyuan32/simple-admin-job/ent/tasklog"
	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-job/types/job"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
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

	if in.TaskId != nil {
		predicates = append(predicates, tasklog.HasTasksWith(task.IDEQ(*in.TaskId)))
	}

	if in.Result != nil && *in.Result != 0 {
		predicates = append(predicates, tasklog.ResultEQ(uint8(*in.Result)))
	}

	result, err := l.svcCtx.DB.TaskLog.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &job.TaskLogListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &job.TaskLogInfo{
			Id:         &v.ID,
			StartedAt:  pointy.GetPointer(v.StartedAt.UnixMilli()),
			FinishedAt: pointy.GetPointer(v.FinishedAt.UnixMilli()),
			Result:     pointy.GetPointer(uint32(v.Result)),
		})
	}

	return resp, nil
}
