package controller

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"gindemo/model"
	"gindemo/util"
)

// Login 登录
func Login(ctx *gin.Context) {
	var user model.User
	if ctx.ShouldBind(&user) != nil {
		return // 错误
	}
	result := model.DB.First(&user, user)
	// result.RowsAffected // 返回找到的记录数
	// result.Error        // returns error
	fmt.Println(&user)

	if result.Error != nil {
		notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)
		fmt.Println(notFound, result.Error, gorm.ErrRecordNotFound)
		if notFound {
			ctx.JSON(200, gin.H{
				"code": 0,
				"msg":  "未查询到记录",
				// "token":  token,
				// "claims": claims,
			})
			return
		}
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg":  "未知原因",
		})
		return
	}

	token, _ := util.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), user.Name)
	claims, _ := util.ParseToken(token)

	util.CheckExp(token)
	ctx.JSON(200, gin.H{
		"code":   1,
		"msg":    "成功",
		"token":  token,
		"claims": claims,
	})
}

// CheckLogin 检查登录
func CheckLogin(ctx *gin.Context) {
	token, _ := ctx.GetQuery("token")
	Check := util.CheckExp(token)
	ctx.JSON(200, gin.H{"code": 1, "msg": "成功", "data": Check})
}

//Reg 注册
func Reg(ctx *gin.Context) {
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

//Upload 上传
func Upload(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"a": 1,
	})
}
