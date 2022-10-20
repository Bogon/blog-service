package dao

import (
	"app"
	"model"
)

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

// CountArticle 获取文章数目
func (d *Dao) CountArticle(title, desc, createdBy string, state uint8) (int, error) {
	// 创建model
	article := model.Article{
		Title: title,
		Desc:  desc,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return article.Count(d.engine)
}

// GetArticleList 获取文章列表
func (d *Dao) GetArticleList(title, desc, createdBy string, state uint8, page, pageSize int) ([]*model.Article, error) {
	// 创建Model
	article := model.Article{
		Title: title,
		Desc:  desc,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	pageoffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageoffset, pageSize)
}

// UpdateArticle 更新文章
func (d *Dao) UpdateArticle(id uint32, title, desc, content, coverImageURL, modifiedBy string, state uint8, isDel uint8) error {
	// 创建 model
	article := model.Article{
		Model: &model.Model{
			ID: id,
		},
	}

	values := map[string]interface{}{
		"state":  state,
		"is_del": isDel,
	}

	if modifiedBy != "" {
		values["modified_by"] = modifiedBy
	}

	if title != "" {
		values["title"] = title
	}

	if desc != "" {
		values["desc"] = desc
	}

	if coverImageURL != "" {
		values["cover_image_url"] = coverImageURL
	}

	if content != "" {
		values["content"] = content
	}

	return article.Update(d.engine, values)
}

// DeleteArticle 删除文章
func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}
