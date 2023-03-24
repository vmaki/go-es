package jobs

import (
	"fmt"
)

var SayHelloSpec = "*/5 * * * * ?"

func SayHelloHandle() {
	fmt.Println("Hello World")
}
