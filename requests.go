package youtrack

// requests.go contains structs for JSON request bodies. You can add other fields if needed.

// IssuesCountReq - IssuesCount request body.
// Source: https://www.jetbrains.com/help/youtrack/devportal/resource-api-issuesGetter-count.html#create-IssueCountResponse-method-sample-1
type IssuesCountReq struct {
	Query string `json:"query,omitempty"`
}
