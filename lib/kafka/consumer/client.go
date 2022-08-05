package consumer
//
//import (
//	"../../config"
//	kafka "src/github.com/segmentio/kafka-go"
//	"log"
//	"fmt"
//	"time"
//	"src/github.com/segmentio/kafka-go/sasl/plain"
//	"context"
//)
//
//type Consumer struct{
//	pool []string
//	conn *kafka.Conn
//}
//
//func NewConsumer() *Consumer {
//
//	fmt.Println("Kafka consumer...")
//
//	dialer := &kafka.Dialer{
//		Timeout:   10 * time.Second,
//		DualStack: true,
//		//TLS:       &tls.Config{
//		//	Certificates:
//		//},
//		SASLMechanism: plain.Mechanism{
//			Username: "DIT_test_kafka",
//			Password: "FRq5NBqpmei8",
//		},
//	}
//
//	ctx := context.TODO()
//
//	c, err := dialer.DialContext(ctx, "tcp", config.Data.KafkaAddressPool[0])
//	if err != nil {
//		log.Fatalln("Failed to connect to kafka", err)
//	}
//
//	return &Consumer{
//		pool : config.Data.KafkaAddressPool,
//		conn : c,
//	}
//}
//
//func (c *Consumer) ListTopics() {
//
//	partitions, err := c.conn.ReadPartitions()
//	if err != nil {
//		panic(err.Error())
//	}
//
//	m := map[string]struct{}{}
//
//	for _, p := range partitions {
//		m[p.Topic] = struct{}{}
//	}
//	for k := range m {
//		fmt.Println(k)
//	}
//
//	c.conn.Close()
//}
