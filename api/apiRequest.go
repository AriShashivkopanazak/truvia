package api

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)


type Response struct {
    ResponseCode    int    `json:"response_code"`
    Questions []Question `json:"results"`
}


func apiRequest(difficulty string, typeOf string, questions uint) {
	var responseObject Response

	requestUrl := "https://opentdb.com/api.php" + difficulty + typeOf + questions

	response, err := http.Get(requestUrl)
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
