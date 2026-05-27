package gval

import (
	"context"
	"reflect"
)

// Selector allows for custom variable selection from structs
//
// Return value is again handled with variable() until end of the given path
type Selector interface {
	SelectGVal(c context.Context, key string) (interface{}, error)
}

// Evaluable evaluates given parameter
type Evaluable func(c context.Context, parameter interface{}) (interface{}, error)

// EvalInt evaluates given parameter to an int
func (e Evaluable) EvalInt(c context.Context, parameter interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EvalFloat64 evaluates given parameter to a float64
func (e Evaluable) EvalFloat64(c context.Context, parameter interface{}) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EvalBool evaluates given parameter to a bool
func (e Evaluable) EvalBool(c context.Context, parameter interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// EvalString evaluates given parameter to a string
func (e Evaluable) EvalString(c context.Context, parameter interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Const Evaluable represents given constant
func (*Parser) Const(value interface{}) Evaluable {
	_ = "STUB: not implemented"
	return *

	//go:noinline
	new(Evaluable)
}

func constant(value interface{}) Evaluable { _ = "STUB: not implemented"; return *new(Evaluable) }

// Var Evaluable represents value at given path.
// It supports with default language VariableSelector:
//
//		map[interface{}]interface{},
//		map[string]interface{} and
//		[]interface{} and via reflect
//		struct fields,
//		struct methods,
//		slices and
//	 map with int or string key.
func (p *Parser) Var(path ...Evaluable) Evaluable {
	_ = "STUB: not implemented"
	return *new(Evaluable)
}

// Evaluables is a slice of Evaluable.
type Evaluables []Evaluable

// EvalStrings evaluates given parameter to a string slice
func (evs Evaluables) EvalStrings(c context.Context, parameter interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func variable(path Evaluables) Evaluable { _ = "STUB: not implemented"; return *new(Evaluable) }

func reflectSelect(key string, value interface{}) (selection interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// key didn't exist. Check if there is a bound method

// key not an int. Check if there is a bound method

func resolvePotentialPointer(value reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func reflectConvertTo(k reflect.Kind, value string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (*Parser) callFunc(fun function, args ...Evaluable) Evaluable {
	_ = "STUB: not implemented"
	return *new(Evaluable)
}

func (*Parser) callEvaluable(fullname string, fun Evaluable, args ...Evaluable) Evaluable {
	_ = "STUB: not implemented"
	return *new(Evaluable)
}

// IsConst returns if the Evaluable is a Parser.Const() value
func (e Evaluable) IsConst() bool { _ = "STUB: not implemented"; return false }

func regEx(a, b Evaluable) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

func notRegEx(a, b Evaluable) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}
