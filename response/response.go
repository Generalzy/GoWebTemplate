package response

import (
	"fmt"
	"github.com/Generalzy/GeneralSaaS/global"
	"github.com/Generalzy/GeneralSaaS/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const (
	Code0 = iota
	Code1
)

type HttpResponse struct {
	code       int
	data       interface{}
	error      interface{}
	httpStatus int
}

func NewHttpResponse() *HttpResponse {
	return &HttpResponse{
		code:       Code0,
		data:       nil,
		error:      nil,
		httpStatus: http.StatusOK,
	}
}

func (response *HttpResponse) SetCode(code int) *HttpResponse {
	response.code = code
	return response
}

func (response *HttpResponse) SetHttpStatus(status int) *HttpResponse {
	response.httpStatus = status
	return response
}

func (response *HttpResponse) SetData(data interface{}) *HttpResponse {
	response.data = data
	return response
}

func (response *HttpResponse) SetError(err interface{}) *HttpResponse {
	response.error = err
	return response
}

func (response *HttpResponse) ReturnJson(ctx *gin.Context) {
	if response.error != nil {
		// 判断错误是否是校验器引起的
		if errs, ok := response.error.(validator.ValidationErrors); ok {
			response.error = utils.TrimMapStructNamePrefix(errs.Translate(global.GlobalTranslator))
		} else {
			// 如果error允许string则将error转化为string
			if errs, ok := response.error.(error); ok {
				response.error = errs.Error()
			}
		}

		// 将error的形式放入ctx
		_ = ctx.Error(fmt.Errorf("%v", response.error))
	}

	ctx.JSON(response.httpStatus, gin.H{
		"code":  response.code,
		"data":  response.data,
		"error": response.error,
	})
}
