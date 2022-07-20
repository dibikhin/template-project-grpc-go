package client

import (
	"log"
	"time"
)

type TODOService interface {
	TODO() error
}

func RunLoop(s TODOService) {
	for {
		err := s.TODO()
		log.Println("client: called TODO()")
		if err != nil {
			log.Printf("client: TODO(): %v", err)
		}
		time.Sleep(time.Second)
	}
}
