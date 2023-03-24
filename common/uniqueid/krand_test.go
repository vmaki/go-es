package uniqueid

import (
	"fmt"
	"testing"
)

func TestRandomStringByNum(t *testing.T) {
	str := RandomString(8, KcRandKindNum)
	fmt.Println("随机字符串：" + str)
}

func TestRandomStringByLower(t *testing.T) {
	str := RandomString(8, KcRandKindLower)
	fmt.Println("随机字符串：" + str)
}

func TestRandomStringByUpper(t *testing.T) {
	str := RandomString(8, KcRandKindUpper)
	fmt.Println("随机字符串：" + str)
}

func TestRandomStringByAll(t *testing.T) {
	str := RandomString(8, KcRandKindAll)
	fmt.Println("随机字符串：" + str)
}
