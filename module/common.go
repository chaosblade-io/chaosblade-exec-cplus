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
	"fmt"
	"strings"
)

func buildArgs(flags []string) string {
	args := ""
	for _, flag := range flags {
		if flag != "" {
			args = fmt.Sprintf(`%s %s`, args, flag)
		} else {
			args = fmt.Sprintf(`%s ''`, args)
		}
	}
	return strings.TrimSpace(args)
}

// containsGdbError checks if the result contains gdb execution errors
func containsGdbError(result string) bool {
	errorPatterns := []string{
		"couldn't execute \"gdb\"",
		"no such file or directory",
		"command not found",
		"gdb: not found",
		"spawn: command not found",
		"no symbol table is loaded",
		"no executable file now",
		"no debugging symbols found",
		"attach: no such file or directory",
		"could not attach to process",
		"permission denied",
		"operation not permitted",
		"signal: killed",
		"make breakpoint pending on future shared library load",
	}

	resultLower := strings.ToLower(result)
	for _, pattern := range errorPatterns {
		if strings.Contains(resultLower, strings.ToLower(pattern)) {
			return true
		}
	}
	return false
}
