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

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/chaosblade-io/chaosblade-spec-go/util"
)

var proPath string

// GetProgramPath
func GetProgramPath() string {
	if proPath != "" {
		return proPath
	}
	dir, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatalf("can not get the process path, %v", err)
	}
	if p, err := os.Readlink(dir); err == nil {
		dir = p
	}
	proPath, err = filepath.Abs(filepath.Dir(dir))
	if err != nil {
		log.Fatalf("can not get the full process path, %v", err)
	}
	return proPath
}

var scriptPath string

func GetScriptPath() string {
	if scriptPath != "" {
		return scriptPath
	}
	scriptPath = path.Join(GetProgramPath(), "script")
	return scriptPath
}

var chaosbladeLogPath string

// chaosblade-VERSION/lib/cplus
func GetChaosBladeLogPath() string {
	if chaosbladeLogPath != "" {
		return chaosbladeLogPath
	}
	chaosbladeLogPath = path.Join(path.Dir(path.Dir(GetProgramPath())), "logs", "chaosblade.log")

	if util.IsExist(chaosbladeLogPath) {
		return chaosbladeLogPath
	}
	err := os.MkdirAll(path.Dir(chaosbladeLogPath), os.ModePerm)
	if err != nil {
		chaosbladeLogPath = path.Join(GetProgramPath(), "chaosblade.log")
	}
	return chaosbladeLogPath
}
