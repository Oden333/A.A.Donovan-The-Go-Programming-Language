package main

import "github.com/Oden333/A.A.Donovan-The-Go-Programming-Language/tree/master/7.Interfaces/7.9_Example_Expression_Evaluator/eval"

func main() {
	e, _ := eval.Parse("sqrt(A / pi)")
	eval.Display("e", e)
}
