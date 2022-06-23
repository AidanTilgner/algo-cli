package data

import (
	h "algo-cli/helpers"
	t "algo-cli/types"
	"fmt"
)

func getRank(q t.Question, query string) t.Result {
	nameRank := h.CompareStrings(q.Name, query)
	return t.Result{
		Rank:     nameRank,
		Question: q,
	}
}

func Search(query string) {
	questions := GetQuestions("./data/questions.json")
	// Contains Duplicate
	fmt.Println("Rank: ", getRank(questions[0], query))
}
