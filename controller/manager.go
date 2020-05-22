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
	"sync"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"

	"github.com/chaosblade-io/chaosblade-exec-cplus/module"
)

type ExpManager struct {
	// <uid, expModel>
	Experiments map[string]*spec.ExpModel
	// <actionName, actionModel>
	Actions map[string]*spec.ActionModel
}

var Manager *ExpManager
var once sync.Once

func init() {
	once.Do(func() {
		Manager = &ExpManager{
			Experiments: make(map[string]*spec.ExpModel, 0),
			Actions:     make(map[string]*spec.ActionModel, 0),
		}

		modelSpec := module.NewCPlusCommandModelSpec()
		actions := modelSpec.Actions()
		for _, action := range actions {
			actionModel := &spec.ActionModel{
				ActionName:      action.Name(),
				ActionAliases:   action.Aliases(),
				ActionShortDesc: action.ShortDesc(),
				ActionLongDesc:  action.LongDesc(),
				ActionMatchers: func() []spec.ExpFlag {
					matchers := make([]spec.ExpFlag, 0)
					for _, m := range action.Matchers() {
						matchers = append(matchers, spec.ExpFlag{
							Name:     m.FlagName(),
							Desc:     m.FlagDesc(),
							NoArgs:   m.FlagNoArgs(),
							Required: m.FlagRequired(),
						})
					}
					return matchers
				}(),
				ActionFlags: func() []spec.ExpFlag {
					flags := make([]spec.ExpFlag, 0)
					for _, m := range action.Flags() {
						flags = append(flags, spec.ExpFlag{
							Name:     m.FlagName(),
							Desc:     m.FlagDesc(),
							NoArgs:   m.FlagNoArgs(),
							Required: m.FlagRequired(),
						})
					}
					for _, m := range modelSpec.Flags() {
						flags = append(flags, spec.ExpFlag{
							Name:     m.FlagName(),
							Desc:     m.FlagDesc(),
							NoArgs:   m.FlagNoArgs(),
							Required: m.FlagRequired(),
						})
					}
					return flags
				}(),
			}
			actionModel.SetExecutor(action.Executor())
			Manager.Actions[action.Name()] = actionModel
		}
	})
}

func (e *ExpManager) Record(suid string, expModel *spec.ExpModel) error {
	// TODO
	e.Experiments[suid] = expModel
	return nil
}

func (e *ExpManager) Remove(suid string) error {
	// TODO
	delete(e.Experiments, suid)
	return nil
}
