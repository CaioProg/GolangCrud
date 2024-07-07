package controller

import (
	"fmt"

	"github.com/CaioProg/GolangCrud/src/configuration/rest_err"
	"github.com/CaioProg/GolangCrud/src/controller/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(

			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
