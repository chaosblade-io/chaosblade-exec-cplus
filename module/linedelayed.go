package module

import (
	"context"
	"path"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/sirupsen/logrus"

	"github.com/chaosblade-io/chaosblade-exec-cplus/common"
)

type LineDelayedActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewLineDelayedActionSpec() spec.ExpActionCommandSpec {
	return &LineDelayedActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:     "delayDuration",
					Desc:     "delay time, unit is second",
					Required: true,
				},
			},
			ActionExecutor: &LineDelayedExecutor{},
		},
	}
}

func (l LineDelayedActionSpec) Name() string {
	return "delay"
}

func (l LineDelayedActionSpec) Aliases() []string {
	return []string{}
}

func (l LineDelayedActionSpec) ShortDesc() string {
	return "Code line delayed"
}

func (l LineDelayedActionSpec) LongDesc() string {
	return "Code line delayed"
}

type LineDelayedExecutor struct {
	channel spec.Channel
}

func (l *LineDelayedExecutor) Name() string {
	return "delay"
}

func (l *LineDelayedExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	delayDuration := model.ActionFlags["delayDuration"]
	if delayDuration == "" {
		return spec.ReturnFail(spec.Code[spec.IllegalParameters], "less necessary delayDuration value")
	}
	// search pid by process name
	processName := model.ActionFlags["processName"]
	if processName == "" {
		return spec.ReturnFail(spec.Code[spec.IllegalParameters], "less necessary processName value")
	}
	processCtx := context.WithValue(context.Background(), channel.ExcludeProcessKey, "blade")
	localChannel := channel.NewLocalChannel()
	pids, err := localChannel.GetPidsByProcessName(processName, processCtx)
	if err != nil {
		logrus.Warnf("get pids by %s process name err, %v", processName, err)
	}
	if pids == nil || len(pids) == 0 {
		args := buildArgs([]string{
			model.ActionFlags["fileLocateAndName"],
			model.ActionFlags["forkMode"],
			model.ActionFlags["libLoad"],
			model.ActionFlags["breakLine"],
			delayDuration,
			model.ActionFlags["initParams"],
		})
		return localChannel.Run(context.Background(), path.Join(common.GetScriptPath(), common.ResponseDelayScript), args)
	} else {
		args := buildArgs([]string{
			pids[0],
			model.ActionFlags["forkMode"],
			"",
			"",
			model.ActionFlags["breakLine"],
			delayDuration,
			model.ActionFlags["initParams"],
		})
		return localChannel.Run(context.Background(), path.Join(common.GetScriptPath(), common.ResponseDelayAttachScript), args)
	}
}

func (l *LineDelayedExecutor) SetChannel(channel spec.Channel) {
	l.channel = channel
}
