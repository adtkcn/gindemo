package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gindemo/controller"
)

//go:embed static
var static embed.FS

func main() {

	// dsn := "xx:FMKmpFkKD2GMkcjc@tcp(192.168.137.3:3306)/xx?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "heng_", // 表名前缀，`User` 的表名应该是 `t_users`
	// 		SingularTable: true,    // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
	// 	},
	// })

	// // db, err := gorm.Open("mysql", "root:xiangheng@/meddit?charset=utf8&parseTime=True&loc=Local")

	// if err != nil {
	// 	log.Println("数据库连接错误", err)
	// 	return
	// }
	// db.AutoMigrate(&model.User{}) //自动迁移仅仅会创建表，缺少列和索引

	// router:= gin.New()
	router := gin.Default()
	router.RedirectTrailingSlash = false //自动重定向

	// router.Static("/static", "./static") //静态文件

	router.POST("/login", controller.Login)

	router.GET("/reg", controller.Reg)
	router.GET("/checklogin", controller.CheckLogin)

	router.POST("/upload", controller.Upload)

	/**
	生成路径
	/public/static/
	*/
	router.StaticFS("/public", http.FS(static)) //静态文件 嵌入程序中

	fmt.Println("地址： http://127.0.0.1:8080")

	// util.GenRsaKey(2048)

	router.Run()

}
