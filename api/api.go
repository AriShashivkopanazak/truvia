package api

// API Processer
func API(difficulty string, typeOf string, questions uint, guesses uint) {
	apiRequest(difficulty, typeOf, questions)
	apiProcess(typeOf, guesses, questions)
}
