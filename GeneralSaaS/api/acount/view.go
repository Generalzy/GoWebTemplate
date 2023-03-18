package acount

import (
	"fmt"
	"github.com/Generalzy/GeneralSaaS/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfoList(ctx *gin.Context) {
	userInfos := make([]WebUserInfo, 0)

	var Form GetUserInfoListForm
	if err := ctx.ShouldBindJSON(&Form); err != nil {
		_ = ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  1,
			"error": err.Error(),
			"data":  userInfos,
		})
		return
	}
	fmt.Println(Form)
	fmt.Println(fmt.Sprintf("%s %s", Form.SortField, Form.SortOrder))

	err := global.GlobalMysqlClient.Table("web_userinfo").Order(
		fmt.Sprintf("%s %s", Form.SortField, Form.SortOrder)).Limit(Form.PageSize).Offset(Form.PageSize * (Form.PageIndex - 1)).Find(&userInfos).Error

	if err != nil {
		_ = ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  1,
			"error": err.Error(),
			"data":  userInfos,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": "",
			"data":  userInfos,
		})
	}
}
