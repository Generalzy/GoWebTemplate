package acount

import "github.com/gin-gonic/gin"

func Urls(engine *gin.RouterGroup) {
	engine.GET("get_user_info_list", GetUserInfoList)
}
