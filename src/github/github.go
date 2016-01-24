package github


import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
	"strings"
	// "fmt"
)

const IssuesURL = "https://api.github.com/search/issues"

//field tags only necessacry to have underscore 
//in the json name but not the go name fields are case insensitve 
//when matching
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items	[]*Issue	
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string 
	State string 
	User *User
	Milestone *Milestone
	CreatedAt time.Time `json:"created_at"` 
	Body string //in markdown format
	MILESTONEURL string `json:"milestones_url`
	CONTRIBUTORSURL string `json: "contributors_url"`
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}
type Milestone struct {
	HTMLURL string `json:"html_url"`
	State string
	Title string
	Description string
	CreatedAt time.Time `json:"created_at"`
	Number int
}
//search issues queries the github issue tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " ")) // terms could have special charcaters like ? &
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// fmt.Println(IssuesURL + "?q=" + q);
	//we must close resp.Body on all execution paths.
	//(chapter 5 shows defer which make this simpler)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}