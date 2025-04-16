package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(25);not null;comment:'用户�?"`
	Telphone string `json:"telphone" gorm:"type:varchar(11);not null;comment:'手机�?"`
	Password string `json:"password" gorm:"type:varchar(255);not null;comment:'密码'"`
}

func main() {
	// 数据库初始化
	db := InitDB()
	sqldb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqldb.SetMaxIdleConns(10)   // 设置空闲连接池的最大连接数
	sqldb.SetMaxOpenConns(100)  // 设置打开数据库连接的最大数�?
	sqldb.SetConnMaxLifetime(0) // 设置了连接的最大生命周期，0表示不限�?

	// 延迟关闭数据库连�?
	router := gin.Default()
	router.POST("/api/auth/register", func(ctx *gin.Context) {

		// 获取请求参数
		name := ctx.PostForm("name")
		telphone := ctx.PostForm("telphone")
		password := ctx.PostForm("password")
		// 数据验证
		if len(telphone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号长度必须为11�?",
			})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码长度不能小于6�?",
			})
			return
		}
		if len(name) == 0 {
			// 生成随机十位数的用户�?
			name = RandomString(10)
		}
		// 判断手机号是否存�?
		if isTelphoneExist(db, telphone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号已存在",
			})
			return
		}

		log.Printf("用户�?%v,密码%v,手机�?:%v", name, password, telphone)
		// 返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	router.Run() // 监听并在 0.0.0.0:8080 上启动服�?
}

// RandomString 生成一个指定长度的随机字符�?
func RandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 定义一个包含所有可能字符的字符�?
	result := make([]byte, length)                                                 // 创建一个长度为 length 的字节切�?
	for i := range result {
		val, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars)))) // 生成一个随机数，范围为 0 �? len(chars)-1
		result[i] = chars[val.Int64()]                                 // 将随机数映射到相应的字符�?
	}
	return string(result) // 将字节切片转换为字符串并返回
}
func isTelphoneExist(db *gorm.DB, telphone string) bool {
	var user User
	// 查询手机号是否存�?
	result := db.Where("telphone = ?", telphone).First(&user)
	if result.RowsAffected > 0 {
		return true // 手机号已存在
	}
	return false // 手机号不存在
}

func InitDB() *gorm.DB {
	// driveName := "mysql"
	user := "root"
	pass := "123456"
	host := "localhost"
	port := "3306"
	database := "ceshi"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	return db
}

// curl 117.72.11.171:8080/ping
