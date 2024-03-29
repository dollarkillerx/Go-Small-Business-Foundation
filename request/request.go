/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-16
* Time: 上午9:48
* */
package request

import (
	"PRE-Cartoon/defs"
	"github.com/gin-gonic/gin"
)

func RepMsg(ctx *gin.Context, data *defs.RequestDate) {
	ctx.JSON(data.Code, data.Data)
}

func RepData(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, &gin.H{
		"Code": "200",
		"Data": data,
	})
}

func RepBaseData(ctx *gin.Context, data string) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(200)
	ctx.Writer.WriteString(data)
}
