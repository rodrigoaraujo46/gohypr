package main

import (
	"time"

	"github.com/rodrigoaraujo46/gohypr"
)

func main() {
	c := gohypr.NewClient()

	c.OnOpenWindow(func(e gohypr.EventOpenWindow, err error) {
	})

	time.Sleep(time.Hour)
}
