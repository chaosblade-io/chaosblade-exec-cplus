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

package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
	"github.com/sirupsen/logrus"

	"github.com/chaosblade-io/chaosblade-exec-cplus/common"
)

const CreateName = "create"

type CreateController struct {
}

func (c *CreateController) GetControllerName() string {
	return CreateName
}

func (c *CreateController) GetRequestHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		expModel, suid, err := convertRequestToExpModel(request)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
			return
		}
		actionModel := Manager.Actions[expModel.ActionName]
		if actionModel == nil {
			util.Errorf(suid, util.GetRunFuncName(), fmt.Sprintf(spec.ResponseErr[spec.CplusActionNotSupport].ErrInfo, expModel.ActionName))
			response := spec.ResponseFailWaitResult(spec.CplusActionNotSupport, fmt.Sprintf(spec.ResponseErr[spec.CplusActionNotSupport].Err, expModel.ActionName),
				fmt.Sprintf(spec.ResponseErr[spec.CplusActionNotSupport].ErrInfo, expModel.ActionName))
			fmt.Fprintf(writer, response.Print())
			return
		}
		// record
		err = Manager.Record(suid, expModel)
		if err != nil {
			// todo : need not edit, because Manager.Record alawys return nil, now!
			util.Errorf(suid, util.GetRunFuncName(), "the experiment exists")
			fmt.Fprintf(writer, spec.ReturnFail(spec.Code[spec.IllegalParameters], "the experiment exists").Print())
			return
		}
		// TODO 开启 debug
		logrus.SetLevel(logrus.DebugLevel)
		response := actionModel.Executor().Exec(suid, context.Background(), expModel)
		if !response.Success {
			Manager.Remove(suid)
		}
		fmt.Fprintf(writer, response.Print())
	}
}

func convertRequestToExpModel(request *http.Request) (*spec.ExpModel, string, error) {
	err := request.ParseForm()
	if err != nil {
		return nil, "", err
	}
	flags := make(map[string]string, 0)

	suid := request.Form.Get("suid")
	if suid == "" {
		util.Errorf(suid, util.GetRunFuncName(), fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "suid"))
		return nil, "", spec.ResponseFailWaitResult(spec.ParameterLess, fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].Err, "suid"),
			fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "suid"))
	}
	target := request.Form.Get("target")
	if target != common.TargetName {
		util.Errorf(suid, util.GetRunFuncName(), fmt.Sprintf(spec.ResponseErr[spec.ParameterInvalidCplusTarget].ErrInfo, target))
		return nil, suid, spec.ResponseFailWaitResult(spec.ParameterInvalidCplusTarget, fmt.Sprintf(spec.ResponseErr[spec.ParameterInvalidCplusTarget].Err, target),
			fmt.Sprintf(spec.ResponseErr[spec.ParameterInvalidCplusTarget].ErrInfo, target))
	}
	action := request.Form.Get("action")
	if action == "" {
		util.Errorf(suid, util.GetRunFuncName(), fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "action"))
		return nil, suid, spec.ResponseFailWaitResult(spec.ParameterLess, fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].Err, "action"),
			fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "action"))
	}
	breakLine := request.Form.Get("breakLine")
	if breakLine == "" {
		util.Errorf(suid, util.GetRunFuncName(), fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "breakLine"))
		return nil, suid, spec.ResponseFailWaitResult(spec.ParameterLess, fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].Err, "breakLine"),
			fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "breakLine"))
	}
	flags["breakLine"] = breakLine
	fileLocateAndName := request.Form.Get("fileLocateAndName")
	if fileLocateAndName == "" {
		util.Errorf(suid, util.GetRunFuncName(), fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "fileLocateAndName"))
		return nil, suid, spec.ResponseFailWaitResult(spec.ParameterLess, fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].Err, "fileLocateAndName"),
			fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "fileLocateAndName"))
	}
	flags["fileLocateAndName"] = fileLocateAndName
	forkMode := request.Form.Get("forkMode")
	if forkMode == "" {
		util.Errorf(suid, util.GetRunFuncName(), fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "forkMode"))
		return nil, suid, spec.ResponseFailWaitResult(spec.ParameterLess, fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].Err, "forkMode"),
			fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "forkMode"))
	}
	flags["forkMode"] = forkMode
	processName := request.Form.Get("processName")
	if processName == "" {
		util.Errorf(suid, util.GetRunFuncName(), fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "processName"))
		return nil, suid, spec.ResponseFailWaitResult(spec.ParameterLess, fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].Err, "processName"),
			fmt.Sprintf(spec.ResponseErr[spec.ParameterLess].ErrInfo, "processName"))
	}
	flags["processName"] = processName
	libLoad := request.Form.Get("libLoad")
	if libLoad != "" {
		libLoad = common.SetEnvLdLibraryPath + libLoad
	}
	flags["libLoad"] = libLoad
	initParams := request.Form.Get("initParams")
	if initParams != "" {
		flags["initParams"] = initParams
	}

	// TODO delay
	delayDuration := request.Form.Get("delayDuration")
	if delayDuration != "" {
		flags["delayDuration"] = delayDuration
	}
	// return
	returnValue := request.Form.Get("returnValue")
	if returnValue != "" {
		flags["returnValue"] = returnValue
	}

	// modify
	variableValue := request.Form.Get("variableValue")
	if variableValue != "" {
		flags["variableValue"] = variableValue
	}
	variableName := request.Form.Get("variableName")
	if variableName != "" {
		flags["variableName"] = variableName
	}

	return &spec.ExpModel{
		Target:      common.TargetName,
		Scope:       "",
		ActionName:  action,
		ActionFlags: flags,
	}, suid, nil
}
