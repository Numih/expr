package env

import (
	"os"

	"github.com/numih/expr"
)

// Get parses an expression and substitutes environment variables. If the
// expression does not contain environment variables, the expression is returned
// unchanged.
//
// Examples:
//
//	  Get("${VAR:default}") -> value of VAR if defined, "default" otherwise
//	  Get("foo") -> "foo"
//		Get("${VAR}") -> value of VAR if defined, "" otherwise
func Get(e expr.Expr) string {
	variable, defaultValue := expr.Parse(e)
	if variable == "" {
		return defaultValue
	}

	value, ok := os.LookupEnv(variable)
	if ok {
		return value
	}
	return defaultValue
}
