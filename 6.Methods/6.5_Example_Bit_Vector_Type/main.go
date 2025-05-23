package main

import (
	"bytes"
	"fmt"
)

//* A set represented by a map[T]bool is very flexible but,
//* for certain problems, a specialized representation may outperform it.

//? For example, in domains such as dataflow analysis
//? where set elements are small non-negative integers, sets have many elements,
//? and set operations like union and intersection are common, a bit vector is ideal.

//! A  bit  vector  uses  a  slice  of  unsigned  integer  values  or  “words,”
//! each  bit  of  which represents a possible element of the set.
//! The set contains i if the i-th bit is set.

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// ! Since each word has 64 bits, to locate the bit for x,
// ! we use the quotient x/64 as the word  index
// ! the  remainder  x%64  as  the  bit  index  within  that  word
//
// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&
		(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// ! The UnionWith  operation  uses  the  bitwise  OR  operator  |
// ! to  compute  the  union  64 elements at a time
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
		for j := range 64 {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// IntSet value does not have a String method, occasionally leading to surprises like this

//IntSet pointer, which does have a String method.
//* fmt.Println(&x)         // "{1 9 42 144}"

// we call String() on an IntSet variable;
// the compiler inserts the implicit & operation, giving us a pointer,
// which has the String method.
//* fmt.Println(x.String()) // "{1 9 42 144}"

// the  IntSet  value  does  not  have  a  String  method,
// fmt.Println prints the representation of the struct instead.
//* fmt.Println(x)          // "{[4398046511618 0 65536]}"

//Exercise 6.1 - Exercise  6.5 implemented in DataTypes dir
