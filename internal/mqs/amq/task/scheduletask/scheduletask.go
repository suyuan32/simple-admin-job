// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scheduletask

import (
	"fmt"
	"log"

	"github.com/suyuan32/simple-admin-job/internal/svc"
)

type SchedulerTask struct {
	svcCtx *svc.ServiceContext
}

func NewSchedulerTask(svcCtx *svc.ServiceContext) *SchedulerTask {
	return &SchedulerTask{
		svcCtx: svcCtx,
	}
}

// Start starts the server.
func (s *SchedulerTask) Start() {
	s.Register()
	if err := s.svcCtx.AsynqScheduler.Run(); err != nil {
		log.Fatal(fmt.Errorf("failed to start mqtask server, error: %v", err))
	}
}

// Stop stops the server.
func (s *SchedulerTask) Stop() {
	s.svcCtx.AsynqScheduler.Shutdown()
}
