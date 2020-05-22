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
			fmt.Fprintf(writer, spec.ReturnFail(spec.Code[spec.IllegalParameters], err.Error()).Print())
			return
		}
		actionModel := Manager.Actions[expModel.ActionName]
		if actionModel == nil {
			fmt.Fprintf(writer, spec.ReturnFail(spec.Code[spec.IllegalParameters], "action not supported").Print())
			return
		}
		// record
		err = Manager.Record(suid, expModel)
		if err != nil {
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
		return nil, "", fmt.Errorf("illegal suid parameter")
	}
	target := request.Form.Get("target")
	if target != common.TargetName {
		return nil, suid, fmt.Errorf("the target not support")
	}
	action := request.Form.Get("action")
	if action == "" {
		return nil, suid, fmt.Errorf("less action parameter")
	}
	breakLine := request.Form.Get("breakLine")
	if breakLine == "" {
		return nil, suid, fmt.Errorf("less breakLine parameter")
	}
	flags["breakLine"] = breakLine
	fileLocateAndName := request.Form.Get("fileLocateAndName")
	if fileLocateAndName == "" {
		return nil, suid, fmt.Errorf("less fileLocateAndName parameter")
	}
	flags["fileLocateAndName"] = fileLocateAndName
	forkMode := request.Form.Get("forkMode")
	if forkMode == "" {
		return nil, suid, fmt.Errorf("less forkMode parameter")
	}
	flags["forkMode"] = forkMode
	processName := request.Form.Get("processName")
	if processName == "" {
		return nil, suid, fmt.Errorf("less processName parameter")
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
