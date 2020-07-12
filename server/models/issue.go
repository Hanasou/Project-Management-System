package models

// Issue model
type Issue struct {
	IssueID     string
	ProjectID   string
	Title       string
	Description string
	Type        string
	Priority    string
	Timestamp   string
	Status      string
	Creator     string
}
