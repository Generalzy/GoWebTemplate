package validators

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
	"time"
)

const (
	DateValidatorKetWord = "checkDate"
)

// InitTrans 初始化翻译器
// 具体参考:https://blog.csdn.net/General_zy/article/details/127938889
func InitTrans(locale string) (ut.Translator, error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok := uni.GetTranslator(locale)
		if !ok {
			return nil, fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		var err error
		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}

		// 在校验器注册自定义的校验方法
		if err = v.RegisterValidation(DateValidatorKetWord, DateValidator); err != nil {
			return nil, err
		}
		// 注册自定义的校验字段
		if err = v.RegisterTranslation(
			DateValidatorKetWord,
			trans,
			RegisterTranslator(DateValidatorKetWord, "{0}必须要晚于当前日期"),
			Translate,
		); err != nil {
			return nil, err
		}

		return trans, err
	}
	return nil, errors.New("failed to initialize translator")
}

// DateValidator 日期校验器,校验传入time.Date是否晚于当前时间
func DateValidator(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}

// RegisterTranslator 为自定义字段添加翻译功能
func RegisterTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// Translate 自定义字段的翻译方法
func Translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
