package routes

import (
	"context"
	"net/http"

	"github.com/RaevNur/simple-go-microservices-api-gateway/api/parser/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ParseStatus(ctx *gin.Context, c pb.ParserServiceClient) {
	res, err := c.ParseStatus(context.Background(), &emptypb.Empty{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
