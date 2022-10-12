# blog-service

## 生成文档 Swagger
| 注解       | 描述                                                            |
|----------|---------------------------------------------------------------|
| @Summary | 摘要                                                            |
| @Produce | API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等 |
 | @Param   | 参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释                            |
| @Success |响应成功，从左到右分别为：状态码、参数类型、数据类型、注释|
|@Failure|响应失败，从左到右分别为：状态码、参数类型、数据类型、注释|
|@Router|路由，从左到右分别为：路由地址，HTTP 方法|


## 国际化/接口参数校验 `go-playground/validator `
安装：
```Bash
go get -u github.com/go-playground/validator/v10
```

 | 标签       | 描述   |
|------|---------|
 | required | 必填   |
| gt       | 大于   |
| gte      | 大于等于 |
 | lt       | 小于   |
 | lte      | 小于等于   |
|min |	最小值|
|max|	最大值|
|oneof|	参数集内的其中之一|
|len|	长度要求与 len 给定的一致|

## 上传图片
```Go
UploadSavePath：上传文件的最终保存目录。
UploadServerUrl：上传文件后的用于展示的文件服务地址。
UploadImageMaxSize：上传文件所允许的最大空间大小（MB）。
UploadImageAllowExts：上传文件所允许的文件后缀。
```

> 在实际项目中，应当将应用服务和文件服务给拆分开来，因为从安全角度来讲，文件资源不应当与应用资源摆放在一起，否则会有风险，又或是直接采用市面上的 OSS 也是可以的。