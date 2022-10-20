package v1

import (
	"app"
	"convert"
	"errcode"
	"github.com/gin-gonic/gin"
	"global"
	"service"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// Get
// @Summary 获取文章
// @Produce  json
// @Param id query uint32 1 "文章id"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/:id [get]
func (a *Article) Get(c *gin.Context) {
	// 1. 获取参数
	param := service.GetArticleRequest{convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	// 2. 绑定参数验证参数
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "srv.Get error:", errors)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 3. 查询数据库
	svc := service.New(c)
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArtivle error:", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}
	// 4. 返回结果
	response.ToResponse(article)
	return
}

// List
// @Summary 获取文章列表
// @Produce  json
// @Param title query string false "标签名称" maxlength(100)
// @Param desc query string false "描述" maxlength(100)
// @Param content query string false "文章内容" maxlength(100)
// @Param coverImageURL query string false "文章图片" maxlength(100)
// @Param CreatedBy query string false "文章创建者" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param isDel query int false "是否删除" Enums(0, 1) default(0)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a *Article) List(c *gin.Context) {
	// 1. 获取参数
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	// 2. 绑定参数、校验参数
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "svc.List error:", errors)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	//3. 查询数据库
	svc := service.New(c)
	// 3.1 查询记录符合条件的总条数
	totalRows, err := svc.CountArticle(&service.CountArticleRequest{
		Title:     param.Title,
		Desc:      param.Desc,
		State:     param.State,
		CreatedBy: param.CreatedBy,
	})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountArticle error:", err)
		response.ToErrorResponse(errcode.ErrorCountArticleFail)
		return
	}
	// 3.2 查询所有的符合条件的记录
	paper := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articleList, err := svc.GetArticleList(&param, &paper)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticleList error:", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}
	// 5.返回结果
	response.ToResponseList(articleList, totalRows)
}

// Create
// @Summary 创建文章
// @Produce  json
// @Param title query string false "标签名称" maxlength(100)
// @Param desc query string false "描述" maxlength(100)
// @Param content query string false "文章内容" maxlength(100)
// @Param coverImageURL query string false "文章图片" maxlength(100)
// @Param createdBy query string false "文章创建者" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a *Article) Create(c *gin.Context) {
	// 1.获取传入的参数,绑定参数
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	// 2.检验参数
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		// 2.1 打错误日志
		global.Logger.Errorf(c, "src.CreateArticle error:", errors)
		// 2.2 返回
		app.NewResponse(c).ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 3.插入数据库
	// 3.1获取服务句柄
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateArticle error:", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	// 4.返回结果
	response.ToResponse(gin.H{})
}

// Update
// @Summary 更新文章
// @Produce  json
// @Param title query string false "标签名称" maxlength(100)
// @Param desc query string false "描述" maxlength(100)
// @Param content query string false "文章内容" maxlength(100)
// @Param coverImageURL query string false "文章图片" maxlength(100)
// @Param createdBy query string false "文章创建者" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (a *Article) Update(c *gin.Context) {
	// 1.获取参数
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	// 2.绑定并校验参数
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "svc.Update error:", errors)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 3.提交更改到数据库
	svc := service.New(c)
	if err := svc.UpdateArticle(&param); err != nil {
		global.Logger.Errorf(c, "svc.UpdateArticle error:", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	// 4.返回修改结果
	response.ToResponse(nil)
	return
}

// Delete
// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a *Article) Delete(c *gin.Context) {
	// 1.获取参数
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	// 2. 检验参数
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "svc.Delete error:", errors)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 3.从数据库删除
	svc := service.New(c)
	if err := svc.DeleteArticle(&param); err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticle error:", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	// 4.返回结果
	response.ToResponse(nil)
	return
}
