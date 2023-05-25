package youtrack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetProjects requests YouTrack projects.
// Source: https://www.jetbrains.com/help/youtrack/devportal/resource-api-admin-projects.html
func (c Client) GetProjects(ctx context.Context, queryParams map[string]string) ([]Project, error) {
	var response []Project

	respBody, err := c.sendReq(ctx, http.MethodGet, c.url+projectsURL, http.NoBody, queryParams)
	if err != nil {
		return nil, fmt.Errorf("c.sendReq: %w", err)
	}
	defer respBody.Close()

	err = json.NewDecoder(respBody).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return response, nil
}

// GetIssuesCount requests YouTrack issues count.
// Source: https://www.jetbrains.com/help/youtrack/devportal/resource-api-issuesGetter-count.html
func (c Client) GetIssuesCount(ctx context.Context, req IssuesCountReq, queryParams map[string]string) (IssuesCountResponse, error) {
	var response IssuesCountResponse

	reqJSON, err := json.Marshal(req)
	if err != nil {
		return IssuesCountResponse{}, fmt.Errorf("json.Marshal: %w", err)
	}

	respBody, err := c.sendReq(ctx, http.MethodPost, c.url+issuesCountURL, bytes.NewBuffer(reqJSON), queryParams)
	if err != nil {
		return IssuesCountResponse{}, fmt.Errorf("c.sendReq: %w", err)
	}
	defer respBody.Close()

	err = json.NewDecoder(respBody).Decode(&response)
	if err != nil {
		return IssuesCountResponse{}, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return response, nil
}

// GetIssues requests YouTrack issues list.
// Source: https://www.jetbrains.com/help/youtrack/devportal/resource-api-issues.html
func (c Client) GetIssues(ctx context.Context, queryParams map[string]string) ([]Issue, error) {
	var response []Issue

	respBody, err := c.sendReq(ctx, http.MethodGet, c.url+issuesURL, http.NoBody, queryParams)
	if err != nil {
		return nil, fmt.Errorf("c.sendReq: %w", err)
	}
	defer respBody.Close()

	err = json.NewDecoder(respBody).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return response, nil
}

// GetIssueActivityItems requests YouTrack issue activity items.
// Source: https://www.jetbrains.com/help/youtrack/devportal/resource-api-issues-issueID-activities.html
func (c Client) GetIssueActivityItems(ctx context.Context, issueID string, queryParams map[string]string) ([]ActivityItem, error) {
	var (
		response []ActivityItem
		url      = c.url + fmt.Sprintf(issueActivityItemsURL, issueID)
	)

	respBody, err := c.sendReq(ctx, http.MethodGet, url, http.NoBody, queryParams)
	if err != nil {
		return nil, fmt.Errorf("c.sendReq: %w", err)
	}
	defer respBody.Close()

	err = json.NewDecoder(respBody).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return response, nil
}

// GetUsers requests YouTrack users list.
// Source: https://www.jetbrains.com/help/youtrack/devportal/resource-api-users.html
func (c Client) GetUsers(ctx context.Context, queryParams map[string]string) ([]User, error) {
	var response []User

	respBody, err := c.sendReq(ctx, http.MethodGet, c.url+usersURL, http.NoBody, queryParams)
	if err != nil {
		return nil, fmt.Errorf("c.sendReq: %w", err)
	}
	defer respBody.Close()

	err = json.NewDecoder(respBody).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return response, nil
}

// GetWorkItems requests YouTrack work items list.
// Source: https://www.jetbrains.com/help/youtrack/devportal/resource-api-workItems.html
func (c Client) GetWorkItems(ctx context.Context, queryParams map[string]string) ([]IssueWorkItem, error) {
	var response []IssueWorkItem

	respBody, err := c.sendReq(ctx, http.MethodGet, c.url+workItemsURL, http.NoBody, queryParams)
	if err != nil {
		return nil, fmt.Errorf("c.sendReq: %w", err)
	}
	defer respBody.Close()

	err = json.NewDecoder(respBody).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return response, nil
}

// Requests to YouTrack Hub Api

// GetAllUsersHub requests YouTrack Hub all users list.
// Source: https://www.jetbrains.com/help/youtrack/devportal/HUB-REST-API_Users_Get-All-Users.html
func (c Client) GetAllUsersHub(ctx context.Context, queryParams map[string]string) (HubUsers, error) {
	var response HubUsers

	respBody, err := c.sendReq(ctx, http.MethodGet, c.url+hubUsersURL, http.NoBody, queryParams)
	if err != nil {
		return HubUsers{}, fmt.Errorf("c.sendReq: %w", err)
	}
	defer respBody.Close()

	err = json.NewDecoder(respBody).Decode(&response)
	if err != nil {
		return HubUsers{}, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return response, nil
}

// GetAllProjectsHub requests YouTrack Hub all projects list.
// Source: https://www.jetbrains.com/help/youtrack/devportal/HUB-REST-API_Projects_Get-All-Projects.html
func (c Client) GetAllProjectsHub(ctx context.Context, queryParams map[string]string) (HubProjects, error) {
	var response HubProjects

	respBody, err := c.sendReq(ctx, http.MethodGet, c.url+hubProjectsURL, http.NoBody, queryParams)
	if err != nil {
		return HubProjects{}, fmt.Errorf("c.sendReq: %w", err)
	}
	defer respBody.Close()

	err = json.NewDecoder(respBody).Decode(&response)
	if err != nil {
		return HubProjects{}, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return response, nil
}

func (c Client) Ping(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, c.url, http.NoBody)
	if err != nil {
		return fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusInternalServerError {
		return fmt.Errorf("checking status code: status code - %d", resp.StatusCode)
	}

	return nil
}
