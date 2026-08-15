package main

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/rodrigoaraujo46/gohypr"
)

func main() {
	l, err := gohypr.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	for res := range l.Events() {
		if res.Err != nil {
			slog.Warn("event failed to parse", "error", res.Err)
			continue
		}

		fmt.Printf("%+v\n", res.Event)
	}

	time.Sleep(time.Hour)
}
