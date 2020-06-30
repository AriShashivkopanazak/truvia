package api

// API Processer
func Process(difficulty string, typeOf string, questions uint, guesses uint) {
	requestOutput := apiRequest(difficulty, typeOf, questions)
	processedQuestions := apiProcess(responseOutput)
	return processedQuestions
}
