package workers

import (
	"crypto/md5"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gocraft/work"
	"github.com/sirupsen/logrus"
	"github.com/vavas/workchan/pkg/db"
	"time"
)

type Context struct {
	Logger        logrus.Logger
	DB            *db.DB
	KafkaProducer *kafka.Producer
	JobID         string
}

func uniqueJobID(job *work.Job) string {
	uniqueStr := fmt.Sprintf("%v %v %v", job.Name, job.Args, time.Now())
	hash := md5.Sum([]byte(uniqueStr))
	return fmt.Sprintf("%x", hash)
}
