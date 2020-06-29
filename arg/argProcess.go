package arg

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/AriShashivkopanazak/truvia/api"
	"github.com/AriShashivkopanazak/truvia/save"
)

func argProcess(rawDifficulty string, rawTypeof string, guesses uint, questions uint, willSave bool, rawPopSaved string, rawDelSaved string) {

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
		rand.Seed(time.Now().Unix())
		randomrawDifficultyIndex := []string{
			"&rawDifficulty=easy",
			"&rawDifficulty=medium",
			"&rawDifficulty=hard",
		}
		randomrawDifficulty := rand.Int() % len(randomrawDifficultyIndex)
		difficulty = randomrawDifficultyIndex[randomrawDifficulty]
	default:
		fmt.Println("Error: unknown rawDifficulty")
		os.Exit(0)
	}

	// type of question processing
	var typeOf string
	switch rawTypeof {
	case "tf":
		typeOf = "&type=boolean"
	case "mc":
		typeOf = "&type=multiple"
	case "ai":
		typeOf = "&type=multiple"
	case "random":
		rand.Seed(time.Now().Unix())
		randomTypeIndex := []string{
			"ai",
			"tf",
			"mc",
		}
		typeRandom := rand.Int() % len(randomTypeIndex)
		typeOf = randomTypeIndex[typeRandom]
	default:
		fmt.Println("Error: unknown type of question")
		os.Exit(1)
	}

	var popSaved uint
	if rawPopSaved != "null" {
		intPopSaved, err := strconv.Atoi(rawPopSaved)
		if err != nil {
			fmt.Println("Error: Please use a valid save name")
			os.Exit(1)
		}
		popSaved = uint(intPopSaved)
	}

	var delSaved uint
	if rawDelSaved != "null" {
		intDelSaved, err := strconv.Atoi(rawDelSaved)
		if err != nil {
			fmt.Println("Error: Please use a valid save name")
			os.Exit(1)
		}
		delSaved = uint(intDelSaved)
	}
	if willSave == true || popSaved > 0 || delSaved > 0 {
		save.Save(willSave, popSaved, delSaved)
	} else {
		api.API(difficulty, typeOf, questions, guesses)
	}
}
