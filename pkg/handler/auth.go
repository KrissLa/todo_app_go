package handler

import (
	todo "github.com/KrissLa/todo_app_go"
	"github.com/gin-gonic/gin"
	"net/http"
)



func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})


}

func (h *Handler) signIn(c *gin.Context) {

}
