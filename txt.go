package hangmanlesbg

import (
	"fmt"
	"math/rand"
	"os"
)

func SelectWord(s string) []string {
	content, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("file error")
		return []string(nil)
	}
	list := ByteToString(content)
	mot := list[rand.Intn(len(list)-1)]
	mot = strings.TrimSpace(mot)
	run := []rune(mot)
	str := []string{}
	for i := 0; i < len(mot); i++ {
		str = append(str, string(run[i]))
	}
	return str
}

func ByteToString(b []byte) []string {
	word := ""
	result := []string{}
	for _, i := range b {
		if string(i) == "\n" {
			result = append(result, word)
			word = ""
		} else {
			word += string(i)
		}
	}
	return result
}
