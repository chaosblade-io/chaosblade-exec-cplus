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
