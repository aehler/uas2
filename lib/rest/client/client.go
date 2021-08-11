package client

import (
	"../../config"
	httpclient "github.com/bozd4g/go-http-client"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
)

var (
	clnt string
	httpc httpclient.Client
	token string
)

func NewClient() {
	clnt = config.Data.RestAddress
	httpc = httpclient.New(clnt)
	token = "64253be6-370a-44f4-bc95-707964de9f54"
}

func Listen() {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {

		fmt.Println("Starting listener on 8089")

		http.HandleFunc("/uas-test-listener", lh)

		http.ListenAndServe(":8089", nil)

	}()

	for {
		s := <-sigChan
		switch s {
		case os.Signal(syscall.SIGHUP):
			log.Printf("graceful shutdown from signal: %v\n", s)
		default:
			log.Fatalf("exiting from signal: %v\n", s)
		}
	}


}

func lh(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r)

}

func ConnectAndAuth() {

	req, err := httpc.PostWith("/auth", config.Data.RestCreds)
	//req, err := httpc.Post("/auth")

	if err != nil {
		log.Fatal("REST Auth request error", err)
	}

	req.Header.Set("Content-type", "application/json; charset=utf-8")

	resp, err := httpc.Do(req)

	if err != nil {
		log.Fatal("REST Auth error", err)
	}

	r := resp.Get()

	jr := struct{
		UniqueToken string `json:"uniqueToken"`
	}{}

	if err := json.Unmarshal(r.Body, &jr); err == nil {
		token = jr.UniqueToken
	} else {
		log.Fatalln("Error parsing response", err)
	}

}

func GetLotList(id uint) {

	fmt.Println("Getting GetLotById ", id)

	Listen()

	req, err := httpc.Get(fmt.Sprintf("/GetLotList"))

	if err != nil {
		log.Fatal("REST GetLotById request error ", err)
	}

	req.Header.Set("accept", "*/*")

	req.AddCookie(&http.Cookie{Name: "externalSystem", Value: "DIT_test_http"})
	req.AddCookie(&http.Cookie{Name: "token", Value: token})

	fmt.Println(req)

	resp, err := httpc.Do(req)

	if err != nil {
		log.Fatal("REST GetLotById error", err)
	}

	r := resp.Get()

	fmt.Println(r.Status, string(r.Body))

	jr := struct{
		PackageGuid string `json:"packageGuid"`
	}{}

	if err := json.Unmarshal(r.Body, &jr); err == nil {

	} else {
		log.Fatalln("Error parsing response", err)
	}

}