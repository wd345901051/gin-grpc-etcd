package middlerware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("123")
	}
}
