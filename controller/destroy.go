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
	"path"
	"strings"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/sirupsen/logrus"

	"github.com/chaosblade-io/chaosblade-exec-cplus/common"
)

const DestroyName = "destroy"

type DestroyController struct {
}

func (d *DestroyController) GetControllerName() string {
	return DestroyName
}

func (d *DestroyController) GetRequestHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			fmt.Fprintf(writer,
				spec.ReturnFail(spec.Code[spec.IllegalParameters], err.Error()).Print())
			return
		}

		suid := request.Form.Get("suid")
		if suid == "" {
			fmt.Fprintf(writer,
				spec.ReturnFail(spec.Code[spec.IllegalParameters], "illegal suid parameter").Print())
			return
		}
		expModel := Manager.Experiments[suid]
		if expModel == nil {
			fmt.Fprintf(writer, spec.ReturnSuccess("the experiment not found").Print())
			return
		}
		processName := expModel.ActionFlags["processName"]
		if processName == "" {
			fmt.Fprintf(writer, spec.ReturnSuccess("success").Print())
			return
		}
		debug := expModel.ActionFlags["debug"] == "true"
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		localChannel := channel.NewLocalChannel()
		pids, err := localChannel.GetPidsByProcessName(processName, context.Background())
		if err == nil && len(pids) == 0 {
			fmt.Fprintf(writer, spec.ReturnSuccess("success").Print())
			return
		}
		var pid string
		if len(pids) > 0 {
			pid = strings.Join(pids, ",")
		}

		response := localChannel.Run(context.Background(),
			path.Join(common.GetScriptPath(), common.RemoveProcessScript), pid)
		fmt.Fprintf(writer, response.Print())
	}
}
