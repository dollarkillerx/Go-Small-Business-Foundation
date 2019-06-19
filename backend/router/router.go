/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 上午10:11
* */
package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("./view/static"))

	return router
}
