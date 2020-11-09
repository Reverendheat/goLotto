package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Lottery struct {
	stateNumbers  [5]int
	playerNumbers [5]int
	rounds        int
	winner        bool
}

func (l Lottery) play() {
RoundLoop:
	for n := 1; n <= l.rounds; n++ {
		fmt.Println("Round number", n)
		//TODO: Probably don't need to generate 100 numbers just to cherry pick 10
		randomStateNumbers := rand.Perm(100)
		l.stateNumbers = [5]int{randomStateNumbers[0], randomStateNumbers[1], randomStateNumbers[2], randomStateNumbers[3], randomStateNumbers[4]}
		fmt.Println("States numbers are: ", l.stateNumbers)
		for i, v := range l.stateNumbers {
			if v != l.playerNumbers[i] {
				continue RoundLoop
			} else {
				continue
			}
		}
		fmt.Println("Probably won't happen again, but you win! Go buy a real lotto ticket!")
		return
	}
	fmt.Println("Shocking, you lose!")
}

func validate(a [5]int) bool {
	fmt.Println("Validating...")
	for i := range a {
		if (a[i] <= 100) && (a[i] >= 1) {
			continue
		} else {
			fmt.Println("Failed: Enter numbers between 1-100, separated by spaces. '11 22 33 44 55'.")
			return false
		}
	}
	fmt.Println("Valid!")
	return true
}

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println("Enter 5 numbers between 1-100, separated by spaces.")
	var playerIntInput [5]int
	fmt.Scanln(&playerIntInput[0], &playerIntInput[1], &playerIntInput[2], &playerIntInput[3], &playerIntInput[4])

	fmt.Println("Enter the number of weeks you would like to play.")
	var playerRoundsInput int
	fmt.Scanln(&playerRoundsInput)

	//Validate the array of integers the player input
	valid := validate(playerIntInput)
	if valid == false {
		return
	}

	player := Lottery{
		//stateNumbers:  [5]int{randomStateNumbers[0], randomStateNumbers[1], randomStateNumbers[2], randomStateNumbers[3], randomStateNumbers[4]},
		playerNumbers: [5]int{playerIntInput[0], playerIntInput[1], playerIntInput[2], playerIntInput[3], playerIntInput[4]},
		rounds:        playerRoundsInput,
	}
	player.play()
}
