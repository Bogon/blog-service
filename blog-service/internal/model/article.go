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
func (a *Article) Article(db *gorm.DB) (*Article, error) {
	// 创建 临时变量
	var article Article
	// 按照条件查询
	if err := db.Where("id = ?", a.ID).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// Count 获取数字
func (a *Article) Count(db *gorm.DB) (int, error) {
	// 创建临时变量
	var count int
	// 构造查询参数
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	if a.Desc != "" {
		db = db.Where("desc = ?", a.Desc)
	}

	if a.CreatedBy != "" {
		db = db.Where("created_by = ?", a.CreatedBy)
	}
	db = db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	// 返回查询结果
	return count, nil
}

// List 获取文章列表
func (a *Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	// 创建临时存储列表变量
	var articleList []*Article
	var err error
	if pageOffset > 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	// 构造查询参数
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	if a.Desc != "" {
		db = db.Where("desc = ?", a.Desc)
	}

	if a.CreatedBy != "" {
		db = db.Where("created_by = ?", a.CreatedBy)
	}
	db = db.Where("state = ?", a.State)
	// 查询结果
	if err = db.Where("is_del=?", 0).Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}

// Update 更新文章
func (a *Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(a).Where("id = ?", a.ID).Update(values).Error; err != nil {
		return err
	}
	return nil
}
