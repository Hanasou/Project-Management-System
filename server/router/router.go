package router

import (
	"github.com/gorilla/mux"
	"github.com/projmanserver/middleware"
	"github.com/projmanserver/middleware/auth"
	"github.com/projmanserver/middleware/comments"
	"github.com/projmanserver/middleware/issues"
	"github.com/projmanserver/middleware/projects"
	"github.com/projmanserver/middleware/teams"
)

// NewRouter creates a router
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Test apis
	router.HandleFunc("/test", middleware.TestFunc)
	router.HandleFunc("/test/readReq", middleware.ReadRequest)
	router.HandleFunc("/test/testpost", middleware.TestPost).Methods("POST", "OPTIONS")
	router.Handle("/test/authReq", auth.JwtVerify(middleware.AuthRequest))
	router.HandleFunc("/test/path/{id}", middleware.TestPathParams).Methods("GET", "OPTIONS")

	// Auth apis
	router.HandleFunc("/auth/signup", auth.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/login", auth.Login).Methods("POST", "OPTIONS")

	// Project apis
	router.HandleFunc("/projects/add/{email}", projects.AddProject).Methods("POST", "OPTIONS")
	router.HandleFunc("/projects/getAll/{email}", projects.GetProjects).Methods("GET", "OPTIONS")
	router.HandleFunc("/projects/get/{projectID}/{userEmail}", projects.GetProject).Methods("GET", "OPTIONS")
	router.HandleFunc("/projects/delete/{projectID}/{email}", projects.DeleteProject).Methods("DELETE", "OPTIONS")

	// Issue apis
	router.HandleFunc("/issues/add", issues.CreateIssue).Methods("POST", "OPTIONS")
	router.HandleFunc("/issues/update", issues.UpdateIssue).Methods("POST", "OPTIONS")
	router.HandleFunc("/issues/getAll/{projectID}", issues.GetIssues).Methods("GET", "OPTIONS")
	router.HandleFunc("/issues/getByUser/{userEmail}", issues.GetIssuesByUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/issues/get/{projectID}/{issueID}", issues.GetIssue).Methods("GET", "OPTIONS")
	router.HandleFunc("/issues/delete/{issueID}/{projectID}", issues.DeleteIssue).Methods("DELETE", "OPTIONS")

	// Team apis
	router.HandleFunc("/teams/add", teams.CreateTeam).Methods("POST", "OPTIONS")
	router.HandleFunc("/teams/getAll/{memberEmail}", teams.GetTeams).Methods("GET", "OPTIONS")
	router.HandleFunc("/teams/get/{projectID}", teams.GetTeam).Methods("GET", "OPTIONS")

	// Comment apis
	router.HandleFunc("/comments/add", comments.AddComment).Methods("POST", "OPTIONS")
	router.HandleFunc("/comments/getAll/{issueID}", comments.GetComments).Methods("GET", "OPTIONS")
	return router
}
