package env

import (
	"fmt"
	"os"

	"github.com/numih/expr"
)

const (
	EnvVariableName = "ENV"
	Development     = "development"
	Production      = "production"
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

// IsDevelopment returns true if the environment is set to "development".
func IsDevelopment() bool {
	return Get(fmt.Sprintf("${%s}", EnvVariableName)) == Development
}

// IsProduction returns true if the environment is set to "production".
func IsProduction() bool {
	return Get(fmt.Sprintf("${%s}", EnvVariableName)) == Production
}
