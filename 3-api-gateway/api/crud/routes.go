package crud

import (
	"github.com/RaevNur/simple-go-microservices-api-gateway/api/crud/routes"
	"github.com/RaevNur/simple-go-microservices-api-gateway/configs"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *configs.Config) {
	srv := &ServiceClient{
		Client: InitServiceClinet(c),
	}

	routes := r.Group("/post")
	routes.GET("/:id", srv.Get)
	routes.POST("/getposts", srv.GetPosts)
	routes.POST("/delete", srv.Delete)
	routes.POST("/update", srv.Update)
}

func (srv *ServiceClient) Get(ctx *gin.Context) {
	routes.Get(ctx, srv.Client)
}

func (srv *ServiceClient) GetPosts(ctx *gin.Context) {
	routes.GetPosts(ctx, srv.Client)
}

func (srv *ServiceClient) Delete(ctx *gin.Context) {
	routes.Delete(ctx, srv.Client)
}

func (srv *ServiceClient) Update(ctx *gin.Context) {
	routes.Update(ctx, srv.Client)
}
