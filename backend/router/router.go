/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-18
* Time: 上午11:18
* */
package router

import "github.com/gin-gonic/gin"

func RegisterRouter() *gin.Engine {
	engine := gin.New()

	router(engine)

	return engine
}

func router(app *gin.Engine) {

}
