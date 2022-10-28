package views

import (
	"net/http"

	"github.com/Jinx-Heniux/jun-golang-blog/common"
	"github.com/Jinx-Heniux/jun-golang-blog/config"
	"github.com/Jinx-Heniux/jun-golang-blog/models"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}

	/*
		var hr = &models.HomeResponse{
			config.Cfg.Viewer,
			categorys,
			posts,
			1,
			1,
			[]int{1},
			true,
		}
	*/

	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}

	index.WriteData(w, hr)
}
