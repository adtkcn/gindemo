package main

import (
	"embed"
	"fmt"

	"github.com/gin-gonic/gin"

	"gindemo/controller"
)

//go:embed static/**
var static embed.FS

func main() {

	// router:= gin.New()
	router := gin.Default()
	router.RedirectTrailingSlash = false //自动重定向

	// router.Static("/static", "./static") //静态文件

	router.POST("/login", controller.Login)
	router.POST("/upload", controller.Upload)

	router.GET("/reg", controller.Reg)
	router.GET("/checklogin", controller.CheckLogin)

	router.GET("/cookie", func(c *gin.Context) {
		fmt.Printf("%#v \n", c.Request)
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", c.Request.Host, false, true)
		}
		// c.SetCookie("gin_cookie", "test", 3600, "/", c.Request.Host, false, true)

		fmt.Printf("Cookie value: %s , %s\n", cookie, err)
	})

	/**
	生成路径
	/public/static/
	*/
	// router.StaticFS("/public", http.FS(static)) //静态文件 嵌入程序中
	// router.StaticFS("/public", static.ReadDir("static")) //静态文件 嵌入程序中
	// router.StaticFS("/", http.FS(static))

	router.StaticFile("/", "./static/index.html")
	router.Static("/static", "./static")

	IP_PORT := ":10080"
	fmt.Println("地址： http://127.0.0.1" + IP_PORT)
	// util.GenRsaKey(2048)
	router.Run(IP_PORT)

}
