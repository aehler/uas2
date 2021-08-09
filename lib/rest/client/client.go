package client

import (
	"../../config"
	httpclient "github.com/bozd4g/go-http-client"
	"log"
	"fmt"
)

var (
	clnt string
	httpc httpclient.Client
)

func NewClient() {
	clnt = config.Data.RestAddress
	httpc = httpclient.New(clnt)
}

func Listen() {

}

func ConnectAndAuth() {

	req, err := httpc.PostWith("/auth", config.Data.RestCreds)
	//req, err := httpc.Post("/auth")

	if err != nil {
		log.Fatal("REST Auth request error", err)
	}

	req.Header.Set("Content-type", "application/json; charset=utf-8")

	//buf := new(bytes.Buffer)
	//
	//rc := config.Data.RestCreds
	//
	//errj := json.NewEncoder(buf).Encode(rc)
	//if errj != nil {
	//	panic(errj)
	//}
	//
	//fmt.Println(buf)
	//
	//req.Body = ioutil.NopCloser(buf)

	//p := []byte{}
	//fmt.Println(req.Body.Read(p))
	//fmt.Println(string(p))

//	fmt.Println(req)

	resp, err := httpc.Do(req)

	if err != nil {
		log.Fatal("REST Auth error", err)
	}

	r := resp.Get()

	fmt.Println(r.StatusCode, r.Status, string(r.Body))

}

func send() {

}