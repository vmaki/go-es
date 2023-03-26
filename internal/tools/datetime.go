package tools

import (
	"fmt"
	"go-es/global"
	"time"
)

func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

func TimeNowByTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(global.GConfig.Timezone)
	return time.Now().In(chinaTimezone)
}
