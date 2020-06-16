package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
		"html"
)


// A Response struct to map the Entire Response
type Response struct {
    ResponseCode    int    `json:"response_code"`
    Questions []Question `json:"results"`
}

// A Question Struct to map every trivia question to.
type Question struct {
    Category string            `json:"category"`
		QuestionType string	`json:"type"`
		QuestionDifficulty string `json:"difficulty"`
		Question string `json:"question"`
		CorrectAnswer string `json:"correct_answer"`
		IncorrectAnswers []string `json:"incorrect_answers"`
}


func processApi() {
		//make a request to the api for 10 random questions
    response, err := http.Get("https://opentdb.com/api.php?amount=10")
    if err != nil { //check for errors in the request
        fmt.Print(err.Error())
        os.Exit(1)
    }

		//read the api response
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

		//convert the json to go structs
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)

		//print the status code and the number of questions
    fmt.Println(responseObject.ResponseCode)
    fmt.Println(len(responseObject.Questions))

		//print all the questions
		var q string
    for i := 0; i < len(responseObject.Questions); i++ {
				q = responseObject.Questions[i].Question
				q = html.UnescapeString(q) //convert html special characters to text
				if err != nil {
					log.Fatal(err)
				}

        fmt.Println(q)
    }

}
