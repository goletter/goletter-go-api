package article

import (
	"goletter-go-api/app/models"
	"goletter-go-api/app/models/user"
	"goletter-go-api/pkg/route"
	"goletter-go-api/pkg/types"
	"html/template"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title    string `gorm:"type:varchar(255);not null;" valid:"title"`
	Preview  string `gorm:"type:varchar(255);not null;" valid:"preview"`
	Describe string `gorm:"type:varchar(255);not null;" valid:"describe"`
	Content  string `gorm:"type:longtext;not null;" valid:"content"`
	Sort     string `gorm:"type:int(5);" valid:"sort"`
	Htmlcontent template.HTML

	UserID uint64 `gorm:"not null;index"`
	User   user.User

	ArticleTypeID uint64 `gorm:"not null;default:4;index"`
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

// GetStringID 获取 ID 的字符串格式
func (a Article) GetStringID() string {
	return types.Uint64ToString(a.ID)
}

// CreatedAtDate 创建日期
func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}