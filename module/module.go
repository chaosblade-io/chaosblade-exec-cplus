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
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

type CPlusExpModuleSpec struct {
	spec.BaseExpModelCommandSpec
}

func NewCPlusCommandModelSpec() spec.ExpModelCommandSpec {
	return &CPlusExpModuleSpec{
		spec.BaseExpModelCommandSpec{
			ExpScope: "host",
			ExpActions: []spec.ExpActionCommandSpec{
				NewErrorReturnedActionSpec(),
				NewLineDelayedActionSpec(),
				NewVariableModifiedActionSpec(),
			},
			ExpFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:     "breakLine",
					Desc:     "Injection line in source code",
					Required: true,
				},
				&spec.ExpFlag{
					Name:     "fileLocateAndName",
					Desc:     "Startup file location and name",
					Required: true,
				},
				&spec.ExpFlag{
					Name:     "initParams",
					Desc:     "Initialization parameters for program startup (such as port number)",
					Required: true,
				},
				&spec.ExpFlag{
					Name:     "forkMode",
					Desc:     "Fault injection into child or parent processes (sub process:child ; main process:parent)",
					Required: true,
				},
				&spec.ExpFlag{
					Name:     "processName",
					Desc:     "Application process name",
					Required: true,
				},
				&spec.ExpFlag{
					Name:     "libLoad",
					Desc:     "If the class library needs to be loaded when the program starts, input the class library address",
					Required: false,
				},
			},
		},
	}
}

func (c *CPlusExpModuleSpec) Name() string {
	return "cplus"
}

func (c *CPlusExpModuleSpec) ShortDesc() string {
	return "C++ chaos experiments"
}

func (c *CPlusExpModuleSpec) LongDesc() string {
	return "C++ chaos experiments contain code line delayed, variable modified and err returned"
}

func (c *CPlusExpModuleSpec) Example() string {
	// TODO
	return ""
}
