package api

// API Processer
func Process(difficulty string, typeOf string, questions uint, guesses uint) []Question {
	requestOutput := apiRequest(difficulty, typeOf, questions)
	processedQuestions := apiProcess(requestOutput)
	return processedQuestions
}
