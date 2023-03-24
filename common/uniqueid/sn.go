package uniqueid

import (
	"fmt"
	"time"
)

// SnPrefix 生成唯一sn单号
type SnPrefix string

const (
	SnPrefixAsynq SnPrefix = "ASYNQ" // 延迟任务唯一id
)

// GenSn 生成唯一sn单号
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), RandomString(8, KcRandKindNum))
}
