package service

import "model"

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

// CreateArticle 创建博客文章
func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.CreatedBy, param.State)
}

// GetArticle 获取文章
func (svc *Service) GetArticle(param *GetArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}
