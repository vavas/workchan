package endpoints

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/vavas/workchan/app/server/serializers"
	"log"
	"net/http"
)

// JobRequestData uses gin's binding API
// to extract the request parameters.
type JobRequestData struct {
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Company     string `binding:"required"`
	Salary      string `binding:"required"`
}

func CreateJob(c *gin.Context) {
	data := JobRequestData{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid request data"))
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	var jobCreate = "job_create"
	err = KafkaProducer(c).Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &jobCreate, Partition: kafka.PartitionAny},
		Value:          jsonData,
	}, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	response := &serializers.Job{}

	c.JSON(http.StatusCreated, response)
}
