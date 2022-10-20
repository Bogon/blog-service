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
func (a *Article) List(c *gin.Context) {
	// 获取文章列表

}

// Create 创建文章
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

func (a *Article) Update(c *gin.Context) {}
func (a *Article) Delete(c *gin.Context) {}
