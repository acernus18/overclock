package overclock

import (
	"log"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	pool := NewExecutor(2, 2, func(err error) {
		log.Fatal(err)
	})
	if err := pool.Start(); err != nil {
		log.Fatal(err)
	}
	if err := pool.Execute(func() error {
		time.Sleep(3 * time.Second)
		log.Println("hello")
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := pool.Execute(func() error {
		time.Sleep(2 * time.Second)
		log.Println("world")
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := pool.Close(); err != nil {
		log.Fatal(err)
	}
}
