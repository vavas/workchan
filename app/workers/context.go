package workers

import (
	"crypto/md5"
	"fmt"
	"github.com/gocraft/work"
	"time"
)

func uniqueJobID(job *work.Job) string {
	uniqueStr := fmt.Sprintf("%v %v %v", job.Name, job.Args, time.Now())
	hash := md5.Sum([]byte(uniqueStr))
	return fmt.Sprintf("%x", hash)
}
