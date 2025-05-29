//* Expression language consists of
//* floating-point literals;
//* the binary operators +, -, *, and /;
//* the unary operators -x and +x;
//* function calls pow(x,y), sin(x), and sqrt(x);
//* variables  such  as  x  and  pi;  and  of  course
//* parentheses  and  standard operator  precedence.

// * All  values  are  of  type  float64
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return expr.Eval(Env{"x": x, "y": y, "r": r})
	})
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", plot)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
