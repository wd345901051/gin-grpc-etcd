package middlerware

import (
	"api-gateway/pkg/e"
	"api-gateway/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = 200
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != 200 {
			ctx.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    e.GetMsg(uint(code)),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
