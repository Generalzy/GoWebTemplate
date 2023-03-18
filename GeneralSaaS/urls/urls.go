package urls

import (
	"github.com/Generalzy/GeneralSaaS/GeneralSaaS/api/acount"
	"github.com/gin-gonic/gin"
)

func InitUrls(engine *gin.Engine) {
	apiV1 := engine.Group("/api/v1")

	countGroup := apiV1.Group("/acount")
	acount.Urls(countGroup)
}
