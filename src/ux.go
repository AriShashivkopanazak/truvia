package main

// UX for Truvia

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Parses args
func argParse() {

	// Random Example
	// // NOTE: figure out why this is not actully random
	// rand.Seed(34)
	// randomDifficulty := []string{
	// 	"easy",
	// 	"medium",
	// 	"hard",
	// }
	// random := randomDifficulty[rand.Intn(len(randomDifficulty))]

	// difficulty flag
	var difficulty string
	var randomDifficulty string = "easy" // make random difficulty generator
	flag.StringVar(&difficulty, "d", randomDifficulty, "Difficulty of Questions, defaults to random")

	// type of question flag
	var typeQuestion string
	var randomType string = "tf" // make random string generator
	flag.StringVar(&typeQuestion, "t", randomType, "Type of question, (true/false, multiple choice, ai matched answer, or random), defaults to random")

	// chances
	var chances string
	flag.StringVar(&chances, "c", "3", "number of chances, defaults to 3 chances")

	// number of questions
	var questionQuanity string
	flag.StringVar(&questionQuanity, "q", "1", "number of questions, defaults to 1 question request")

	// cache
	var cached string
	flag.StringVar(&cached, "s", "false", "download your json file and save it, test case")

	// load saved cache
	var popCached string
	flag.StringVar(&popCached, "p", "0", "pop existing saves from library, defaults to first most previous save")

	// delete saved cache
	var delCached string
	flag.StringVar(&delCached, "D", "0", "delete saved request, defaults to most previous save, BE CAREFUL!")

	flag.Parse()

	// process difficulty
	if difficulty == "easy" {
		// give easy request
	} else if difficulty == "medium" {
		// give medium request
	} else if difficulty == "hard" {
		// give hard request
	} else {
		fmt.Println("unknown difficulty")
		os.Exit(1)
	}

	// process types of questions
	if typeQuestion == "tf" {
		// request true/false
	} else if typeQuestion == "mc" {
		// request multiple choice
	} else if typeQuestion == "ai" {
		// request multiple choice
		// request ai to work on this
		//ai()
	} else if typeQuestion == "random" {
		// do random value
	} else {
		fmt.Println("unknown type of question")
		os.Exit(1)
	}

	// process chances
	numChances, errChances := strconv.Atoi(chances)
	if errChances != nil {
		fmt.Println("chances must use a integer value")
		os.Exit(1)
	}

	// converted string number into integer
	fmt.Printf("using %v chances\n", numChances)

	// process number of questions
	numQuestion, errQuestion := strconv.Atoi(questionQuanity)
	if errQuestion != nil {
		fmt.Println("question must use a integer value")
		os.Exit(1)

	}

	// converted string number into integer
	fmt.Printf("using %v questions\n", numQuestion)

	// process cache
	willCache, errCache := strconv.ParseBool(cached)
	if errCache != nil {
		fmt.Println("save arg must use boolean value")
		os.Exit(1)
	}

	if willCache == true {
		// save file
		fmt.Println("saving query...")
	}

	// process saved cache

	// process deleted cache
}
