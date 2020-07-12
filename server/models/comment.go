package models

// Comment represents a comment someone can post on an issue
type Comment struct {
	IssueID   string
	CommentID string
	Email     string
	Content   string
	Timestamp string
}
