package controller

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada!")
	c.JSON(err.Code, err)
}
