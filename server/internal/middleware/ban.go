package middleware

import (
	"github.com/gin-gonic/gin"
	"interastral-peace.com/alnitak/internal/resp"
)

func Ban() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.GetInt("status")
		if status == 1 {
			resp.Result(ctx, 500, nil, "用户封禁中")
			ctx.Abort()
			return
		}
	}
}
