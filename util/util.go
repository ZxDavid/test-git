package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("fdsdfsjflsdjfskldjflsfjslHFSJHjjjLJLjld")
	result := make([]byte, n)

	//根据系统时间设置时间戳
	rand.Seed(time.Now().Unix())
	for i := range result {
		//随机从letters中获取一个自字符串
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
