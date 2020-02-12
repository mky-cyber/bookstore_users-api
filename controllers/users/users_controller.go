package users_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement get user")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement create user")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement search user")
}
