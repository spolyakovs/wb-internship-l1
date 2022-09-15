package main

/*
line "justString = v[:100]" can be problematic
because all strings encoded in Unicode, 1 letter could take 1-4 bytes
and "v[:100]" takes 100 BYTES, not letters

lower example of more correct implementation

func someFunc() {
	v := createHugeString(1 << 10)
	vRunes := []rune(v)
	justRunes := vRunes[:100]
	justString = string(justRunes)
}

if "v" is huge, instead of lines 12-13 could use utf8.DecodeRuneInString
to get runes from string one by one in cycle

*/
