package main

import (
	"fmt"
	"otter-v2/config"
	"otter-v2/http/router"
	"otter-v2/service/jobqueue"
)

func main() {
	// init config
	config.Load("./config.ini")
	var cfg = config.Get()

	// init jobqueue
	jobqueue.Init()

	// set headers
	router.SetHeader("Access-Control-Allow-Origin", "*")
	router.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	router.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// init api
	router.Init()
	// start http server
	if err := router.ListenAndServe(cfg.ServerName, cfg.ServerPort); err != nil {
		panic(err)
	}
	// start https server
	// if err = router.ListenAndServeTLS("6300", cfg.SSLCertFilePath, cfg.SSLKeyFilePath); err != nil {
	// 	panic(err)
	// }

	// waiting for jobqueue finished
	jobqueue.Wait()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start server error: ", err)
		}
	}()
}
