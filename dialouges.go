package main

import "fmt"

const DIALOGUE_WELCOME = `Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have multiple chances to guess the correct number.`

func printWelcome() {
	fmt.Println(DIALOGUE_WELCOME)
}

const DIALOGUE_DIFFICULTY_LEVEL_PROMPT = `Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: `

func printDifficultyPrompt() {
	fmt.Printf("\n%s", DIALOGUE_DIFFICULTY_LEVEL_PROMPT)
}

func printGuessPrompt() {
	fmt.Print("\nEnter your guess: ")
}

func printDifficultyChoiceDialogue(d Difficulty) {
	fmt.Printf("\nGreat! You have selected the %s difficulty level.\nLet's start the game!\n", d)
}

func printIncorrectGuessDialogue(i Indicator, g uint8) {
	fmt.Printf("\nIncorrect! The number is %s than %d.\n", i, g)
}

func printCongratsWithNumberOfGuesses(numberOfGuesses uint8) {
	fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts.\n\n", numberOfGuesses)
}

func printGuessLimitExceededDialogue(target uint8) {
	fmt.Printf("\nyou exceeded the guess limit\nthe number was %d\n", target)
}

func printContinuePrompt() {
	fmt.Print("Do you want to continue (y/n): ")
}

func printInvalidInputErr() {
	fmt.Printf("\n%s\n", ErrInvalidInput)
}

func printGoodbye() {
	fmt.Println("See you next time")
}
