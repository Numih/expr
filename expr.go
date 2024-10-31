package expr

import (
	"strings"
)

// Expr is an expression that can be used to parse variables. The
// expression must be in the format `${VAR:default}` where `VAR` is the name of
// the variable and `default` is the default value to be used if the
// variable is not defined.
//
// The expressions are meant to be used by Getters like `env.Get` to retrieve
// the value of the variable.
//
// Use the `Parse` function to parse variables in an expression.
type Expr = string

// Parse parses an expression, returning the variable and default value if
// present. If a variable is not present, the expression is returned as the
// default value.
func Parse(expr Expr) (variable string, defaultValue string) {

	length := len(expr)

	// Check for valid placeholder format early
	if length < 3 || !strings.HasPrefix(expr, "${") || !strings.HasSuffix(expr, "}") {
		return "", expr
	}

	// Strip leading `${` and trailing `}`
	expr = expr[2 : length-1]

	// Split the variable and default value if present
	variable, defaultValue, _ = strings.Cut(expr, ":")

	return variable, defaultValue
}
