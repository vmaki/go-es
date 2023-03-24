package common

import "fmt"

func MaskPhone(phone string) string {
	if len(phone) < 10 {
		return phone
	}

	return fmt.Sprintf("%s****%s", phone[:3], phone[len(phone)-4:])
}
