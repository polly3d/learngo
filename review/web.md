# go web开发总结

## 路由

```
r := gin.Default()
r.GET("/", func(c *gin.Context) {
	c.String(200, "hello world")
})
```

## 数据库

```
type tableName struct {
	filedName string
}

db, err := gorm.Open("mysql", "username:password@/database?charset=utf8&parseTime=True")

var list []tableName
db.Find(&list)
```

## 模板渲染

```
r := gin.Default()
r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Title from go",
		})
	})
```

模板内容
```
<h1>{{ .title }}</h1>
```