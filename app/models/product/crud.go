package product

import (
	"goletter-go-api/pkg/model"
	"goletter-go-api/pkg/pagination"
	"goletter-go-api/pkg/route"
	"goletter-go-api/pkg/types"
	"net/http"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Product, error) {
	var product Product
	id := types.StringToInt(idstr)
	if err := model.DB.First(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}

// GetAll 获取全部商品
func GetAll(r *http.Request, perPage int) ([]Product, pagination.ViewData, error) {

	// 1. 初始化分页实例
	db := model.DB.Model(Product{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("products.index"), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var products []Product
	_pager.Results(&products)

	return products, viewData, nil
}