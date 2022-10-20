package model

import (
	"app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

// Create 插入数据库
func (a *Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

// Article 获取数据
func (a Article) Article(db *gorm.DB) (*Article, error) {
	// 创建 临时变量
	var article Article
	// 按照条件查询
	if err := db.Where("id = ?", a.ID).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
