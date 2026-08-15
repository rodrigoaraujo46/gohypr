package gohypr

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func parseEvents(typ reflect.Type, payload string) (Event, error) {
	fields := strings.Split(payload, ",")

	if len(fields) != typ.NumField() {
		return nil, fmt.Errorf("expected %d fields, got %d", typ.NumField(), len(fields))
	}

	value := reflect.New(typ).Elem()

	for i, raw := range fields {
		if err := setField(value.Field(i), raw); err != nil {
			return nil, fmt.Errorf("field %d: %w", i, err)
		}
	}

	return value.Interface().(Event), nil
}

func setField(field reflect.Value, raw string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(raw)
		return nil

	case reflect.Bool:
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return err
		}

		field.SetBool(v)
		return nil
	}

	return fmt.Errorf("unsupported field type %s", field.Type())
}
