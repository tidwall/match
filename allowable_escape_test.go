package match

import "testing"

func TestAllowableEscapedLiterals(t *testing.T) {
	for _, tt := range []struct{ pattern, min, max string }{
		{`a\?b`, "a?b", "a?b"},
		{`\*abc`, "*abc", "*abc"},
		{`a\\b`, `a\b`, `a\b`},
		{`a\?*`, "a?", "a@"},
		{`a\界*`, "a界", "a畍"},
	} {
		t.Run(tt.pattern, func(t *testing.T) {
			min, max := Allowable(tt.pattern)
			if min != tt.min || max != tt.max {
				t.Fatalf("Allowable(%q) = (%q,%q), want (%q,%q)", tt.pattern, min, max, tt.min, tt.max)
			}
			if !Match(tt.min, tt.pattern) {
				t.Fatalf("minimum %q does not match pattern", tt.min)
			}
		})
	}
}
