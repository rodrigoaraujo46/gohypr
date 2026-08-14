package gohypr

type option func(*client)

func WithCustomPath(path string) func(*client) {
	return func(l *client) {
		l.addr.Name = path
	}
}
