package api

import (
    "fmt"
    "os"
	"html"
)

type Question struct {
    Category string            `json:"category"`
		QuestionType string	`json:"type"`
		QuestionDifficulty string `json:"difficulty"`
		Question string `json:"question"`
		CorrectAnswer string `json:"correct_answer"`
		IncorrectAnswers []string `json:"incorrect_answers"`
}


func apiProcess(jsonResponse Response, typeOf string, guesses uint, questions uint) {
	const len uint = questions
	var questionsArray []Question

	for i := 0; i < len(responseObject.Questions); i++ {
		q = responseObject.Questions[i].Question
		q = html.UnescapeString(q) //convert html special characters to text
		if err != nil {
			fmt.Print(err.Error)
			os.Exit(1)
		}

		questionsArray = append(questionsArray, q)
	}

	return questionsArray
}
