package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*")
	router.GET("/sayHello/:name", SayHello)
	router.GET("/index", Index)
	router.Run(":9090")
}

//http://localhost:9090/sayHello/dong
func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, "hello,"+name)
}

//http://localhost:9090/index
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": fmt.Sprintf("%v,%v", "CreateByDongTech", time.Now().Local()),
	})
}
