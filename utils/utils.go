package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

// TrimMapStructNamePrefixAndToString 移除map中结构体名称前缀并返回字符串
func TrimMapStructNamePrefixAndToString(errorMaps validator.ValidationErrorsTranslations) string {
	builder := strings.Builder{}
	for field, err := range errorMaps {
		builder.WriteString(fmt.Sprintf("%s:%s \n", field[strings.Index(field, ".")+1:], err))
	}
	return builder.String()
}

// TrimMapStructNamePrefix 移除map中结构体名称前缀
func TrimMapStructNamePrefix(errorMaps validator.ValidationErrorsTranslations) map[string]string {
	result := make(map[string]string, len(errorMaps))
	for field, err := range errorMaps {
		result[field[strings.Index(field, ".")+1:]] = err
	}
	return result
}
