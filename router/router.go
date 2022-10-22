package router

import (
	"net/http"

	"github.com/Jinx-Heniux/jun-golang-blog/views"
)

func Router() {

	// 页面 views

	http.HandleFunc("/", views.HTML.Index)

	// api json 数据

	// Error
	// The stylesheet http://127.0.0.1:8080/resource/css/index.css was not loaded because its MIME type, “text/html”, is not “text/css”.
	// 问题 resource/css/index.css文件没有被找到
	// 解决
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
