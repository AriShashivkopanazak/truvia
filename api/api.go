package api

// API Processer
func API(difficulty string, typeOf string, questions uint, guesses uint) {
	apiRequest(difficulty, typeOf, questions)
	apiProcess(typeOf, guesses, questions)
}

// Process test
func Process(difficulty string, typeOf string, guesses uint, questions uint) {

}
