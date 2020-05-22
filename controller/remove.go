package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
)

const RemoveName = "remove"

type RemoveController struct {
}

func (r *RemoveController) GetControllerName() string {
	return RemoveName
}

func (r *RemoveController) GetRequestHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO 暂时全部杀掉 gdb
		response := channel.NewLocalChannel().Run(context.Background(), "pkill", "-f gdb")
		fmt.Fprintf(writer, response.Print())
	}
}
