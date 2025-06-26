package whyreflection

import "strconv"

// Go provides a mechanism to update variables and inspect their values at run time, to
// call  their  methods,  and  to  apply  the  operations  intrinsic  to  their  representation,  all
// without  knowing  their  types  at  compile  time.
// This  mechanism  is  called  reflection.
// Reflection also lets us treat types themselves as first-class values

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// ...similar cases for int16, uint32, and so on...
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}
