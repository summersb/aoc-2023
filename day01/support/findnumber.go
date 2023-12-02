package support

import (
	"strings"
)

func FindFirstNumber(txt string) int {
    vals := []string{"one","1","two","2","three","3","four","4","five","5","six","6","seven","7","eight","8","nine","9"}
    number := 0
    index := len(txt)

    for i, s := range vals {
        idx := strings.Index(txt, s)
        if idx > -1 && idx < index {
            index = idx
            number = i / 2 + 1
        }
    }

    return number 
}

func FindLastNumber(txt string) int {
    vals := []string{"one","1","two","2","three","3","four","4","five","5","six","6","seven","7","eight","8","nine","9"}
    number := -1
    index := -1

    for i, s := range vals {
        idx := strings.LastIndex(txt, s)
        if idx > -1 && idx > index {
            index = idx
            number = i / 2 + 1
        }
    }

    return number 
}
