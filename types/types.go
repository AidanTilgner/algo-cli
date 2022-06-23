package types

type Question struct {
	Id          int
	Name        string
	Unique_name string
	Description string
	Difficulty  int
}

type Result struct {
	Rank     int
	Question Question
}
