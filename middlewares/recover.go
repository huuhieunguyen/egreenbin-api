package middleware

import (
	"fmt"

	"github.com/GDSC-UIT/egreenbin-api/common"
	"github.com/GDSC-UIT/egreenbin-api/component"
	"github.com/gin-gonic/gin"
)

func Recover(sc component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				fmt.Print("AppError: ")
				if appErr, ok := err.(common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// panic(err)
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				// panic(err)
				return
			}
		}()

		c.Next()
	}
}
