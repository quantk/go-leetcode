package main

import (
	"fmt"
	"strings"
)

func toDigit(c byte) int {
	return int(c - '0')
}

func revertString(str string) string {
	res := strings.Builder{}
	for i := len(str) - 1; i >= 0; i-- {
		res.WriteByte(str[i])
	}

	return res.String()
}

func addBinary(a string, b string) string {
	ab := []byte(a)
	bb := []byte(b)
	carry := 0
	i, j := len(a)-1, len(b)-1

	//if j > i {
	//	i, j = j, i
	//	ab, bb = bb, ab
	//}

	res := strings.Builder{}

	for i >= 0 || j >= 0 {
		curI, curJ := 0, 0
		if i >= 0 {
			curI = toDigit(ab[i])
		}
		if j >= 0 {
			curJ = toDigit(bb[j])
		}
		cur := carry + curI + curJ
		realRes := 0
		switch cur {
		case 0:
			realRes = 0
			carry = 0
		case 1:
			realRes = 1
			carry = 0
		case 2:
			realRes = 0
			carry = 1
		case 3:
			realRes = 1
			carry = 1
		}
		res.WriteString(fmt.Sprint(realRes))
		i--
		j--
	}

	if carry > 0 {
		res.WriteString("1")
	}

	result := revertString(res.String())

	return result
}
