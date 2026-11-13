package main

import (
	"fmt"
)

func gameLoop() {
	for {
		var numberOfGuesses uint8 = 1
		var targetNumber = GetRandomInt(1, 100)

		printDifficultyPrompt()

		var difficulty Difficulty
		if _, err := fmt.Scanf("%d", &difficulty); err != nil || !difficulty.IsValid() {
			printInvalidInputErr()
			continue
		}

		printDifficultyChoiceDialogue(difficulty)
		guessLimit := difficulty.Limit()

		for {
			printGuessPrompt()

			var guess uint8
			if _, err := fmt.Scanf("%d", &guess); err != nil || guess < 1 || guess > 100 {
				printInvalidInputErr()
				continue
			}

			if guess == targetNumber {
				printCongratsWithNumberOfGuesses(numberOfGuesses)
				break
			} else if guess < targetNumber {
				printIncorrectGuessDialogue(IndicatorGreater, guess)
			} else {
				printIncorrectGuessDialogue(IndicatorLess, guess)
			}

			numberOfGuesses++
			if numberOfGuesses > guessLimit {
				printGuessLimitExceededDialogue(targetNumber)
				break
			}
		}

		printContinuePrompt()
		var doContinue string
		if _, err := fmt.Scanf("%s", &doContinue); err != nil || (doContinue != "y" && doContinue != "n") {
			printInvalidInputErr()
			break
		}

		if doContinue == "n" {
			printGoodbye()
			break
		}
	}
}

func main() {
	printWelcome()
	gameLoop()
}
