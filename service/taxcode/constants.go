package taxcode

import (
	"bytes"
	"golang.org/x/exp/slices"
)

var charMonthMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"H": 6,
	"L": 7,
	"M": 8,
	"P": 9,
	"R": 10,
	"S": 11,
	"T": 12,
}

var monthCharMap = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "H",
	7:  "L",
	8:  "M",
	9:  "P",
	10: "R",
	11: "S",
	12: "T",
}

var vowelsSlice = []rune{'A', 'E', 'I', 'O', 'U'}

var controlCharMap = map[int]string{
	0:  "A",
	1:  "B",
	2:  "C",
	3:  "D",
	4:  "E",
	5:  "F",
	6:  "G",
	7:  "H",
	8:  "I",
	9:  "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "O",
	15: "P",
	16: "Q",
	17: "R",
	18: "S",
	19: "T",
	20: "U",
	21: "V",
	22: "W",
	23: "X",
	24: "Y",
	25: "Z",
}

var evenSumMap = map[string]int{
	"0": 0,
	"A": 0,
	"1": 1,
	"B": 1,
	"2": 2,
	"C": 2,
	"3": 3,
	"D": 3,
	"4": 4,
	"E": 4,
	"5": 5,
	"F": 5,
	"6": 6,
	"G": 6,
	"7": 7,
	"H": 7,
	"8": 8,
	"I": 8,
	"9": 9,
	"J": 9,
	"K": 10,
	"L": 11,
	"M": 12,
	"N": 13,
	"O": 14,
	"P": 15,
	"Q": 16,
	"R": 17,
	"S": 18,
	"T": 19,
	"U": 20,
	"V": 21,
	"W": 22,
	"X": 23,
	"Y": 24,
	"Z": 25,
}

var oddSumMap = map[string]int{
	"0": 1,
	"A": 1,
	"1": 0,
	"B": 0,
	"2": 5,
	"C": 5,
	"3": 7,
	"D": 7,
	"4": 9,
	"E": 9,
	"5": 13,
	"F": 13,
	"6": 15,
	"G": 15,
	"7": 17,
	"H": 17,
	"8": 19,
	"I": 19,
	"9": 21,
	"J": 21,
	"K": 2,
	"L": 4,
	"M": 18,
	"N": 20,
	"O": 11,
	"P": 3,
	"Q": 6,
	"R": 8,
	"S": 12,
	"T": 14,
	"U": 16,
	"V": 10,
	"W": 22,
	"X": 25,
	"Y": 24,
	"Z": 23,
}

var case0 = func(vowels string, consonants string) string {
	if len(vowels) > 2 {
		return vowels[0:3]
	}
	if len(vowels) == 2 {
		return vowels + "X"
	}
	if len(vowels) == 1 {
		return vowels + "XX"
	}
	return "XXX"
}

var case1 = func(vowels string, consonants string) string {
	if len(vowels) >= 2 {
		return consonants + vowels[0:2]
	}
	if len(vowels) == 1 {
		return consonants + vowels + "X"
	}
	return consonants + "XX"
}

var case2 = func(vowels string, consonants string) string {
	if len(vowels) >= 2 {
		return consonants + vowels[0:1]
	}
	return consonants + "X"
}

var case3 = func(vowels string, consonants string) string {
	return consonants
}

var surnameCaseDefault = func(vowels string, consonants string) string {
	return consonants[0:3]
}

var nameCaseDefault = func(vowels string, consonants string) string {
	return consonants[0:1] + consonants[2:4]
}

var surnameFunctionMap = map[int]func(vowels string, consonants string) string{
	0: case0,
	1: case1,
	2: case2,
}

var nameFunctionMap = map[int]func(vowels string, consonants string) string{
	0: case0,
	1: case1,
	2: case2,
	3: case3,
}

var consonants = func(word string) string {
	var consonants bytes.Buffer
	for _, char := range word {
		if !slices.Contains(vowelsSlice, char) {
			consonants.WriteRune(char)
		}
	}
	return consonants.String()
}

var vowels = func(word string) string {
	var consonants bytes.Buffer
	for _, char := range word {
		if slices.Contains(vowelsSlice, char) {
			consonants.WriteRune(char)
		}
	}
	return consonants.String()
}
