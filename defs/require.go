/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-16
* Time: 上午9:47
* */
package defs

import "github.com/gin-gonic/gin"

type RequestDate struct {
	Code int
	Data interface{}
}

var (
	ErrorBadRequest = &RequestDate{Code: 400, Data: &gin.H{"Code": "40001", "Msg": "Bad required"}}
	ErrorSQLNotData = &RequestDate{Code: 404, Data: &gin.H{"Code": "40401", "Msg": "SQL Not Found"}}
	ErrorBadServer  = &RequestDate{Code: 500, Data: &gin.H{"Code": "50001", "Msg": "Internal Server Error"}}
)
