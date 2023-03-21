package urls

import (
	"github.com/Generalzy/GeneralSaaS/web/api/user"
	"github.com/gin-gonic/gin"
)

func InitUrls(engine *gin.Engine) {
	apiV1 := engine.Group("/api/v1")
	{
		userGroup := apiV1.Group("/user")
		user.Urls(userGroup)
	}
}
