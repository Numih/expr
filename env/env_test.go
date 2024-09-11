package env

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		expr     Expr
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
		{"${:default:extra}", "default:extra"},
	}

	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			actual := Parse(test.expr)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
