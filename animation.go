package hangmanlesbg

import (
	"fmt"
	"os"
	"time"
)

func frame(r []byte) []string {
	word := ""
	line := 0
	result := []string{}

	for _, i := range r {
		if line == 8 {
			line = 0
			result = append(result, word)
			word = ""
			continue
		}
		if string(i) == "\n" {
			line++
		}
		word += string(i)
	}
	return append(result, word)
}

func HangmanLoseAnimation() {
	text, _ := os.ReadFile("jeu/hangman_frame.txt")
	frames := frame(text)
	frames = append(frames, frames[1])
	fmt.Print(frames[0])
	index := 0

	for i := 0; i < 20; i++ {
		fmt.Printf("\033[8A")
		fmt.Printf("%s", frames[index])
		time.Sleep(500 * time.Millisecond)
		index++
		if index == 4 {
			index = 0
		}
	}
	fmt.Println("\033[1;31mGAME OVER\033[0m")
}

func HangmanWinAnimation() {
	text, _ := os.ReadFile("jeu/hangman_frame2.txt")
	frames := frame(text)
	frames = append(frames, frames[1])
	fmt.Print(frames[0])
	index := 0

	for i := 0; i < 20; i++ {
		fmt.Printf("\033[8A")
		fmt.Printf("%s", frames[index])
		time.Sleep(300 * time.Millisecond)
		index++
		if index == 4 {
			index = 0
		}
	}
	fmt.Println("\033[1;32mYou saved Jose ! \U0001F60A\033[0m")
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")
}

func GuilloLoseAnimation() {
	text, _ := os.ReadFile("jeu/guillotineAnime.txt")
	frames := frame(text)
	frames = append(frames, frames[1])
	fmt.Print(frames[0])
	index := 0

	for i := 0; i < 3; i++ {
		fmt.Printf("\033[8A")
		fmt.Printf("%s", frames[index])
		time.Sleep(500 * time.Millisecond)
		index++
		if index == 4 {
			index = 0
		}
	}
	fmt.Println("\033[1;31mGAME OVER\033[0m")
}

func GuilloWinAnimation() {
	text, _ := os.ReadFile("jeu/guillotineAnime2.txt")
	frames := frame(text)
	frames = append(frames, frames[1])
	fmt.Print(frames[0])
	index := 0

	for i := 0; i < 20; i++ {
		fmt.Printf("\033[8A")
		fmt.Printf("%s", frames[index])
		time.Sleep(500 * time.Millisecond)
		index++
		if index == 4 {
			index = 0
		}
	}
	fmt.Println("\033[1;32mYou saved Jose ! \U0001F60A\033[0m")
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")
}
