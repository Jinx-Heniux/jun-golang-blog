package main

import (
	"log"
	"net/http"

	"github.com/Jinx-Heniux/jun-golang-blog/common"
	"github.com/Jinx-Heniux/jun-golang-blog/router"
)

func init() {
	//模板加载
	common.LoadTemplate()
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// 路由
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
