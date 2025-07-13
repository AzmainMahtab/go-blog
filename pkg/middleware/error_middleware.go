// Package middleware
package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/AzmainMahtab/go-blog/pkg/handler"
)

// ErrorHandler catches and formats unhandled errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			handlers.Error(c, c.Errors.Last())
		}
	}
}
