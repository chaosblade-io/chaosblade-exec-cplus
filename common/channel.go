package common

import (
	"context"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

type AsyncChannel struct {
	localChannel channel.OsChannel
}

func NewAsyncChannel() spec.Channel {
	return &AsyncChannel{
		localChannel: channel.NewLocalChannel(),
	}
}

func (a *AsyncChannel) Run(ctx context.Context, script, args string) *spec.Response {
	go a.localChannel.Run(ctx, script, args)
	// less judgement
	return spec.ReturnSuccess("success")
}

func (a *AsyncChannel) GetScriptPath() string {
	return a.localChannel.GetScriptPath()
}
