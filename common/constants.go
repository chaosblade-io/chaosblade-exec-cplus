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

package common

const (
	TargetName                = "cplus"
	DelayActionName           = "delay"
	ReturnErrorDataActionName = "return"
	VariableModifyActionName  = "modify"
	FileKeyword               = "file "
	SetEnvLdLibraryPath       = "set env LD_LIBRARY_PATH "
	SleepMilliseconds         = 1000
	CorePoolSize              = 1
	MaximumPoolSize           = 1
	KeepAliveTime             = 0
	BinName                   = "chaosblade-exec-cplus"
)

const (
	InitializationScript        = "shell_initialization.sh"
	RemoveProcessScript         = "shell_remove_process.sh"
	ResponseDelayScript         = "shell_response_delay.sh"
	ResponseDelayAttachScript   = "shell_response_delay_attach.sh"
	BreakAndReturnScript        = "shell_break_and_return.sh"
	BreakAndReturnAttachScript  = "shell_break_and_return_attach.sh"
	ModifyVariableScript        = "shell_modify_variable.sh"
	ModifyVariableAttachScript  = "shell_modify_variable_attach.sh"
	CheckProcessIdScript        = "shell_check_process_id.sh"
	CheckProcessDuplicateScript = "shell_check_process_duplicate.sh"
)
