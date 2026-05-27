package gval

import (
	"fmt"
	"text/scanner"
)

// Parser parses expressions in a Language into an Evaluable
type Parser struct {
	scanner scanner.Scanner
	Language
	lastScan   rune
	camouflage error
}

func newParser(expression string, l Language) *Parser { _ = "STUB: not implemented"; return nil }

func (p *Parser) resetScannerProperties() { _ = "STUB: not implemented"; return }

// SetWhitespace sets the behavior of the whitespace matcher. The given
// characters must be less than or equal to 0x20 (' ').
func (p *Parser) SetWhitespace(chars ...rune) { _ = "STUB: not implemented"; return }

// SetMode sets the tokens that the underlying scanner will match.
func (p *Parser) SetMode(mode uint) { _ = "STUB: not implemented"; return }

// SetIsIdentRuneFunc sets the function that matches ident characters in the
// underlying scanner.
func (p *Parser) SetIsIdentRuneFunc(fn func(ch rune, i int) bool) {
	_ = "STUB: not implemented"
	return
}

// Scan reads the next token or Unicode character from source and returns it.
// It only recognizes tokens t for which the respective Mode bit (1<<-t) is set.
// It returns scanner.EOF at the end of the source.
func (p *Parser) Scan() rune { _ = "STUB: not implemented"; return 0 }

func (p *Parser) isCamouflaged() bool { _ = "STUB: not implemented"; return false }

// Camouflage rewind the last Scan(). The Parser holds the camouflage error until
// the next Scan()
// Do not call Rewind() on a camouflaged Parser
func (p *Parser) Camouflage(unit string, expected ...rune) { _ = "STUB: not implemented"; return }

// Peek returns the next Unicode character in the source without advancing
// the scanner. It returns EOF if the scanner's position is at the last
// character of the source.
// Do not call Peek() on a camouflaged Parser
func (p *Parser) Peek() rune { _ = "STUB: not implemented"; return 0 }

var errCamouflageAfterNext = fmt.Errorf("Camouflage() after Next()")

// Next reads and returns the next Unicode character.
// It returns EOF at the end of the source.
// Do not call Next() on a camouflaged Parser
func (p *Parser) Next() rune { _ = "STUB: not implemented"; return 0 }

// TokenText returns the string corresponding to the most recently scanned token.
// Valid after calling Scan().
func (p *Parser) TokenText() string { _ = "STUB: not implemented"; return "" }

// Expected returns an error signaling an unexpected Scan() result
func (p *Parser) Expected(unit string, expected ...rune) error {
	_ = "STUB: not implemented"
	return nil
}

type unexpectedRune struct {
	unit     string
	expected []rune
	got      rune
}

func (err unexpectedRune) Error() string { _ = "STUB: not implemented"; return "" }
