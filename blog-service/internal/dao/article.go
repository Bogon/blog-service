package dao

import "model"

// CreateArticle 创建文章
func (d *Dao) CreateArticle(title, desc, content, coverImageURL, createdBy string, state uint8) error {
	// 1、创建model
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageURL,
		State:         state,
		Model:         &model.Model{CreatedBy: createdBy},
	}
	// 使用gorm插入数据
	return article.Create(d.engine)
}

// GetArticle 获取文章
func (d *Dao) GetArticle(id uint32) (*model.Article, error) {
	// 创建查询model
	article := model.Article{
		Model: &model.Model{
			ID: id,
		},
	}
	return article.Article(d.engine)
}
