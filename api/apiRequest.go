package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Response raw json from api request
type Response struct {
	ResponseCode int        `json:"response_code"`
	Questions    []Question `json:"results"`
}

func apiRequest(difficulty string, typeOf string, questions uint) Response {
	var responseObject Response
	var requestURL string

	//questionString := strconv.FormatUint(questions, 10)

	requestURL = "https://opentdb.com/api.php" + difficulty + typeOf // + questionString

	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
