package env

import (
	"os"
	"strings"
)

// Expr is an expression that can be used to parse environment variables. The
// expression must be in the format `${VAR:default}` where `VAR` is the name of
// the environment variable and `default` is the default value to be used if the
// environment variable is not defined.
//
// Use the `Parse` function to parse environment variables in an expression.
type Expr = string

// Parse parses an expression and substitutes environment variables. If the
// expression does not contain environment variables, the expression is returned
// unchanged.
//
// Examples:
//
//	  Parse("${VAR:default}") -> value of VAR if defined, "default" otherwise
//	  Parse("foo") -> "foo"
//		Parse("${VAR}") -> value of VAR if defined, "" otherwise
func Parse(expr Expr) string {

	length := len(expr)

	// Check for valid placeholder format early
	if length < 3 || !strings.HasPrefix(expr, "${") || !strings.HasSuffix(expr, "}") {
		return expr
	}

	// Strip leading `${` and trailing `}`
	expr = expr[2 : length-1]

	// Split the key and default value if present
	key, defaultValue, _ := strings.Cut(expr, ":")

	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
