package middleware

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

const kafkaProducerCtxKey = "kafkaProducer"

// InjectKafkaProducer creates a gin middleware that adds the
// kafka producer to all request contexts.
func InjectKafkaProducer(kafkaProducer *kafka.Producer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(kafkaProducerCtxKey, kafkaProducer)
		c.Next()
	}
}

func GetKafkaProducer(c *gin.Context) *kafka.Producer {
	val, exists := c.Get(kafkaProducerCtxKey)
	if !exists {
		return nil
	}
	kafkaProducer, ok := val.(*kafka.Producer)
	if !ok {
		return nil
	}
	return kafkaProducer
}
