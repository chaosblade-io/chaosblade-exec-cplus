package controller

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"

	"github.com/chaosblade-io/chaosblade-exec-cplus/common"
)

const DestroyName = "destroy"

type DestroyController struct {
}

func (d *DestroyController) GetControllerName() string {
	return DestroyName
}

func (d *DestroyController) GetRequestHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			fmt.Fprintf(writer,
				spec.ReturnFail(spec.Code[spec.IllegalParameters], err.Error()).Print())
			return
		}

		suid := request.Form.Get("suid")
		if suid == "" {
			fmt.Fprintf(writer,
				spec.ReturnFail(spec.Code[spec.IllegalParameters], "illegal suid parameter").Print())
			return
		}
		expModel := Manager.Experiments[suid]
		if expModel == nil {
			fmt.Fprintf(writer, spec.ReturnSuccess("the experiment not found").Print())
			return
		}
		processName := expModel.ActionFlags["processName"]
		// TODO remove? kill process?
		response := channel.NewLocalChannel().Run(context.Background(),
			path.Join(common.GetScriptPath(), common.RemoveProcessScript), processName)
		fmt.Fprintf(writer, response.Print())
	}
}
