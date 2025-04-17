package controller

import (
	"log"
	"net/http"

	"github.com/260444/ginEssential/common"
	"github.com/260444/ginEssential/model"
	"github.com/260444/ginEssential/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB() // 获取数据库连接
	// 获取请求参数
	name := ctx.PostForm("name")
	telphone := ctx.PostForm("telphone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telphone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号长度必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码长度不能小于6位",
		})
		return
	}
	if len(name) == 0 {
		// 生成随机十位数的用户名
		name = util.RandomString(10)
	}
	// 判断手机号是否存在
	if isTelphoneExist(DB, telphone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号已存在",
		})
		return
	}
	newuser := model.User{
		Name:     name,
		Telphone: telphone,
		Password: password,
	}
	DB.Create(&newuser) // 插入数据
	log.Printf("用户名%v,密码%v,手机号:%v", name, password, telphone)
	// 返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

// 判断手机号是否存在
func isTelphoneExist(db *gorm.DB, telphone string) bool {
	var user model.User
	// 查询手机号是否存在
	result := db.Where("telphone = ?", telphone).First(&user)
	if result.RowsAffected > 0 {
		return true // 手机号已存在
	}
	return false // 手机号不存在
}
