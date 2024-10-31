package expr

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		e               Expr
		expectedVar     string
		expectedDefault string
	}{
		{"${VAR:default}", "VAR", "default"},
		{"foo", "", "foo"},
		{"${VAR}", "VAR", ""},
		{"${VAR:}", "VAR", ""},
		{"${VAR::}", "VAR", ":"},
		{"${VAR::default}", "VAR", ":default"},
		{"${VAR:default:extra}", "VAR", "default:extra"},
		{"${:default}", "", "default"},
		{"${:}", "", ""},
		{"${:default:extra}", "", "default:extra"},
	}

	for _, test := range tests {
		t.Run(test.e, func(t *testing.T) {
			actualVar, actualDefault := Parse(test.e)
			if actualVar != test.expectedVar {
				t.Errorf("expected var %q, got %q", test.expectedVar, actualVar)
			}
			if actualDefault != test.expectedDefault {
				t.Errorf("expected default %q, got %q", test.expectedDefault, actualDefault)
			}
		})
	}
}
