package api

import (
	"html"
)

type Question struct {
	Category           string   `json:"category"`
	QuestionType       string   `json:"type"`
	QuestionDifficulty string   `json:"difficulty"`
	Question           string   `json:"question"`
	CorrectAnswer      string   `json:"correct_answer"`
	IncorrectAnswers   []string `json:"incorrect_answers"`
}

func apiProcess(jsonResponse Response) []Question {
	//const len uint = questions
	var questionsArray []Question
	//var q Question

	for i := 0; i < len(jsonResponse.Questions); i++ {
		questionString := jsonResponse.Questions[i].Question
		questionString = html.UnescapeString(questionString) //convert html special characters to text

		q := jsonResponse.Questions[i]
		q.Question = questionString

		questionsArray = append(questionsArray, q)
	}

	return questionsArray
}
