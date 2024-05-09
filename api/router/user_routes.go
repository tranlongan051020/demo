package router

import (
	"demo/constant"
	"net/http"
	"strings"
)

func (r *Router) userRoutes() []routeType {
	return []routeType{
		{
			method:  http.MethodGet,
			path:    userPath,
			handler: r.allCtrl.UserCtrl.Get,
		},
		{
			method:  http.MethodPost,
			path:    strings.Join([]string{userPath, constant.PathParamKey.UserId}, "/"),
			handler: r.allCtrl.UserCtrl.Update,
		},
	}
}
