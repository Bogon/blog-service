package service

import (
	"app"
	"model"
)

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=100"`
	Content       string `form:"content" binding:"required"`
	CoverImageUrl string `form:"cover_image_url"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state"`
}

type GetArticleRequest struct {
	ID uint32 `form:"id" binding:"required"`
}

type ArticleListRequest struct {
	Title     string `form:"title"`
	Desc      string `form:"desc"`
	CreatedBy string `form:"created_by"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CountArticleRequest struct {
	Title     string `form:"title" binding:"max=100"`
	Desc      string `form:"desc" binding:"max=100"`
	CreatedBy string `form:"created_by"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=100"`
	Content       string `form:"content" binding:"max=500"`
	ModifiedBy    string `form:"modified_by"`
	CoverImageUrl string `form:"cover_image_url"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	IsDel         uint8  `form:"is_del,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// CreateArticle 创建博客文章
func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.CreatedBy, param.State)
}

// GetArticle 获取文章
func (svc *Service) GetArticle(param *GetArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}

// CountArticle 博客数量
func (svc *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return svc.dao.CountArticle(param.Title, param.Desc, param.CreatedBy, param.State)
}

// GetArticleList 获取文章列表
func (svc *Service) GetArticleList(param *ArticleListRequest, paper *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.Title, param.Desc, param.CreatedBy, param.State, paper.Page, paper.PageSize)
}

// UpdateArticle 更新文章
func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.CoverImageUrl, param.ModifiedBy, param.State, param.IsDel)
}

// DeleteArticle 删除文章
func (svc Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
