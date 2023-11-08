package hangmanlesbg

import (
	"fmt"
	"math/rand"
	"os"
)

func StartHard(s []string) []string {
	var index int
	word := []string{}
	for _, i := range s {
		word = append(word, i)
	}
	for i := 0; i < len(word); i++ {
		index = rand.Intn(len(word))
		for word[index] == "_" {
			index = rand.Intn(len(s))
		}
		word[index] = "_"
	}
	return word
}

func StartEasy(s []string) []string {
	var index int
	word := []string{}
	for _, i := range s {
		word = append(word, i)
	}
	for i := 0; i < ((len(word) / 2) + 1); i++ {
		index = rand.Intn(len(word))
		for word[index] == "_" {
			index = rand.Intn(len(s))
		}
		word[index] = "_"
	}
	return word
}

func Turn(s, answer []string) ([]string, int, string) {
	var input string
	var cd int
	fmt.Printf("LETTER : ")
	_, err := fmt.Scan(&input)
	fmt.Print("\033[H\033[2J")
	if err != nil {
		return []string(nil), 0, ""
	}
	if len(input) != 1 {
		return []string(nil), 0, ""
	}
	if input == "+" {
		os.Exit(0)
	}
	l := Check(input, answer)
	if len(l) == 0 {
		fmt.Printf("You type %s, \033[1;31mWrong !\033[0m\n \n", input)
		cd = -1
	} else {
		fmt.Printf("You type %s, \033[1;32mGreat !\033[0m\n \n", input)
		cd = 0
	}
	for _, i := range l {
		s[i] = answer[i]
	}
	return s, cd, string(input)
}

func Check(s string, answer []string) []int {
	l := []int{}
	for i := 0; i < len(answer); i++ {
		if s == answer[i] {
			l = append(l, i)
		}
	}
	return l
}

func Victory(s, answer []string) bool {
	for f, i := range answer {
		if s[f] != i {
			return false
		}
	}
	fmt.Print("\033[H\033[2J")
	fmt.Printf("\033[1;32m%s\033[0m\n \n", s)
	fmt.Println("\033[1;32mYou got this one Congrats !\033[0m")
	return true
}
