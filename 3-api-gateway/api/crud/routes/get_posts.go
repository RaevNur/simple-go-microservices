package routes

import (
	"context"
	"net/http"

	"github.com/RaevNur/simple-go-microservices-api-gateway/api/crud/pb"
	"github.com/gin-gonic/gin"
)

type GetPostsRequestBody struct {
	Id []int64 `json:"id"`
}

func GetPosts(ctx *gin.Context, c pb.CrudServiceClient) {
	reqBody := GetPostsRequestBody{}

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.GetPosts(context.Background(), &pb.GetPostsRequest{
		Id: reqBody.Id,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
