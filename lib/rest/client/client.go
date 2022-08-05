package client

import (
	"../../config"
	httpclient "github.com/bozd4g/go-http-client"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

var (
	clnt string
	httpc httpclient.Client
	token string
)

func NewClient() {
	clnt = config.Data.RestAddress
	httpc = httpclient.New(clnt)
	token = "64253be6-370a-44f4-bc95-707964de9f54" //This is a temp token from Auth method
}

func Listen() {

	func() {

		fmt.Println("Starting listener on 8089")

		http.HandleFunc("/uas-test-listener", lh)

		http.ListenAndServe(":8089", nil)

	}()

}

func lh(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r)

	htmlData, err := ioutil.ReadAll(r.Body) //<--- here!

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(htmlData))


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

	go Listen()

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

		log.Println(jr, string(r.Body))

	} else {
		log.Fatalln("Error parsing response", err)
	}



}