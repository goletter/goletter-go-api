package controllers

import (
	"goletter-go-api/app/models/product"
	"goletter-go-api/pkg/route"
	"goletter-go-api/pkg/view"
	"html/template"
	"net/http"
)

// ArticlesController 处理静态页面
type ProductsController struct {
	BaseController
}

// Index 文章列表页
func (ac *ProductsController) Index(w http.ResponseWriter, r *http.Request) {
	// 1. 获取结果集
	products, pagerData, err := product.GetAll(r, 10)

	if err != nil {
		ac.ResposeForSQLError(w, err)
	} else {
		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Products":  products,
			"PagerData": pagerData,
		}, "products.index")
	}
}


func (ac *ProductsController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	product, err := product.Get(id)
	product.Htmlcontent = template.HTML(product.Content)

	// 3. 如果出现错误
	if err != nil {
		ac.ResposeForSQLError(w, err)
	} else {
		// ---  4. 读取成功，显示文章 ---
		view.Render(w, view.D{
			"Product": product,
		}, "products.show")
	}
}