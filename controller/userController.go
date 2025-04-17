package controller

import (
	"log"
	"net/http"

	"github.com/260444/ginEssential/common"
	"github.com/260444/ginEssential/model"
	"github.com/260444/ginEssential/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	// 创建用户

	hashdPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "加密错误",
		})
		return
	}

	newuser := model.User{
		Name:     name,
		Telphone: telphone,
		Password: string(hashdPassword),
	}
	// 插入数据
	DB.Create(&newuser)
	// 返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {
	// 获取数据库连接
	DB := common.GetDB()
	// 获取参数
	telphone := ctx.PostForm("telphone")
	password := ctx.PostForm("password")
	//数据验证
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
	// 判断手机号是否存在
	var user model.User
	DB.Where("telphone = ?", telphone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return

	}
	// 判断密码是否正确
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统错误",
		})
		log.Printf("token generate error: %v", err)
		return
	}
	// 返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{"token": token},
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": user},
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
