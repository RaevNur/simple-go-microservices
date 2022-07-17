package parser

import (
	"github.com/RaevNur/simple-go-microservices-api-gateway/api/parser/routes"
	"github.com/RaevNur/simple-go-microservices-api-gateway/configs"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *configs.Config) {
	srv := &ServiceClient{
		Client: InitServiceClinet(c),
	}

	routes := r.Group("/parser")
	routes.POST("/parse", srv.ParsePosts)
	routes.GET("/status", srv.ParseStatus)
}

func (srv *ServiceClient) ParsePosts(ctx *gin.Context) {
	routes.ParsePosts(ctx, srv.Client)
}

func (srv *ServiceClient) ParseStatus(ctx *gin.Context) {
	routes.ParseStatus(ctx, srv.Client)
}
