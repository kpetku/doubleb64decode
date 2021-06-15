package main

import (
	"bytes"
	"unicode"
)

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var i int
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i = bytes.IndexByte(data, doubleQuote); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if i = bytes.IndexByte(data, singleQuote); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

// https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
