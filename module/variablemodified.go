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

type VariableModifiedActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewVariableModifiedActionSpec() spec.ExpActionCommandSpec {
	return &VariableModifiedActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:     "variableName",
					Desc:     "The name of the modified variable",
					Required: true,
				},
				&spec.ExpFlag{
					Name:     "variableValue",
					Desc:     "The value of the modified variable",
					Required: true,
				},
			},
			ActionExecutor: &VariableModifiedExecutor{},
		},
	}
}

func (v *VariableModifiedActionSpec) Name() string {
	return "modify"
}

func (v *VariableModifiedActionSpec) Aliases() []string {
	return []string{}
}

func (v *VariableModifiedActionSpec) ShortDesc() string {
	return "Modify value of the variable in source code when program running"
}

func (v *VariableModifiedActionSpec) LongDesc() string {
	return "Modify value of the variable in source code when program running"
}

type VariableModifiedExecutor struct {
	channel spec.Channel
}

func (v *VariableModifiedExecutor) Name() string {
	return "modify"
}

func (v *VariableModifiedExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	variableName := model.ActionFlags["variableName"]
	if variableName == "" {
		return spec.ReturnFail(spec.Code[spec.IllegalParameters], "less necessary variableName value")
	}
	variableValue := model.ActionFlags["variableValue"]
	if variableValue == "" {
		return spec.ReturnFail(spec.Code[spec.IllegalParameters], "less necessary variableValue value")
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
			variableName,
			variableValue,
			model.ActionFlags["initParams"],
		})
		return localChannel.Run(context.Background(), path.Join(common.GetScriptPath(), common.ModifyVariableScript), args)
	} else {
		args := buildArgs([]string{
			pids[0],
			model.ActionFlags["forkMode"],
			"",
			"",
			model.ActionFlags["breakLine"],
			variableName,
			variableValue,
			model.ActionFlags["initParams"],
		})
		return localChannel.Run(context.Background(), path.Join(common.GetScriptPath(), common.ModifyVariableAttachScript), args)
	}
}

func (v *VariableModifiedExecutor) SetChannel(channel spec.Channel) {
	v.channel = channel
}
