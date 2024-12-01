package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const numWords = 854

type Game interface {
	SetUp() *GameState
	Run()
	Restart()
}

type GameState struct {
	word            string
	currentState    []bool
	lettersGuessed  string
	wrongGuessCount uint
	totalGuesses    uint
}

func newGameState() *GameState {
	game := GameState{}

	randomLine := rand.Intn(numWords) + 1
	word, err := getWord(randomLine)

	if word == "" {
		panic("Error getting word")
	}

	handleError(err)

	game.word = word
	game.totalGuesses = 5

	game.currentState = make([]bool, len([]rune(word)))
	for i := range game.currentState {
		game.currentState[i] = false
	}

	game.wrongGuessCount = 0
	game.lettersGuessed = ""

	return &game
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}
func getWord(lineNumber int) (string, error) {
	file, err := os.Open("./words.txt")
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentline := 1

	for scanner.Scan() {
		if currentline == lineNumber {
			return scanner.Text(), nil
		}

		currentline++
	}

	err = scanner.Err()

	handleError(err)
	return "", err
}

func checkForWin(game *GameState) bool {
	for _, isGuessed := range game.currentState {
		if !isGuessed {
			return false
		}
	}
	return true
}

func (game *GameState) Run() {
	game.PrintState()

	for {
		var runes []rune
		var input string

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a single character: ")
		input, _ = reader.ReadString('\n') // Read input
		input = strings.TrimSpace(input)   // Remove extra whitespace or newline

		runes = []rune(input)

		for len(runes) > 1 || strings.ContainsRune(game.lettersGuessed, runes[0]) {

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter a single character: ")
			input, _ = reader.ReadString('\n') // Read input
			input = strings.TrimSpace(input)   // Remove extra whitespace or newline

			runes = []rune(input)
		}

		guess := runes[0]

		if strings.ContainsRune(game.word, guess) {
			for index, char := range game.word {
				if char == guess {
					game.currentState[index] = true
				}
			}
			fmt.Printf("%c is in the word!\n\n", guess)
		} else {
			fmt.Printf("%c is not in the word\n\n", guess)
			game.wrongGuessCount++
		}

		game.lettersGuessed += string(guess)

		checkForWin(game)

		game.PrintState()
		if game.wrongGuessCount >= game.totalGuesses {
			fmt.Println("Game Over")
			fmt.Printf("The Word was %s\n", game.word)
			fmt.Println("Do you want to play again? (Y\\n)")

			reader := bufio.NewReader(os.Stdin)
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "n" || input == "N" {
				return
			} else {
				game.Restart()
			}
		}

	}
}

func (game *GameState) PrintState() {
	// Print the current state of the word
	fmt.Print("Word: ")
	for i, revealed := range game.currentState {
		if revealed {
			fmt.Printf("%c ", game.word[i])
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()

	// Print the letters guessed
	fmt.Printf("Letters guessed: %s\n", strings.Join(strings.Split(game.lettersGuessed, ""), ", "))

	// Print guesses left
	guessesLeft := game.totalGuesses - game.wrongGuessCount
	fmt.Printf("Guesses left: %d\n", guessesLeft)
}

func (game *GameState) Restart() {
	game = newGameState()
}

func main() {
	game := newGameState()
	game.Run()

}
