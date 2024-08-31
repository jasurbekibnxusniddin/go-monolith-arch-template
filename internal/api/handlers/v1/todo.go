package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/models"
)

func (h *handler) CreateTodo(ctx *gin.Context) {

	var reqBody models.CreateTodoReq

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	todo, err := h.service.GetTodoService().Create(ctx, &reqBody)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, todo)
}
