package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 配置模板
	router.LoadHTMLGlob("resources/templates/*")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	router.StaticFS("/static", http.Dir("resources/static"))
	// 请求
	group := router.Group("/")
	{
		group.GET("/", Index)
		group.GET("/index", Index)
		group.GET("/hello", Hello)
		group.GET("/sayHello/:name", SayHello)
	}
	router.Run(":9090")
}

//http://localhost:9090/sayHello/dong
func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, "hello,"+name)
}

//http://localhost:9090/hello
func Hello(c *gin.Context) {
	c.HTML(http.StatusOK, "hello.html", gin.H{
		"Hello": fmt.Sprintf("%v	%v", "HelloWorld!", time.Now().Local()),
	})
}

//http://localhost:9090/index
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		`WEBSITE_TITLE`: `東の博客`,
	})
}
