package eval

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"slices"
	"strconv"
)

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64

	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error

	// String - fmt.Stringer impementation
	String() string
}

// A Var identifies a variable, e.g., x.
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) String() string {
	return string(v)
}

// A literal is a numeric constant, e.g., 3.141.
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}
func (v literal) String() string {
	return strconv.FormatFloat(float64(v), 'f', 2, 64)
}

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}
func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x)
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}
func (u binary) String() string {
	return fmt.Sprintf("(%s %c %s)", u.x, u.op, u.y)
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env),
			c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (u call) String() string {
	str := &bytes.Buffer{}
	res := bufio.NewWriter(str)
	res.WriteString(u.fn)
	res.WriteRune('(')

	for i, arg := range u.args {
		if i > 0 {
			res.WriteString(", ")
		}
		res.WriteString(arg.String())
	}
	res.WriteRune(')')
	res.Flush()
	return str.String()
}

// To evaluate an expression containing variables, we’ll need an environment that maps variable names to values
type Env map[Var]float64

// A min is a numeric constant, e.g., 3.141.
type min struct {
	// fn   string = min
	args []Expr
}

func (l min) Eval(env Env) float64 {
	tmp := make([]float64, len(l.args))
	for i, arg := range l.args {
		tmp[i] = arg.Eval(env)
	}
	return slices.Min(tmp)
}

func (v min) String() string {
	str := new(bytes.Buffer)
	res := bufio.NewWriter(str)
	res.WriteString("min")
	res.WriteRune('(')
	for i, a := range v.args {
		if i > 0 {
			res.WriteString(", ")
		}
		res.WriteString(a.String())
	}
	res.WriteRune(')')
	res.Flush()
	return str.String()
}
