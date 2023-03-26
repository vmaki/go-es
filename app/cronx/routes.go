package cronx

import (
	"github.com/robfig/cron"
	"go-es/app/cronx/jobs"
)

func RegisterApiJobs(c *cron.Cron) {
	c.AddFunc(jobs.SayHiSpec, jobs.SayHiHandle)
	c.AddFunc(jobs.SayHelloSpec, jobs.SayHelloHandle)
}
