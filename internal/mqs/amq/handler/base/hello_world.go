package base

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/types/payload"
	"github.com/suyuan32/simple-admin-job/internal/svc"
)

type HelloWorldHandler struct {
	svcCtx *svc.ServiceContext
}

func NewHelloWorldHandler(svcCtx *svc.ServiceContext) *HelloWorldHandler {
	return &HelloWorldHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask if return err != nil , asynq will retry | 如果返回错误不为空则会重试
func (l *HelloWorldHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p payload.HelloWorldPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(err, "failed to umarshal the payload :%s", string(t.Payload()))
	}

	fmt.Printf("Hi! %s\n", p.Name)

	return nil
}
