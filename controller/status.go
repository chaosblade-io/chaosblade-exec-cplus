package controller

import (
	"fmt"
	"net/http"
)

const StatusName = "status"

type StatusController struct {
}

func (r *StatusController) GetControllerName() string {
	return StatusName
}

func (r *StatusController) GetRequestHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "success")
	}
}
