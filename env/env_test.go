package env

import (
	"testing"

	"github.com/numih/expr"
)

func TestGet(t *testing.T) {
	tests := []struct {
		e        expr.Expr
		expected string
	}{
		{"${VAR:default}", "default"},
		{"foo", "foo"},
		{"${VAR}", ""},
		{"${VAR:}", ""},
		{"${VAR::}", ":"},
		{"${VAR::default}", ":default"},
		{"${VAR:default:extra}", "default:extra"},
		{"${:default}", "default"},
		{"${:}", ""},
		{"${}", ""},
		{"${:default:extra}", "default:extra"},
	}

	for _, test := range tests {
		t.Run(test.e, func(t *testing.T) {
			actual := Get(test.e)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
