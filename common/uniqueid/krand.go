package uniqueid

import (
	"math/rand"
	"time"
)

const (
	KcRandKindNum   = 0 // 纯数字
	KcRandKindLower = 1 // 小写字母
	KcRandKindUpper = 2 // 大写字母
	KcRandKindAll   = 3 // 数字、大小写字母
)

// RandomString 随机字符串
func RandomString(size int, kind int) string {
	k, ks, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		if isAll {
			k = rand.Intn(3)
		}

		scope, base := ks[k][0], ks[k][1]
		result[i] = uint8(base + rand.Intn(scope))
	}

	return string(result)
}
