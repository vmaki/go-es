package uniqueid

import (
	"fmt"
	"testing"
)

func TestGenSn(t *testing.T) {
	str := GenSn(SnPrefixAsynq)
	fmt.Println("随机id：" + str)
}
