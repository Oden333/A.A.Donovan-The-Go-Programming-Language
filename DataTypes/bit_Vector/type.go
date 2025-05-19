// Package bit_Vector - representation for an IntSet is a set of small non-negative integers.
// * A set represented by a map[T]bool is very flexible but,
// * for certain problems, a specialized representation may outperform it.
//
// ? For example, in domains such as dataflow analysis
// ? where set elements are small non-negative integers, sets have many elements,
// ? and set operations like union and intersection are common, a bit vector is ideal.
//
// ! A  bit  vector  uses  a  slice  of  unsigned  integer  values  or  “words,”
// ! each  bit  of  which represents a possible element of the set.
// ! The set contains i if the i-th bit is set.
package bit_Vector

import (
	"bytes"
	"fmt"
	"math/rand"
	"slices"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// ! Since each word has UINTSIZE bits, to locate the bit for x,
// ! we use the quotient x/UINTSIZE as the word  index
// ! the  remainder  x%UINTSIZE  as  the  bit  index  within  that  word
//
// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/UINTSIZE, uint(x%UINTSIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/UINTSIZE, uint(x%UINTSIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// ! The UnionWith  operation  uses  the  bitwise  OR  operator  |
// ! to  compute  the  union  UINTSIZE elements at a time
//
// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINTSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", UINTSIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len - return the number of elements
func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		for word > 0 {
			// fmt.Println(fmt.Sprintf("%[1]d : %[1]b", word))
			if a := word & 1; a == 1 {
				len++
			}
			word = word >> 1
		}
	}
	return len
}

// Remove - remove x from the set
func (s *IntSet) Remove(num uint) {
	word, bit := num/UINTSIZE, num%UINTSIZE
	if len(s.words) < int(word) {
		return
	}
	s.words[int(word)] = s.words[int(word)] & ^(1 << bit)
}

// Clear - remove all elements from the set
func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

// Copy - return a copy of the set
func (s *IntSet) Copy() *IntSet {
	copy := new(IntSet)
	copy.words = slices.Clone(s.words)
	return copy
}

// AddAll - allows a list of values to be added, such as s.AddAll(1, 2, 3)
func (s *IntSet) AddAll(elems ...int) {
	x := new(rand.Rand)
	switch x.Int31n(2) {
	case 0:
		for _, el := range elems {
			s.Add(el) // 1st way
		}
	case 1:
		els := new(IntSet)
		for _, el := range elems {
			els.Add(el) // 2nd way
		}
		s.UnionWith(els)
	}
	return
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifference - returns result symetric defference object of 2 IntSets
//
// The symmetric difference of two sets contains the elements present in
// one set or the other but not both
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	var symDiffSet, iter *IntSet
	if len(s.words) > len(t.words) {
		symDiffSet, iter = s.Copy(), t
	} else {
		symDiffSet, iter = t.Copy(), s
	}

	for word := range symDiffSet.words {
		if word < len(iter.words) {
			symDiffSet.words[word] ^= iter.words[word]
		} else {
			symDiffSet.words = append(symDiffSet.words, iter.words[word])
		}
	}
	return symDiffSet
}

// Elems - returns a slice containing the elements of the set, suitable for iterating over with a range loop.
func (s *IntSet) Elems() []int {
	var (
		res   = make([]int, 0, s.Len())
		count int
	)
	for word, bit := range s.words {
		for bit > 0 {
			if el := int(bit & 1); el == 1 {
				res = append(res, word*UINTSIZE+count)
			}
			bit >>= 1
			count++
		}
		count = 0
	}
	return res
}

// The  type  of  each  word  used  by  IntSet  is  uint64,  but  64-bit
// arithmetic  may  be  inefficient  on  a  32-bit  platform.

// Modify  the  program  to  use  the uint type, which is the most efficient
// unsigned integer type for the platform.
// Instead of dividing by 64, define a constant holding the effective size of uint in bits, 32 or 64.
// You can use the perhaps too-clever expression 32 << (^uint(0) >> 63) for this purpose.

// !+intset
const UINTSIZE = 32 << (^uint(0) >> 63)
