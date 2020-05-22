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
