package consumer

import (
	"../../config"
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
	"log"
)

func TestConn() {
	brokers := "10.204.192.111:9192"
	splitBrokers := strings.Split(brokers, ",")
	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Version = sarama.V0_10_0_0
	conf.Consumer.Return.Errors = true
	conf.ClientID = config.Data.RestCreds.ExtSystem
	conf.Metadata.Full = false
	conf.Net.SASL.Enable = true
	conf.Net.SASL.User =  config.Data.RestCreds.Login
	conf.Net.SASL.Password = config.Data.SaslPWD
	conf.Net.SASL.Handshake = true
	//conf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
	conf.Net.SASL.Mechanism = sarama.SASLMechanism(sarama.SASLTypePlaintext)

	certs := x509.NewCertPool()
	pemPath := "./cert.crt"
	pemData, err := ioutil.ReadFile(pemPath)
	if err != nil {
		fmt.Println("Couldn't load cert: ", err.Error())
		// handle the error
	}
	certs.AppendCertsFromPEM(pemData)

	conf.Net.TLS.Enable = true
	conf.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: true,
		RootCAs: certs,
	}

	fmt.Println("Connecting")

	fmt.Println("Config errors:", conf.Validate())

	master, err := sarama.NewConsumer(splitBrokers, conf)
	if err != nil {
		fmt.Println("Coulnd't create consumer: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("New consumer - ok")

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	topic := "eaistf2.eaist.uas2.kafka-output.DIT_test_kafka.lot.GetLotList.v-1"

	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {

		fmt.Println("Created consumer")

		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
}