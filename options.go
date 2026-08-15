package gohypr

type option func(*client)

func WithCustomSocketFinder(f func() (string, error)) func(*client) {
	return func(l *client) {
		l.socketFinder = f
	}
}
