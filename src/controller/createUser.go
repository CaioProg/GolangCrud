package controller

import (
	"github.com/CaioProg/GolangCrud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("You called the wrong route")
	c.JSON(err.Code, err)
}
