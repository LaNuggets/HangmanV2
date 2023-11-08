package hangmanlesbg

import (
	"fmt"
	hangmanlesbg "hangmanlesbg/jeu"
	"math/rand"
	"os"
	"slices"
	"strconv"
	//"os/exec"
)

func main() {
	score := 0
	player := []string{}

	for {
		fmt.Print("\033[H\033[2J")
		death, level := hangmanlesbg.Menu()
		fmt.Print("\033[1;35mNew Game\033[0m \n \n \n")

		for {
			listLetter := []string{}
			text := "jeu/words" + strconv.Itoa(rand.Intn(4)) + ".txt"
			answer := hangmanlesbg.SelectWord(text)
			if level == "e" {
				player = hangmanlesbg.StartEasy(answer)
			} else if level == "h" {
				player = hangmanlesbg.StartHard(answer)
			}
			hangmanlesbg.Display(player)
			countdown := 10
			for hangmanlesbg.Victory(player, answer) == false && countdown != 0 {
				fmt.Printf("\033[1;35mRemaining Trials : \033[0m%d             ", countdown)
				fmt.Printf("\033[1;35mletters already tried :\033[0m ")
				hangmanlesbg.DisplayLetterUsed(listLetter)
				mot, score, lettre := hangmanlesbg.Turn(player, answer)
				if score == -1 && !slices.Contains(listLetter, lettre) {
					countdown += score
					listLetter = append(listLetter, lettre)
				}
				if death == "c" {
					hangmanlesbg.DisplayHangman(countdown)
				} else if death == "f" {
					hangmanlesbg.DisplayGuillo(countdown)
				}
				hangmanlesbg.Display(mot)
			}
			if countdown == 0 {
				fmt.Print("\033[H\033[2J")
				fmt.Println("\nThe answer was : ")
				hangmanlesbg.Display(answer)
				fmt.Printf("\033[1;31mJose is Dead, Sorry \033[1;37m\u2620\033[0m\n")
				if death == "c" {
					hangmanlesbg.HangmanLoseAnimation()
				}
				if death == "f" {
					hangmanlesbg.GuilloLoseAnimation()
				}
				os.Exit(0)
			} else {
				if death == "c" {
					hangmanlesbg.HangmanWinAnimation()
				} else if death == "f" {
					hangmanlesbg.GuilloWinAnimation()
				}
				score++
				fmt.Printf("Score : %d \n", score)
				fmt.Println("\033[1;35mNew Game\033[0m")
				fmt.Printf("\n \n \n \n")
			}
		}
	}
}
