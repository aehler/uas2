package rest

import ("fmt"
	"./client"
	"os"
	"os/signal"
	"syscall"
	"log"
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

func (r *Rest) GetLotList(id uint) {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	client.GetLotList(id)


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