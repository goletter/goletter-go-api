package admin

import (
	"fmt"
	"goletter-go-api/app/models/product"
	"net/http"
)

// ProductController 处理静态页面
type ProductController struct {
	BaseController
}

// Index 文章列表页
func (ac *ProductController) Index(w http.ResponseWriter, r *http.Request) {

	// 1. 获取结果集
	articles, pagerData, err := product.GetAll(r, 10)

	if err != nil {
		ac.ResposeForSQLError(w, err)
	} else {
		fmt.Println(articles)
		fmt.Println(pagerData)
	}
}
