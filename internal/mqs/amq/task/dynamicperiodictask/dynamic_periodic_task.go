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

package dynamicperiodictask

import (
	"fmt"
	"log"

	"github.com/suyuan32/simple-admin-job/internal/svc"
)

type DPTask struct {
	svcCtx *svc.ServiceContext
}

func NewDPTask(svcCtx *svc.ServiceContext) *DPTask {
	return &DPTask{
		svcCtx: svcCtx,
	}
}

// Start starts the server.
func (m *DPTask) Start() {
	if err := m.svcCtx.AsynqPTM.Run(); err != nil {
		log.Fatal(fmt.Errorf("failed to start dptask server, error: %v", err))
	}
}

// Stop stops the server.
func (m *DPTask) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("DPTask shuts down successfully")
		}
	}()
	m.svcCtx.AsynqPTM.Shutdown()
}
