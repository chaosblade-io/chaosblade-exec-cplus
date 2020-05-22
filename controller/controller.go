package controller

import "net/http"

type Controller interface {
	GetControllerName() string
	GetRequestHandler() func(writer http.ResponseWriter, request *http.Request)
}

var Controllers = []Controller{
	&CreateController{},
	&DestroyController{},
	&RemoveController{},
	&StatusController{},
}
