package ux

import (
	"flag"
	"math/rand"
	"time"
)

func argParse() {

	// difficulty flag
	var rawDifficulty string
	rand.Seed(time.Now().Unix())
	randomDifficultyIndex := []string{
		"easy",
		"medium",
		"hard",
	}
	randomDifficulty := rand.Int() % len(randomDifficultyIndex)
	flag.StringVar(&rawDifficulty, "d", randomDifficultyIndex[randomDifficulty], "Difficulty of Question: easy, medium, or hard")

	// type of question flag
	var rawQuestion string
	rand.Seed(time.Now().Unix())
	randomTypeIndex := []string{
		"ai",
		"tf",
		"mc",
	}
	randomType := rand.Int() % len(randomTypeIndex)
	flag.StringVar(&rawQuestion, "t", randomTypeIndex[randomType], "Type of question, (ai, tf, mc)")

	// chances
	var numGuess int
	flag.IntVar(&numGuess, "c", 3, "number of guesses")

	// number of questions
	var numQuestion int
	flag.IntVar(&numQuestion, "q", 1, "number of questions")

	// save
	var saved bool
	flag.BoolVar(&saved, "s", false, "save your request")

	// load saved save
	var popSaved string
	flag.StringVar(&popSaved, "p", "null", "pop existing saves from library")

	// delete saved save
	var delSaved string
	flag.StringVar(&delSaved, "D", "null", "delete saved request, BE CAREFUL!")

	flag.Parse()
}

func processArgs() {

	argParse()

	// // process difficulty
	// if rawDifficulty == "easy" {
	// 	// give easy request
	// } else if rawDifficulty == "medium" {
	// 	// give medium request
	// } else if rawDifficulty == "hard" {
	// 	// give hard request
	// } else if rawDifficulty == "random" {
	// 	rawDifficulty = randomDifficultyIndex[randomDifficulty]
	// } else {
	// 	fmt.Println("unknown difficulty")
	// 	os.Exit(1)
	// }

	// // process types of questions
	// if rawQuestion == "tf" {
	// 	// request true/false
	// } else if rawQuestion == "mc" {
	// 	// request multiple choice
	// } else if rawQuestion == "ai" {
	// 	// request multiple choice
	// 	// request ai to work on this
	// 	//ai()
	// } else if rawQuestion == "random" {
	// 	// do random value
	// 	rawQuestion = randomTypeIndex[randomType]
	// } else {
	// 	fmt.Println("unknown type of question")
	// 	os.Exit(1)
	// }

	// guessQuanity := numGuess
	// // converted string number into integer
	// fmt.Printf("using %v chances\n", guessQuanity)

	// questionQuanity := numQuestion
	// // converted string number into integer
	// fmt.Printf("using %v questions\n", questionQuanity)

	// willSave := saved
	// if willSave == true {
	// 	// save file
	// 	fmt.Println("saving query...")
	// }

	// // process saved save
	// popSave, errPop := strconv.Atoi(popSaved)
	// if errPop != nil && popSaved != "null" {
	// 	fmt.Println("use a valid save number")
	// 	os.Exit(1)
	// }
	// if popSaved != "null" {
	// 	// converted number string into a intefer
	// 	fmt.Printf("loading %v from saves\n", popSave)
	// }

	// // process deleted save
	// delSave, errDel := strconv.Atoi(delSaved)
	// if errDel != nil && delSaved != "null" {
	// 	fmt.Println("use a valid save number to delete")
	// 	os.Exit(1)
	// }
	// if delSaved != "null" {
	// 	// converted number string into a integer
	// 	fmt.Printf("deleting %v from saves\n", delSave)
	// }
}
