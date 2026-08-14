package main

import (
	"fmt"
	"log"

	"github.com/rodrigoaraujo46/gohypr"
)

func main() {
	l, err := gohypr.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	_ = l.Listen(func(e gohypr.Event) {
		if event, ok := gohypr.AsType[gohypr.EventOpenWindow](e); ok {
			fmt.Println(event)
		}
	})
}
