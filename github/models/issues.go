package models

import (
  "time"
)

// represents our issue object from github api

type GithubIssues []struct {
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL     string `json:"labels_url"`
	CommentsURL   string `json:"comments_url"`
	EventsURL     string `json:"events_url"`
	ID            int    `json:"id"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	User          struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Labels   []interface{} `json:"labels"`
	State    string        `json:"state"`
	Locked   bool          `json:"locked"`
	Assignee struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
	} `json:"assignee"`
	Assignees []struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
	} `json:"assignees"`
	Milestone         interface{} `json:"milestone"`
	Comments          int         `json:"comments"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	ClosedAt          interface{} `json:"closed_at"`
	AuthorAssociation string      `json:"author_association"`
	ActiveLockReason  interface{} `json:"active_lock_reason"`
	Repository        struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		URL              string      `json:"url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		EventsURL        string      `json:"events_url"`
		AssigneesURL     string      `json:"assignees_url"`
		StatusesURL      string      `json:"statuses_url"`
		LanguagesURL     string      `json:"languages_url"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		IssuesURL        string      `json:"issues_url"`
		MilestonesURL    string      `json:"milestones_url"`
		Language         string      `json:"language"`
		HasIssues        bool        `json:"has_issues"`
		HasProjects      bool        `json:"has_projects"`
		MirrorURL        interface{} `json:"mirror_url"`
		Archived         bool        `json:"archived"`
		Disabled         bool        `json:"disabled"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		OpenIssues       int         `json:"open_issues"`
		DefaultBranch    string      `json:"default_branch"`
		Permissions      struct {
			Admin bool `json:"admin"`
			Push  bool `json:"push"`
			Pull  bool `json:"pull"`
		} `json:"permissions"`
	} `json:"repository"`
	Body                  interface{} `json:"body"`
	PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
}
