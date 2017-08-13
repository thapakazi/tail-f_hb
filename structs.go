package main

import (
	"time"
)

// type Results struct{
// 	Message string
// 	Env string
// 	Notices_Count string
// 	Url string
// 	Deploy Deploy
// }

// type Deploy struct{
// 	Id int
// 	Environment string
// 	Revision string
// 	Repository string
// 	Deployed_By string
// 	Created_At string
// 	Changelog []string
// }


type Faults struct {
	Results []struct {
		ProjectID     int           `json:"project_id"`
		Klass         string        `json:"klass"`
		Component     interface{}   `json:"component"`
		Action        interface{}   `json:"action"`
		Environment   string        `json:"environment"`
		Resolved      bool          `json:"resolved"`
		Ignored       bool          `json:"ignored"`
		CreatedAt     time.Time     `json:"created_at"`
		CommentsCount int           `json:"comments_count"`
		Message       string        `json:"message"`
		NoticesCount  int           `json:"notices_count"`
		LastNoticeAt  time.Time     `json:"last_notice_at"`
		Tags          []interface{} `json:"tags"`
		ID            int           `json:"id"`
		Assignee      interface{}   `json:"assignee"`
		URL           string        `json:"url"`
		Deploy        struct {
			ID            int           `json:"id"`
			Environment   string        `json:"environment"`
			Revision      string        `json:"revision"`
			Repository    string        `json:"repository"`
			LocalUsername string        `json:"local_username"`
			CreatedAt     time.Time     `json:"created_at"`
			Changelog     []interface{} `json:"changelog"`
			URL           string        `json:"url"`
		} `json:"deploy"`
	} `json:"results"`
	Links struct {
		Self string `json:"self"`
		Next string `json:"next"`
	} `json:"links"`
}
