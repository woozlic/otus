package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func getRune(s string, index int) int32 {
	for i, cur := range s {
		if i == index {
			return cur
		}
	}
	return -1
}

func Unpack(s string) (string, error) {
	var b strings.Builder
	for i, cur := range s {
		var next int32
		var curIsNum, nextIsNum bool
		if i < len(s)-1 {
			next = getRune(s, i+1)
		}
		_, errCur := strconv.Atoi(string(cur))
		if errCur == nil {
			curIsNum = true
		}
		countNext, errNext := strconv.Atoi(string(next))
		if errNext == nil {
			nextIsNum = true
		}
		switch {
		case (curIsNum && nextIsNum) || (i == 0 && curIsNum):
			return "", ErrInvalidString
		case !curIsNum && nextIsNum:
			b.WriteString(strings.Repeat(string(cur), countNext))
		case !curIsNum && !nextIsNum:
			b.WriteRune(cur)
		default:
			continue
		}
	}
	return b.String(), nil
}
