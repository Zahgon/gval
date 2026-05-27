package gval

import (
	"context"
)

// ParseExpression scans an expression into an Evaluable.
func (p *Parser) ParseExpression(c context.Context) (eval Evaluable, err error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

// ParseNextExpression scans the expression ignoring following operators
func (p *Parser) ParseNextExpression(c context.Context) (eval Evaluable, err error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

// ParseSublanguage sets the next language for this parser to parse and calls
// its initialization function, usually ParseExpression.
func (p *Parser) ParseSublanguage(c context.Context, l Language) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

func (p *Parser) parse(c context.Context) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

func parseString(c context.Context, p *Parser) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

func parseNumber(c context.Context, p *Parser) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

func parseDecimal(c context.Context, p *Parser) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

func parseParentheses(c context.Context, p *Parser) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

func (p *Parser) parseOperator(c context.Context, stack *stageStack, eval Evaluable) (st stage, err error) {
	_ = "STUB: not implemented"
	return *new(stage), nil
}

func parseIdent(c context.Context, p *Parser) (call string, alternative func() (Evaluable, error), err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (p *Parser) parseArguments(c context.Context) (args []Evaluable, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inArray(a, b interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func parseIf(c context.Context, p *Parser, e Evaluable) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

func parseJSONArray(c context.Context, p *Parser) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

func parseJSONObject(c context.Context, p *Parser) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}
