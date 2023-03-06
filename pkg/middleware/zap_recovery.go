package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lackone/go-ws/global"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

func ZapRecovery(recovery ...gin.RecoveryFunc) gin.HandlerFunc {
	var handle gin.RecoveryFunc

	if len(recovery) > 0 {
		handle = recovery[0]
	} else {
		handle = func(c *gin.Context, err any) {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") || strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if global.Logger != nil {
					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					headers := strings.Split(string(httpRequest), "\r\n")
					for idx, header := range headers {
						current := strings.Split(header, ":")
						if current[0] == "Authorization" {
							headers[idx] = current[0] + ": *"
						}
					}
					headersToStr := strings.Join(headers, "\r\n")
					if brokenPipe {
						global.Logger.Error(c.Request.URL.Path,
							zap.Any("error", err),
							zap.String("request", headersToStr),
						)
					} else {
						global.Logger.Error("[Recovery from panic]",
							zap.Any("error", err),
							zap.String("request", headersToStr),
							zap.String("stack", string(debug.Stack())),
						)
					}
				}
				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
				} else {
					handle(c, err)
				}
			}
		}()
		c.Next()
	}
}
