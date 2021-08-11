package consumer

import (
	"../../config"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"fmt"
)

type Consumer struct{
	pool []string
	conn *kafka.Conn
}

func NewConsumer() *Consumer {

	fmt.Println("Kafka consumer...")

	c, err := kafka.Dial("tcp", config.Data.KafkaAddressPool[0])
	if err != nil {
		log.Fatalln("Failed to connect to kafka", err)
	}

	return &Consumer{
		pool : config.Data.KafkaAddressPool,
		conn : c,
	}
}

func (c *Consumer) ListTopics() {

	partitions, err := c.conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}

	c.conn.Close()
}