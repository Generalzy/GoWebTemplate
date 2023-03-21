package user

import "github.com/gin-gonic/gin"

func Urls(group *gin.RouterGroup) {
	group.GET("/login", Login)
}
