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

type ErrorReturnedActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewErrorReturnedActionSpec() spec.ExpActionCommandSpec {
	return &ErrorReturnedActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:     "returnValue",
					Desc:     "Value returned. If you want return null, set --returnValue null",
					Required: true,
				},
			},
			ActionExecutor: &ErrorReturnedExecutor{},
		},
	}
}

func (e ErrorReturnedActionSpec) Name() string {
	return "return"
}

func (e ErrorReturnedActionSpec) Aliases() []string {
	return []string{}
}

func (e ErrorReturnedActionSpec) ShortDesc() string {
	return "error returned"
}

func (e ErrorReturnedActionSpec) LongDesc() string {
	return "error returned"
}

type ErrorReturnedExecutor struct {
	channel spec.Channel
}

func (e *ErrorReturnedExecutor) Name() string {
	return "return"
}

func (e *ErrorReturnedExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	if _, ok := spec.IsDestroy(ctx); ok {
		return spec.ReturnFail(spec.Code[spec.ServerError], "illegal processing")
	}
	returnValue := model.ActionFlags["returnValue"]
	if returnValue == "" {
		return spec.ReturnFail(spec.Code[spec.IllegalParameters], "less necessary returnValue value")
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
			returnValue,
			model.ActionFlags["initParams"],
		})
		return localChannel.Run(context.Background(), path.Join(common.GetScriptPath(), common.BreakAndReturnScript), args)
	} else {
		args := buildArgs([]string{
			pids[0],
			model.ActionFlags["forkMode"],
			"",
			"",
			model.ActionFlags["breakLine"],
			returnValue,
			model.ActionFlags["initParams"],
		})
		return localChannel.Run(context.Background(), path.Join(common.GetScriptPath(), common.BreakAndReturnAttachScript), args)
	}
}

func (e *ErrorReturnedExecutor) SetChannel(channel spec.Channel) {
	e.channel = channel
}
