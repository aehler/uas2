package main

import (
	"fmt"
	"./lib/rest"
	//"./lib/kafka/consumer"
)



func main() {

	fmt.Println("Starting EAIST uas2.0 test client")

	r, _ := rest.NewRest()

	r.GetLotList(1)

	//consumer.TestConn()

	//kc.ListTopics()

}