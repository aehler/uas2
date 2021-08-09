package rest

import ("fmt"
	"./client"
)

type Rest struct {
	Alive bool
	Token string
}

func NewRest() (*Rest, error) {

	fmt.Println("Rest client...")

	client.NewClient()
	client.ConnectAndAuth()

	return &Rest{
		Alive: true,
	}, nil
}