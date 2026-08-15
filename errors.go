package gohypr

import "errors"

var ErrUnexpectedFields = errors.New(
	"unexpected number of fields: Hyprland uses commas as event delimiters, " +
		"so payloads containing commas cannot be reliably parsed",
)
