package youtrack

// entities.go contains YouTrack entities. You can add other fields to entities if needed.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entities.html

// Project entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-Project.html
type Project struct {
	Name string `json:"name,omitempty"`
}

// IssuesCountResponse entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-IssueCountResponse.html
type IssuesCountResponse struct {
	Count int `json:"count,omitempty"`
}

// Issue entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-Issue.html
type Issue struct {
	ID           string             `json:"id,omitempty"`
	IDReadable   string             `json:"idReadable,omitempty"`
	Summary      string             `json:"summary,omitempty"`
	Project      Project            `json:"project,omitempty"`
	CustomFields []IssueCustomField `json:"customFields,omitempty"`
}

// IssueCustomField entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-IssueCustomField.html
type IssueCustomField struct {
	Name  string `json:"name,omitempty"`
	Value any    `json:"value,omitempty"`
}

// PeriodValue entity.
// https://www.jetbrains.com/help/youtrack/devportal/api-entity-PeriodValue.html
type PeriodValue struct {
	ID           string `json:"id,omitempty"`
	Minutes      int    `json:"minutes,omitempty"`
	Presentation string `json:"presentation,omitempty"`
}

// ActivityItem entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-ActivityItem.html
type ActivityItem struct {
	Added     any   `json:"added,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
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
	Author   User          `json:"author,omitempty"`
	Duration DurationValue `json:"duration,omitempty"`
}

// DurationValue entity.
// Source: https://www.jetbrains.com/help/youtrack/devportal/api-entity-DurationValue.html
type DurationValue struct {
	Minutes int `json:"minutes,omitempty"`
}

// Hub entities
// Source: https://www.jetbrains.com/help/youtrack/devportal/HUB-REST-API_JSON-Scheme.html

type HubUsers struct {
	Type  string    `json:"type,omitempty"`
	Next  string    `json:"next,omitempty"`
	Skip  int       `json:"skip,omitempty"`
	Top   int       `json:"top,omitempty"`
	Users []HubUser `json:"users,omitempty"`
}

// HubUser is a "user" entity from hub api.
// Source: https://www.jetbrains.com/help/youtrack/devportal/HUB-REST-API_JSON-Scheme.html#user
type HubUser struct {
	Login string `json:"login,omitempty"`
}

type HubProjects struct {
	Skip     int          `json:"skip,omitempty"`
	Top      int          `json:"top,omitempty"`
	Total    int          `json:"total,omitempty"`
	Projects []HubProject `json:"projects,omitempty"`
}

// HubProject is a "project" entity from hub api.
// Source: https://www.jetbrains.com/help/youtrack/devportal/HUB-REST-API_JSON-Scheme.html#project
type HubProject struct {
	Name string `json:"name,omitempty"`
}
