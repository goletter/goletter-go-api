package product

import (
	"goletter-go-api/app/models"
	"goletter-go-api/pkg/route"
	"goletter-go-api/pkg/types"
	"html/template"
)

type Product struct {
	models.BaseModel

	Title 	string `gorm:"type:varchar(255);not null;" valid:"title"`
	Preview string `gorm:"type:varchar(255);not null;" valid:"preview"`
	Code    string `gorm:"type:varchar(100);not null;" valid:"code"`
	Content string `gorm:"type:varchar(255);not null;" valid:"content"`
	Htmlcontent template.HTML
}

// Link 方法用来生成文章链接
func (a Product) Link() string {
	return route.Name2URL("products.show", "id", a.GetStringID())
}

// GetStringID 获取 ID 的字符串格式
func (a Product) GetStringID() string {
	return types.Uint64ToString(a.ID)
}

// CreatedAtDate 创建日期
func (a Product) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}