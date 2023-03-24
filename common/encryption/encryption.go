package encryption

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(str string, y string) string {
	m := md5.New()

	_, err := io.WriteString(m, str+y)
	if err != nil {
		panic(err)
	}

	arr := m.Sum(nil)

	return fmt.Sprintf("%x", arr)
}
