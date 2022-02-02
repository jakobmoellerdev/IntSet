package main

func CheckBitInWord(word *uint, i *uint) bool {
	return ((*word >> *i) & 1) == 1
}

func ClearBitInWord(word *uint, i *uint) {
	*word &= ^(1 << *i)
}

func SetBitInWord(word *uint, i *uint) {
	*word |= 1 << *i
}
