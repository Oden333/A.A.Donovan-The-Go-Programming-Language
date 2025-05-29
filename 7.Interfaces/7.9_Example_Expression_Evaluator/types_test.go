package main

import (
	"fmt"
	"math"
	"testing"
)

// func TestString(t *testing.T) {
// 	tests := []struct {
// 		expr string
// 		env  Env
// 		want string
// 	}{
// 		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
// 		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
// 		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
// 		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
// 		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
// 		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
// 	}
// 	var prevExpr string
// 	for _, test := range tests {
// 		// Print expr only when it changes.
// 		if test.expr != prevExpr {
// 			fmt.Printf("\n%s\n", test.expr)
// 			prevExpr = test.expr
// 		}
// 		expr, err := Parse(test.expr)
// 		if err != nil {
// 			t.Error(err) // parse error
// 			continue
// 		}

//			got := fmt.Sprintf("%.6g", expr.Eval(test.env))
//			fmt.Printf("\t%v => %s\n", test.env, got)
//			if got != test.want {
//				t.Errorf("%s.Eval() in %s = %q, want %q\n", test.expr, test.env, got, test.want)
//			}
//		}
//	}
func TestString(t *testing.T) {
	tcs := []struct {
		expr   string
		env    Env
		result string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},

		{"min(x,y,z)", Env{"x": 123, "y": 133, "z": 11}, "11"},
	}
	for _, tc := range tcs {
		expr, err := Parse(tc.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		s := expr.String()
		reexpr, err := Parse(s)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Logf("Input: \n\t%s \n\t%s \n\t%s", tc.expr, s, reexpr)

		got := fmt.Sprintf("%.6g", expr.Eval(tc.env))
		regot := fmt.Sprintf("%.6g", reexpr.Eval(tc.env))

		if got != tc.result || regot != tc.result {
			t.Errorf("\n%s.Eval() in %v = %q,\n%s.Eval() in %v = %q, result %q\n",
				tc.expr, tc.env, got, reexpr, tc.env, regot, tc.result)
		}
	}
}
