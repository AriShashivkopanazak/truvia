package arg

import (
	"flag"
	"math/rand"
	"time"
)

// Parse takes in arguments
func Parse() {

	// difficulty
	var difficulty string
	rand.Seed(time.Now().Unix())
	randomDifficultyIndex := []string{
		"easy",
		"medium",
		"hard",
	}
	randomDifficulty := rand.Int() % len(randomDifficultyIndex)
	randomDiff := randomDifficultyIndex[randomDifficulty]
	flag.StringVar(&difficulty, "d", randomDiff, "Difficulty of Question: easy, medium, or hard")

	// type of question
	var typeOf string
	rand.Seed(time.Now().Unix())
	randomTypeIndex := []string{
		"ai",
		"tf",
		"mc",
	}
	typeRandom := rand.Int() % len(randomTypeIndex)
	randomType := randomTypeIndex[typeRandom]
	flag.StringVar(&typeOf, "t", randomType, "Type of question, (ai, tf, mc)")

	// guesses
	var guesses uint
	flag.UintVar(&guesses, "c", 3, "number of guesses")

	// number of questions
	var questions uint
	flag.UintVar(&questions, "q", 1, "number of questions")

	// save
	var willSave bool
	flag.BoolVar(&willSave, "s", false, "save your request")

	// load saved save
	var popSaved string
	flag.StringVar(&popSaved, "p", "null", "pop existing saves from library")

	// delete saved save
	var delSaved string
	flag.StringVar(&delSaved, "D", "null", "delete saved request, BE CAREFUL!")

	flag.Parse()

	// parse args and put them in argProcess for processing
	argProcess(difficulty, typeOf, guesses, questions, willSave, popSaved, delSaved)
}
