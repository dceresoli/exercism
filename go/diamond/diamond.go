package diamond

//package main
//import "fmt"

import (
	"errors"
	"strings"
)

const testVersion = 1

func Gen(c byte) (string, error) {
	if c < 'A' || c > 'Z' {
		return "", errors.New("char out of range")
	}

	half := int(c - 'A')
	size := 2*half + 1
	square := make([]byte, size*size)
	for i := 0; i < size*size; i++ {
		square[i] = ' '
	}

	for i := 0; i <= half; i++ {
		square[(half+i)+size*i] = byte('A' + i)
		square[(half-i)+size*i] = byte('A' + i)
		square[(half+i)+size*(size-i-1)] = byte('A' + i)
		square[(half-i)+size*(size-i-1)] = byte('A' + i)
	}

	res := make([]string, size)
	for i := 0; i < size; i++ {
		res[i] = string(square[i*size : (i+1)*size])
	}
	return strings.Join(res, "\n") + "\n", nil
}

//func main() {
//	dia, _ := Gen('C')
//	for _, line := range strings.Split(dia, "\n") {
//		fmt.Printf("\"%s\"\n", line)
//	}
//}
