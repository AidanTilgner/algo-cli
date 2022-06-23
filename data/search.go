package data

import (
	h "algo-cli/helpers"
	t "algo-cli/types"
	"fmt"
	"path/filepath"
	"runtime"
)

func getRank(q t.Question, query string) int {
	nameRank := int(h.CompareTwoStrings(q.Name, query) * 100)
	uniqueNameRank := int(h.CompareTwoStrings(q.Unique_name, query) * 100)
	return (nameRank + uniqueNameRank) / 2
}

func Search(query string) []t.Question {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	questions := GetQuestions(filepath.Join(basepath, "/questions.json"))
	results := make([]t.Result, 0)

	for _, q := range(questions) {
		results = append(results, t.Result{
			Rank: getRank(q, query),
			Question: q,
		})
	}

	fmt.Println("Results: ", results)

	sortedResults:= h.ReverseResults(h.QuickSortResults(results, 0, len(results) -1))

	fmt.Println("Sorted Results: ", sortedResults)
	sortedQuestions := make([]t.Question, 0)
	for _, r := range(sortedResults) {
		sortedQuestions = append(sortedQuestions, r.Question)
	}

	return sortedQuestions
}
