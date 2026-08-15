package gohypr

import (
	"bufio"
	"net"
)

type client struct {
	addr net.UnixAddr
}

func NewClient(options ...option) (*client, error) {
	l := &client{addr: net.UnixAddr{Net: "unix"}}

	for _, option := range options {
		option(l)
	}

	if l.addr.Name == "" {
		name, err := getDefaultSocketPath()
		if err != nil {
			return nil, err
		}

		l.addr.Name = name
	}

	return l, nil
}

type EventResult struct {
	Event Event
	Err   error
}

func (l *client) Events() <-chan EventResult {
	events := make(chan EventResult, 32)

	go func() {
		defer close(events)

		conn, err := net.DialUnix("unix", nil, &l.addr)
		if err != nil {
			events <- EventResult{Event: nil, Err: err}
			return
		}

		defer func() { _ = conn.Close() }()

		for r := bufio.NewReader(conn); ; {
			line, err := r.ReadString('\n')
			if err != nil {
				events <- EventResult{Err: err}
				return
			}

			e, err := parseEvent(line)
			events <- EventResult{Event: e, Err: err}
		}
	}()

	return events
}
