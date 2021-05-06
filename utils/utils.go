package utils

import (
	"runtime/debug"

	"github.com/nju-iot/cron_jobs/logs"
)

// RecoverPanic ...
func RecoverPanic() {
	if x := recover(); x != nil {
		logs.Error("runtime panic: %v\n%v", x, string(debug.Stack()))
	}
}
