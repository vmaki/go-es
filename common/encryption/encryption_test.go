package encryption

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	pwd := Md5("123456", "go-es")
	fmt.Println(pwd, len(pwd))
}
