package routes

import (
	"context"
	"net/http"

	"github.com/RaevNur/simple-go-microservices-api-gateway/api/crud/pb"
	"github.com/gin-gonic/gin"
)

type UpdateRequestBody struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func Update(ctx *gin.Context, c pb.CrudServiceClient) {
	reqBody := UpdateRequestBody{}

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Update(context.Background(), &pb.UpdateRequest{
		Id:     reqBody.Id,
		UserId: reqBody.UserId,
		Title:  reqBody.Title,
		Body:   reqBody.Body,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
