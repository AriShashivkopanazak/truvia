package api

import "github.com/AriShashivkopanazak/truvia/ask"

func apiProcess(typeOf string, guesses uint, questions uint) {
	ask.Question(typeOf, guesses, questions)
}
