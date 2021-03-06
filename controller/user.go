package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"gindemo/model"
)

func Login(ctx *gin.Context) {

	var user model.User

	// 获取文件
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err)
		// return
	} else {
		log.Println(file.Filename)
		// 保存文件
		ctx.SaveUploadedFile(file, "./upload/"+file.Filename)
	}

	if ctx.ShouldBind(&user) == nil {
		// fmt.Println("待插入数据", user)
		// u := model.User{Name: "Jinzhu", Pwd: "18"}
		// fmt.Println("待插入数据", u)
		if err := model.DB.Create(&user).Error; err != nil {
			fmt.Println("插入失败", err)
			ctx.JSON(200, user)

			return
		}

		ctx.JSON(200, user)
	} else {
		ctx.JSON(200, gin.H{
			"message": "获取错误",
		})
	}
}

//Zhuce 注册
func Reg(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"a": 1,
	})
}

//Upload 上传
func Upload(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"a": 1,
	})
}
