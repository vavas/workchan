package endpoints

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/vavas/workchan/app/server/middleware"
	"github.com/vavas/workchan/pkg/db"
)

// DB extracts the db from the gin context.
func DB(c *gin.Context) *db.DB {
	return middleware.GetDB(c)
}

// KafkaProducer extracts the kafka producer from the gin context.
func KafkaProducer(c *gin.Context) *kafka.Producer {
	return middleware.GetKafkaProducer(c)
}
