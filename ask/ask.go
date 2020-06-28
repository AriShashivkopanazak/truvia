package ask

import "github.com/AriShashivkopanazak/truvia/ai"

// Question asks
func Question(typeOf string, guesses uint, questions uint) {

	var i uint
	for i = 0; i == guesses; i++ {
		if typeOf == "ai" {
			ai.AI()
		} else {

		}
	}
}
