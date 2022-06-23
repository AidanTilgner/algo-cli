package data

import (
	t "algo-cli/types"
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetQuestions(path string) []t.Question {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	questions := make([]t.Question, 0)
	err = json.Unmarshal(content, &questions)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	log.Println("Questions: ", questions)

	return questions
}
