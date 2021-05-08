package workers

import (
	"github.com/gocraft/work"
	"github.com/sirupsen/logrus"
	"time"
)

type jobLog struct {
	logger logrus.Logger
}

func (j *jobLog) JobStarted(job *work.Job) {
	j.logger.Printf("Starting job=%s, id=%s", job.Name, uniqueJobID(job))
}

func (j *jobLog) JobFinished(job *work.Job, elapsed time.Duration) {
	j.logger.Printf("Finished job=%s, id=%s, elapsed=%v", job.Name, uniqueJobID(job), elapsed)
}
