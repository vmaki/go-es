package jobs

import (
	"fmt"
)

var SayHiSpec = "* * * * * *"

func SayHiHandle() {
	fmt.Println("Hi World")
}
