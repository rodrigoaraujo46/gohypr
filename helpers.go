package gohypr

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func validateSocket(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("not a valid socket path: %w", err)
	}

	if f.Mode()&os.ModeSocket == 0 {
		return errors.New("path is not a Unix socket")
	}

	return nil
}

func DefaultSocketFinder() (string, error) {
	dir, ok := os.LookupEnv("XDG_RUNTIME_DIR")
	if !ok {
		return "", errors.New("'XDG_RUNTIME_DIR' not set")
	}

	instance, ok := os.LookupEnv("HYPRLAND_INSTANCE_SIGNATURE")
	if !ok {
		return "", errors.New("'HYPRLAND_INSTANCE_SIGNATURE' not set")
	}

	path := filepath.Join(dir, "hypr", instance, ".socket2.sock")
	if err := validateSocket(path); err != nil {
		return "", err
	}

	return path, nil
}

func AsType[T Event](event Event) (T, bool) {
	e, ok := event.(T)
	return e, ok
}
