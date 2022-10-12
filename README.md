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