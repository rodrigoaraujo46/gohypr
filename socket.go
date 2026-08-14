package gohypr

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	addr net.UnixAddr
	conn *net.UnixConn
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

func parseEvent(line string) (Event, error) {
	line = strings.TrimSuffix(line, "\n")
	_, _, found := strings.Cut(line, ">>")
	if !found {
		return nil, fmt.Errorf("failed to parse event %q", line)
	}

	return nil, nil
}

func (l *client) Listen(handler func(e Event)) error {
	conn, err := net.DialUnix("unix", nil, &l.addr)
	if err != nil {
		return err
	}

	l.conn = conn
	defer func() { _ = l.Close() }()

	for r := bufio.NewReader(conn); ; {
		line, err := r.ReadString('\n')
		if err != nil {
			return err
		}

		event, err := parseEvent(line)
		if err != nil {
			return err
		}

		handler(event)
	}
}

func (l *client) Close() error {
	if l.conn == nil {
		return nil
	}

	err := l.conn.Close()
	l.conn = nil

	return err
}
