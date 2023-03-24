package cronx

import (
	"github.com/robfig/cron"
	"time"
)

type Cron struct {
	C *cron.Cron
}

func NewCron() *Cron {
	return &Cron{
		C: cron.New(),
	}
}

func (c *Cron) Register() {
	RegisterApiJobs(c.C)
}

func (c *Cron) Start() {
	c.C.Start()

	t1 := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
