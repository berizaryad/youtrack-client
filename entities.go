package youtrack

// entities.go contains YouTrack entities. You can add other fields to entities if needed.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entities.html

// Project entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-Project.html
type Project struct {
	Name string `json:"name"`
}

// IssuesCountResponse entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-IssueCountResponse.html
type IssuesCountResponse struct {
	Count int `json:"count"`
}

// Issue entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-Issue.html
type Issue struct {
	ID         string  `json:"id"`
	IDReadable string  `json:"idReadable"`
	Summary    string  `json:"summary"`
	Project    Project `json:"project"`
}

// ActivityItem entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-ActivityItem.html
type ActivityItem struct {
	Added     any   `json:"added"`
	Timestamp int64 `json:"timestamp"`
}

// User entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-User.html
type User struct {
	RingID string `json:"ringId,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

// IssueWorkItem entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-IssueWorkItem.html
type IssueWorkItem struct {
	Author   User          `json:"author"`
	Duration DurationValue `json:"duration"`
}

// DurationValue entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-DurationValue.html
type DurationValue struct {
	Minutes int `json:"minutes"`
}

// Hub entities
// Source: https://www.jetbrains.com/help/youtrack/devportal/HUB-REST-API_JSON-Scheme.html

type HubUsers struct {
	Type  string    `json:"type"`
	Next  string    `json:"next"`
	Skip  int       `json:"skip"`
	Top   int       `json:"top"`
	Users []HubUser `json:"users"`
}

// HubUser is a "user" entity from hub api.
// Source: https://www.jetbrains.com/help/youtrack/devportal/HUB-REST-API_JSON-Scheme.html#user
type HubUser struct {
	Login string `json:"login"`
}
