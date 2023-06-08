package main

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"log"
	"net/http"
)

func main() {

	setting.Init()
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listen and Server on port", setting.HTTPPort)
	log.Fatal(s.ListenAndServe())

}
