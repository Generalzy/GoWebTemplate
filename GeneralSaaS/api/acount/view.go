package acount

import (
	"fmt"
	"github.com/Generalzy/GeneralSaaS/global"
	"github.com/Generalzy/GeneralSaaS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfoList(ctx *gin.Context) {
	// 获取响应对象
	response := utils.NewHttpResponse()

	// binding
	userInfos := make([]WebUserInfo, 0)
	var Form GetUserInfoListForm
	if err := ctx.ShouldBindJSON(&Form); err != nil {
		response.SetCode(utils.Code1).SetHttpStatus(http.StatusBadRequest).SetError(err).SetData(userInfos).ReturnJson(ctx)
		return
	}

	// SQL
	err := global.GlobalMysqlClient.Table("web_userinfo").Order(
		fmt.Sprintf("%s %s", Form.SortField, Form.SortOrder),
	).Limit(Form.PageSize).Offset(Form.PageSize * (Form.PageIndex - 1)).Find(&userInfos).Error
	if err != nil {
		response.SetCode(utils.Code1).SetHttpStatus(http.StatusBadRequest).SetData(userInfos).SetError(err).ReturnJson(ctx)
	} else {
		response.SetData(userInfos).SetData(userInfos).ReturnJson(ctx)
	}
}
