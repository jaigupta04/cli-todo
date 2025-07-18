package internal

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Task struct {
	ID        int
	Title     string
	Completed bool
	Priority  Priority
	Category  string
	DueDate   string
}
