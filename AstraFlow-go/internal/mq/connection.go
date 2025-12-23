package mq

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func NewConnection() (*amqp.Connection, error) {
	mqURL := os.Getenv("MQ_URL")
	if mqURL == "" {
		return nil, fmt.Errorf("MQ_URL is not set")
	}

	return amqp.Dial(mqURL)
}
