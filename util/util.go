package util

import (
	"crypto/rand"
	"math/big"
)

// RandomString 生成一个指定长度的随机字符
func RandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 定义一个包含所有可能字符的字符
	result := make([]byte, length)                                                 // 创建一个长度为 length 的字节切
	for i := range result {
		val, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars)))) // 生成一个随机数
		result[i] = chars[val.Int64()]                                 // 将随机数映射到相应的字符
	}
	return string(result) // 将字节切片转换为字符串并返回
}
