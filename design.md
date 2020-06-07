# Truvia Design

## Features

* give chances
  * defaults to 3 guesses
* asks questions in form of:
  * multiple choice
  * true/false
  * ai matched answer
  * defaults to random

* has difficulty
  * easy
  * hard
  * medium
  * defaults to random

* asks number of questions
  * defaults to 1 question

## Constraints

* Outputs as string
* Requests in form of json
* Receives response in form of json

## design

* ai matched answer
  * finds matches in words
  * or looks up definition in thesauras and dictionary and matches those definitions
  * %90 accurate award correct answer
* user input
  * see man page
* send api request
* process api request
  * question
  * if wrong answer give 3 guesses
