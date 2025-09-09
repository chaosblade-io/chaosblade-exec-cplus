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

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
)

const RemoveName = "remove"

type RemoveController struct {
}

func (r *RemoveController) GetControllerName() string {
	return RemoveName
}

func (r *RemoveController) GetRequestHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO 暂时全部杀掉 gdb
		// 使用 shell 命令组合，确保即使没有匹配的进程也返回成功
		response := channel.NewLocalChannel().Run(context.Background(), "sh", "-c 'pkill -f gdb || true'")
		fmt.Fprintf(writer, response.Print())
	}
}
