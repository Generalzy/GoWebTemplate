package middleware

import (
	"errors"
	"fmt"
	"github.com/Generalzy/GeneralSaaS/global"
	"github.com/Generalzy/GeneralSaaS/response"
	"github.com/Generalzy/GeneralSaaS/utils/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// 跳过jwt的url
var allowedUrl = map[string]struct{}{}

func InitMiddleWare(engine *gin.Engine) {
	engine.Use(GinLogger(), GinRecovery(true), cors.Default(), JsonWebTokenMiddleWare())
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		global.GlobalLogger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			// UA
			// zap.String("user-agent", c.Request.UserAgent()),
			// ctx中的报错信息
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			// 请求响应时间
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.GlobalLogger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					global.GlobalLogger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					global.GlobalLogger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func JsonWebTokenMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := allowedUrl[ctx.Request.URL.Path]; ok {
			ctx.Next()
			return
		}

		r := response.NewHttpResponse()
		t := ctx.GetHeader("Authorization")

		if t == "" {
			r.SetCode(response.Code1).SetHttpStatus(http.StatusBadRequest).SetError(
				errors.New("auth为空")).ReturnJson(ctx)
			ctx.Abort()
			return
		}

		parts := strings.SplitN(t, " ", 2)
		if len(parts) != 2 && parts[0] != "Jwt" {
			r.SetCode(response.Code1).SetHttpStatus(http.StatusBadRequest).SetError(
				errors.New("auth错误")).ReturnJson(ctx)
			ctx.Abort()
			return
		}

		claim, err := token.ParseToken(parts[1])
		if err != nil {
			r.SetCode(response.Code1).SetHttpStatus(http.StatusBadRequest).SetError(
				errors.New("token错误")).ReturnJson(ctx)
			ctx.Abort()
			return
		}

		fmt.Println(claim.UID)
		ctx.Next()
	}
}
