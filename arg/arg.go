package arg

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/AriShashivkopanazak/truvia/api"
	"github.com/AriShashivkopanazak/truvia/save"
)

// Parse takes in the args and outputs them into api or save
func Parse() {

	// difficulty
	var rawDifficulty string
	rand.Seed(time.Now().Unix())
	randomDifficultyIndex := []string{
		"easy",
		"medium",
		"hard",
	}
	randomDifficulty := rand.Int() % len(randomDifficultyIndex)
	randomDiff := randomDifficultyIndex[randomDifficulty]
	flag.StringVar(&rawDifficulty, "d", randomDiff, "Difficulty of Question: easy, medium, or hard")

	// type of question
	var rawTypeOf string
	rand.Seed(time.Now().Unix())
	randomTypeIndex := []string{
		"&type=boolean",
		"&type=multiple",
		"ai",
	}
	typeRandom := rand.Int() % len(randomTypeIndex)
	randomType := randomTypeIndex[typeRandom]
	flag.StringVar(&rawTypeOf, "t", randomType, "Type of question, (ai, tf, mc)")

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
	var popSaved uint
	flag.UintVar(&popSaved, "p", 0, "pop existing saves from library")

	// delete saved save
	var delSaved uint
	flag.UintVar(&delSaved, "D", 0, "delete saved request, BE CAREFUL!")

	flag.Parse()

	// difficulty processing
	var difficulty string
	switch rawDifficulty {
	case "easy":
		difficulty = "&rawDifficulty=easy"
	case "medium":
		difficulty = "&rawDifficulty=medium"
	case "hard":
		difficulty = "&rawDifficulty=hard"
	case "random":
		random := "&rawDifficulty=" + randomDiff
		difficulty = random
		fmt.Println(randomDiff + " difficulty selected")
	default:
		fmt.Println("Error: unknown rawDifficulty")
		os.Exit(0)
	}

	// type of question processing
	var typeOf string
	switch rawTypeOf {
	case "tf":
		typeOf = "&type=boolean"
	case "mc":
		typeOf = "&type=multiple"
	case "ai":
		typeOf = "&type=multiple"
	case "random":
		typeOf = rawTypeOf
	default:
		fmt.Println("Error: unknown type of question")
		os.Exit(1)
	}

	if popSaved > 0 || delSaved > 0 {
		save.Save(willSave, popSaved, delSaved)
	} else {
		api.Process(difficulty, typeOf, guesses, questions)
	}
}
