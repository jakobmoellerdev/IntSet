package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

// wordSize checks if we get 64 bit out of a zero-initialized uint, meaning we are either on a 64 bit platform or 32.
func (s *IntSet) wordSize() int {
	return 32 << (^uint(0) >> 63)
}

// wordIndexAndBitInWord returns the index of the word as well as the position of the bit in that word for a given number
func (s *IntSet) wordIndexAndBitInWord(x int) (int, uint) {
	wordSize := s.wordSize()
	return x / wordSize, uint(x % wordSize)
}

// Add adds a new integer to the set
func (s *IntSet) Add(x int) {
	wordIndex, bit := s.wordIndexAndBitInWord(x)
	s.ensureCapacity(wordIndex)
	SetBitInWord(&s.words[wordIndex], &bit)
}

// Remove removes an integer from the set
func (s *IntSet) Remove(x int) {
	wordIndex, bit := s.wordIndexAndBitInWord(x)
	ClearBitInWord(&s.words[wordIndex], &bit)
}

// Exists checks for existence of a number in the set
func (s *IntSet) Exists(x int) bool {
	wordIndex, bit := s.wordIndexAndBitInWord(x)
	return wordIndex < len(s.words) && CheckBitInWord(&s.words[wordIndex], &bit)
}

// ensureCapacity checks if there is enough capacity for a given wordIndex
func (s *IntSet) ensureCapacity(wordIndex int) {
	for wordIndex >= len(s.words) {
		s.words = append(s.words, 0)
	}
}

// Union adds all elements from the given set to the set while ignoring duplicates
func (s *IntSet) Union(t *IntSet) *IntSet {
	cs := s.Copy()
	for i, word := range t.words {
		if i < len(cs.words) {
			cs.words[i] |= word
		} else {
			cs.words = append(cs.words, word)
		}
	}
	return cs
}

// Intersection returns all elements contained in both sets
func (s *IntSet) Intersection(t *IntSet) *IntSet {
	var cs, ct = s.Copy(), t.Copy()
	for i := 0; len(cs.words) != len(ct.words); i++ {
		if len(cs.words) > len(ct.words) {
			ct.words = append(ct.words, cs.words[i])
		} else {
			cs.words = append(cs.words, ct.words[i])
		}
	}
	for i, word := range cs.words {
		ct.words[i] &= word
	}
	return ct
}

// SymmetricDifference returns all elements contained in either set but not both
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	cs := t.Copy()
	for i, word := range s.words {
		if len(cs.words) <= i {
			cs.words = append(cs.words, 0)
		}
		cs.words[i] ^= word
	}
	return cs
}

// AddAll adds all elements from the given values to the set
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Clear clears the slice
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy uses AppendPreCapped for an avoid-zero, avoid-grow deep copy.
func (s *IntSet) Copy() *IntSet {
	return &IntSet{append(make([]uint, 0, len(s.words)), s.words...)}
}

// Len returns the length of the Set
func (s *IntSet) Len() int {
	count := 0
	wordSize := s.wordSize()
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	wordSize := s.wordSize()
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
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

func (s *IntSet) Elems() []int {
	elems := make([]int, 0, s.Len())
	wordSize := s.wordSize()
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return elems
}
