package v1

import "github.com/gin-gonic/gin"

func (h *handler) Ping(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"message": "pong"})
}
