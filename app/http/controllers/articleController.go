package controllers

import (
	"goletter-go-api/app/models/article"
	"goletter-go-api/pkg/route"
	"goletter-go-api/pkg/view"
	"html/template"
	"net/http"
)

// ArticlesController 处理静态页面
type ArticlesController struct {
	BaseController
}

// Index 文章列表页
func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	// 1. 获取结果集
	articles, pagerData, err := article.GetAll(r, 10)

	if err != nil {
		ac.ResposeForSQLError(w, err)
	} else {
		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index")
	}
}


func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := article.Get(id)
	article.Htmlcontent = template.HTML(article.Content)

	// 3. 如果出现错误
	if err != nil {
		ac.ResposeForSQLError(w, err)
	} else {
		// ---  4. 读取成功，显示文章 ---
		view.Render(w, view.D{
			"Article": article,
		}, "articles.show")
	}
}