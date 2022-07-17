package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/RaevNur/simple-go-microservices-api-gateway/api/crud/pb"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context, c pb.CrudServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.Get(context.Background(), &pb.GetRequest{
		Id: id,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
