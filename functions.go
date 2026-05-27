package gval

import (
	"context"
	"reflect"
)

type function func(ctx context.Context, arguments ...interface{}) (interface{}, error)

func toFunc(f interface{}) function { _ = "STUB: not implemented"; return *new(function) }

func createCallArguments(ctx context.Context, t reflect.Type, args []interface{}) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if first argument is a context, use the given execution context
