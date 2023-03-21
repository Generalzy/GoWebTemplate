package user

import (
	"errors"
	"github.com/Generalzy/GeneralSaaS/response"
	"github.com/Generalzy/GeneralSaaS/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	r := response.NewHttpResponse()
	username := ctx.Query("username")
	if username == "" {
		r.SetError(errors.New("username为必填字段")).SetCode(response.Code1).SetHttpStatus(http.StatusBadRequest).ReturnJson(ctx)
		return
	}

	t, err := token.GenerateToken(username)
	r.SetError(err).SetData(gin.H{"token": t}).ReturnJson(ctx)
}
