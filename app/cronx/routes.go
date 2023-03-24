package cronx

import (
	"github.com/robfig/cron"
	"go-es/app/cronx/jobs"
	"go-es/internal/tools"
)

func RegisterApiJobs(c *cron.Cron) {
	if !tools.IsLocal() {
		c.AddFunc(jobs.SayHiSpec, jobs.SayHiHandle)
		c.AddFunc(jobs.SayHelloSpec, jobs.SayHelloHandle)
	}
}
