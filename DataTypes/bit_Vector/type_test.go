package bit_Vector

import (
	"log"
	"testing"
)

// TestIntSet - checks that its behavior after each operation is equivalent to a set based on built-in maps.
func TestIntSet(t *testing.T) {
	set := IntSet{}
	mapSet := make(map[int]bool, 0)

	nums := []int{1, 3, 5, 6, 142}
	for _, x := range nums {
		set.Add(x)
		mapSet[x] = true

		if set.Len() != len(mapSet) {
			log.Fatal("Length of IntSet and MapSet are not equal")
		}

		if set.Has(x) && mapSet[x] == false {
			log.Fatal("IntSet and MapSet states are not equal")
		}
	}

}
