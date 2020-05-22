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
