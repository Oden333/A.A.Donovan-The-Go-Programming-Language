package main

import (
	"fmt"
	"math"
	"testing"

	gopl_eval "gopl.io/ch7/eval"
)

func TestCoverage(t *testing.T) {
	var tests = []struct {
		input string
		env   gopl_eval.Env
		want  string // expected error from Parse/Check or result from Eval
	}{
		{"x % 2", nil, "unexpected '%'"},
		{"!true", nil, "unexpected '!'"},
		{"log(10)", nil, `unknown function "log"`},
		{"sqrt(1, 2)", nil, "call to sqrt has 2 args, want 1"},
		{"sqrt(A / pi)", gopl_eval.Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", gopl_eval.Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", gopl_eval.Env{"F": -40}, "-40"},
	}
	for _, test := range tests {
		expr, err := gopl_eval.Parse(test.input)
		if err == nil {
			err = expr.Check(map[gopl_eval.Var]bool{})
		}

		if err != nil {
			if err.Error() != test.want {
				t.Errorf("%s: got %q, want %q", test.input, err, test.want)
			}
			continue
		}
		got := fmt.Sprintf("%.6g",
			expr.Eval(test.env))
		if got != test.want {
			t.Errorf("%s: %v => %s, want %s", test.input, test.env, got, test.want)
		}
	}
}
