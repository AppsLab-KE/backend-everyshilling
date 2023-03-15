package router

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define routes

	r.POST("/users", c.createUser)
	r.GET("/login", c.loginUser)
	return r
}

func (c *Controller) createUser(context *gin.Context) {

}

func (c *Controller) loginUser(context *gin.Context) {

}
