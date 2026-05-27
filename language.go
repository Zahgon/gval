package gval

import (
	"context"

	"github.com/shopspring/decimal"
)

// Language is an expression language
type Language struct {
	prefixes        map[interface{}]extension
	operators       map[string]operator
	operatorSymbols map[rune]struct{}
	init            extension
	def             extension
	selector        func(Evaluables) Evaluable
}

// NewLanguage returns the union of given Languages as new Language.
func NewLanguage(bases ...Language) Language { _ = "STUB: not implemented"; return *new(Language) }

func newLanguage() Language { _ = "STUB: not implemented"; return *new(Language) }

// NewEvaluable returns an Evaluable for given expression in the specified language
func (l Language) NewEvaluable(expression string) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

// NewEvaluableWithContext returns an Evaluable for given expression in the specified language using context
func (l Language) NewEvaluableWithContext(c context.Context, expression string) (Evaluable, error) {
	_ = "STUB: not implemented"
	return *new(Evaluable), nil
}

// Evaluate given parameter with given expression
func (l Language) Evaluate(expression string, parameter interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Evaluate given parameter with given expression using context
func (l Language) EvaluateWithContext(c context.Context, expression string, parameter interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Function returns a Language with given function.
// Function has no conversion for input types.
//
// If the function returns an error it must be the last return parameter.
//
// If the function has (without the error) more then one return parameter,
// it returns them as []interface{}.
func Function(name string, function interface{}) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// Constant returns a Language with given constant
func Constant(name string, value interface{}) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// PrefixExtension extends a Language
func PrefixExtension(r rune, ext func(context.Context, *Parser) (Evaluable, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// Init is a language that does no parsing, but invokes the given function when
// parsing starts. It is incumbent upon the function to call ParseExpression to
// continue parsing.
//
// This function can be used to customize the parser settings, such as
// whitespace or ident behavior.
func Init(ext func(context.Context, *Parser) (Evaluable, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// DefaultExtension is a language that runs the given function if no other
// prefix matches.
func DefaultExtension(ext func(context.Context, *Parser) (Evaluable, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// PrefixMetaPrefix chooses a Prefix to be executed
func PrefixMetaPrefix(r rune, ext func(context.Context, *Parser) (call string, alternative func() (Evaluable, error), err error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// PrefixOperator returns a Language with given prefix
func PrefixOperator(name string, e Evaluable) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// PostfixOperator extends a Language.
func PostfixOperator(name string, ext func(context.Context, *Parser, Evaluable) (Evaluable, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// InfixOperator for two arbitrary values.
func InfixOperator(name string, f func(a, b interface{}) (interface{}, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// InfixShortCircuit operator is called after the left operand is evaluated.
func InfixShortCircuit(name string, f func(a interface{}) (interface{}, bool)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// InfixTextOperator for two text values.
func InfixTextOperator(name string, f func(a, b string) (interface{}, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// InfixNumberOperator for two number values.
func InfixNumberOperator(name string, f func(a, b float64) (interface{}, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// InfixDecimalOperator for two decimal values.
func InfixDecimalOperator(name string, f func(a, b decimal.Decimal) (interface{}, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// InfixBoolOperator for two bool values.
func InfixBoolOperator(name string, f func(a, b bool) (interface{}, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// Precedence of operator. The Operator with higher operatorPrecedence is evaluated first.
func Precedence(name string, operatorPrecendence uint8) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

// InfixEvalOperator operates on the raw operands.
// Therefore it cannot be combined with operators for other operand types.
func InfixEvalOperator(name string, f func(a, b Evaluable) (Evaluable, error)) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

func newLanguageOperator(name string, op operator) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}

func (l *Language) makePrefixKey(key string) interface{} { _ = "STUB: not implemented"; return nil }

func (l *Language) makeInfixKey(key string) string { _ = "STUB: not implemented"; return "" }

// VariableSelector returns a Language which uses given variable selector.
// It must be combined with a Language that uses the vatiable selector. E.g. gval.Base().
func VariableSelector(selector func(path Evaluables) Evaluable) Language {
	_ = "STUB: not implemented"
	return *new(Language)
}
