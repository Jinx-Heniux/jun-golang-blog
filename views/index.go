package views

import (
	"errors"
	"log"
	"net/http"

	"github.com/Jinx-Heniux/jun-golang-blog/common"
	"github.com/Jinx-Heniux/jun-golang-blog/service"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	//数据库查询
	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	index.WriteData(w, hr)
}
