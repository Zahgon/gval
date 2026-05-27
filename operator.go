package gval

import (
	"context"

	"github.com/shopspring/decimal"
)

type stage struct {
	Evaluable
	infixBuilder
	operatorPrecedence
}

type stageStack []stage //operatorPrecedence in stacktStage is continuously, monotone ascending

func (s *stageStack) push(b stage) error { _ = "STUB: not implemented"; return nil }

func (s *stageStack) peek() stage { _ = "STUB: not implemented"; return *new(stage) }

func (s *stageStack) pop() stage { _ = "STUB: not implemented"; return *new(stage) }

type infixBuilder func(a, b Evaluable) (Evaluable, error)

func (l Language) isSymbolOperation(r rune) bool { _ = "STUB: not implemented"; return false }

func (l Language) isOperatorPrefix(op string) bool { _ = "STUB: not implemented"; return false }

func (op *infix) initiate(name string) { _ = "STUB: not implemented"; return }

type opFunc func(a, b interface{}) (interface{}, error)

func getStringOpFunc(s func(a, b string) (interface{}, error), f opFunc, typeConversion bool) opFunc {
	_ = "STUB: not implemented"
	return *new(opFunc)
}

func convertToBool(o interface{}) (bool, bool) { _ = "STUB: not implemented"; return false, false }

func getBoolOpFunc(o func(a, b bool) (interface{}, error), f opFunc, typeConversion bool) opFunc {
	_ = "STUB: not implemented"
	return *new(opFunc)
}

func convertToFloat(o interface{}) (float64, bool) { _ = "STUB: not implemented"; return 0, false }

func getFloatOpFunc(o func(a, b float64) (interface{}, error), f opFunc, typeConversion bool) opFunc {
	_ = "STUB: not implemented"
	return *new(opFunc)
}

func convertToDecimal(o interface{}) (decimal.Decimal, bool) {
	_ = "STUB: not implemented"
	return *new(decimal.Decimal), false
}

func getDecimalOpFunc(o func(a, b decimal.Decimal) (interface{}, error), f opFunc, typeConversion bool) opFunc {
	_ = "STUB: not implemented"
	return *new(opFunc)
}

type operator interface {
	merge(operator) operator
	precedence() operatorPrecedence
	initiate(name string)
}

type operatorPrecedence uint8

func (pre operatorPrecedence) merge(op operator) operator {
	_ = "STUB: not implemented"
	return *new(operator)
}

func (pre operatorPrecedence) precedence() operatorPrecedence {
	_ = "STUB: not implemented"
	return *new(operatorPrecedence)
}

func (pre operatorPrecedence) initiate(name string) { _ = "STUB: not implemented"; return }

type infix struct {
	operatorPrecedence
	number       func(a, b float64) (interface{}, error)
	decimal      func(a, b decimal.Decimal) (interface{}, error)
	boolean      func(a, b bool) (interface{}, error)
	text         func(a, b string) (interface{}, error)
	arbitrary    func(a, b interface{}) (interface{}, error)
	shortCircuit func(a interface{}) (interface{}, bool)
	builder      infixBuilder
}

func (op infix) merge(op2 operator) operator { _ = "STUB: not implemented"; return *new(operator) }

type directInfix struct {
	operatorPrecedence
	infixBuilder
}

func (op directInfix) merge(op2 operator) operator {
	_ = "STUB: not implemented"
	return *new(operator)
}

type extension func(context.Context, *Parser) (Evaluable, error)

type postfix struct {
	operatorPrecedence
	f func(context.Context, *Parser, Evaluable, operatorPrecedence) (Evaluable, error)
}

func (op postfix) merge(op2 operator) operator { _ = "STUB: not implemented"; return *new(operator) }
