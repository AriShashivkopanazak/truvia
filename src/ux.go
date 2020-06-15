package main

// UX for Truvia

import (
	"flag"
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
	var randomDifficulty string // make random difficulty generator
	flag.StringVar(&difficulty, "d", randomDifficulty, "Difficulty of Questions, defaults to random")

	// type of question flag
	var typeQuestion string
	var randomType string // make random string generator
	flag.StringVar(&typeQuestion, "t", randomType, "Type of question, (true/false, multiple choice, ai matched answer, or random), defaults to random")

	// chances
	var chances string
	flag.StringVar(&chances, "c", "3", "number of chances, defaults to 3 chances")

	// number of questions
	var questionQuanity string
	flag.StringVar(&questionQuanity, "n", "1", "number of questions, defaults to 1 question request")

	// cache
	var cached string
	flag.StringVar(&cached, "m", "false", "download your json file and save it, test case")

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
	}

	// process types of questions
	// process chances
	// process number of questions
	// process cache
	if cached == "false" {
		// do not save json as file
	} else {
		// have user type in valid path to save and save as 0.json, and move other saves to one below, previous 0.json becomes 1.json
	}

	// process saved cache
	// process deleted cache
}
