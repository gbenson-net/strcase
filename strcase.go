// Package strcase converts strings to various cases.
package strcase

import (
	"regexp"

	"github.com/iancoleman/strcase"
)

var (
	nonwordRx = regexp.MustCompile(`[[:^alnum:]]+`)
	freqRx    = regexp.MustCompile(`\b([a-z]) hz\b`)
)

// pp preprocesses a string we'll convert using strcase.
func pp(s string) string {
	s = nonwordRx.ReplaceAllLiteralString(s, " ")
	s = strcase.ToDelimited(s, ' ')
	s = freqRx.ReplaceAllString(s, "${1}hz")
	return s
}

// ToCamel converts a string to CamelCase.
func ToCamel(s string) string {
	return strcase.ToCamel(pp(s))
}

// ToKebab converts a string to kebab-case.
func ToKebab(s string) string {
	return strcase.ToKebab(pp(s))
}

// ToLowerCamel converts a string to lowerCamelCase.
func ToLowerCamel(s string) string {
	return strcase.ToLowerCamel(pp(s))
}

// ToScreamingKebab converts a string to SCREAMING-KEBAB-CASE.
func ToScreamingKebab(s string) string {
	return strcase.ToScreamingKebab(pp(s))
}

// ToScreamingSnake converts a string to SCREAMING_SNAKE_CASE.
func ToScreamingSnake(s string) string {
	return strcase.ToScreamingSnake(pp(s))
}

// ToSnake converts a string to snake_case.
func ToSnake(s string) string {
	return strcase.ToSnake(pp(s))
}
