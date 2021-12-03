package router

import (
	"github.com/EricChiou/httprouter"
	"github.com/valyala/fasthttp"
)

func Init() {
}

func ListenAndServe(port string) error {
	return newFHServer().ListenAndServe(":" + port)
}

func ListenAndServeTLS(port, certPath, keyPath string) error {
	return newFHServer().ListenAndServeTLS(":"+port, certPath, keyPath)
}

func newFHServer() *fasthttp.Server {
	return &fasthttp.Server{
		Name:    "otter-v2",
		Handler: httprouter.FasthttpHandler(),
	}
}
