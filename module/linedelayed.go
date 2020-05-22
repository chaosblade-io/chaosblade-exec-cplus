/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package module

import (
	"context"
	"path"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/sirupsen/logrus"

	"github.com/chaosblade-io/chaosblade-exec-cplus/common"
)

type LineDelayedActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewLineDelayedActionSpec() spec.ExpActionCommandSpec {
	return &LineDelayedActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:     "delayDuration",
					Desc:     "delay time, unit is second",
					Required: true,
				},
			},
			ActionExecutor: &LineDelayedExecutor{},
		},
	}
}

func (l LineDelayedActionSpec) Name() string {
	return "delay"
}

func (l LineDelayedActionSpec) Aliases() []string {
	return []string{}
}

func (l LineDelayedActionSpec) ShortDesc() string {
	return "Code line delayed"
}

func (l LineDelayedActionSpec) LongDesc() string {
	return "Code line delayed"
}

type LineDelayedExecutor struct {
	channel spec.Channel
}

func (l *LineDelayedExecutor) Name() string {
	return "delay"
}

func (l *LineDelayedExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	delayDuration := model.ActionFlags["delayDuration"]
	if delayDuration == "" {
		return spec.ReturnFail(spec.Code[spec.IllegalParameters], "less necessary delayDuration value")
	}
	// search pid by process name
	processName := model.ActionFlags["processName"]
	if processName == "" {
		return spec.ReturnFail(spec.Code[spec.IllegalParameters], "less necessary processName value")
	}
	processCtx := context.WithValue(context.Background(), channel.ExcludeProcessKey, "blade")
	pids, err := channel.NewLocalChannel().GetPidsByProcessName(processName, processCtx)
	if err != nil {
		logrus.Warnf("get pids by %s process name err, %v", processName, err)
	}
	localChannel := common.NewAsyncChannel()
	if pids == nil || len(pids) == 0 {
		args := buildArgs([]string{
			model.ActionFlags["fileLocateAndName"],
			model.ActionFlags["forkMode"],
			model.ActionFlags["libLoad"],
			model.ActionFlags["breakLine"],
			delayDuration,
			model.ActionFlags["initParams"],
		})
		return localChannel.Run(context.Background(), path.Join(common.GetScriptPath(), common.ResponseDelayScript), args)
	} else {
		args := buildArgs([]string{
			pids[0],
			model.ActionFlags["forkMode"],
			"",
			"",
			model.ActionFlags["breakLine"],
			delayDuration,
			model.ActionFlags["initParams"],
		})
		return localChannel.Run(context.Background(), path.Join(common.GetScriptPath(), common.ResponseDelayAttachScript), args)
	}
}

func (l *LineDelayedExecutor) SetChannel(channel spec.Channel) {
	l.channel = channel
}
