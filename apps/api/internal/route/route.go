package route

import "github.com/gin-gonic/gin"

type Registerer interface {
	Register(router gin.IRouter)
}

func RegisterRoutes(router gin.IRouter, routes ...Registerer) {
	for _, route := range routes {
		route.Register(router)
	}
}
