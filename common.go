package youtrack

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// Youtrack API URLs
const (
	apiURL                = `/api`
	issuesURL             = apiURL + `/issues`
	issuesCountURL        = apiURL + `/issuesGetter/count`
	projectsURL           = apiURL + `/admin/projects`
	issueActivityItemsURL = issuesURL + `/%s/activities`
	usersURL              = apiURL + `/users`
	workItemsURL          = apiURL + `/workItems`
)

// Youtrack Hub Rest API URLs
const (
	hubAPIURL      = `/hub/api/rest`
	hubUsersURL    = hubAPIURL + `/users`
	hubProjectsURL = hubAPIURL + `/projects`
)

// Issue activity categories
const (
	CustomFieldActivityCategory = `CustomFieldCategory`
)

const (
	// If the issues count response returns -1, it means that YouTrack hasn't finished counting the issues yet.
	IssuesCountUnknownValue = -1
)

// addDefaultHeaders adds headers which are required for YouTrack API.
func (c Client) addDefaultHeaders(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/json")
}

func (c Client) addQueryParams(req *http.Request, params map[string]string) {
	q := req.URL.Query()

	for param, value := range params {
		q.Add(param, value)
	}
	req.URL.RawQuery = q.Encode()
}

// checkResponseStatusCode returns error if status != OK.
// Reads response body and returns error with body content, so you don't need to read and close body manually in your methods.
func (c Client) checkResponseStatusCode(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()

		errResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("bad response status code - %d: io.ReadAll: %w", resp.StatusCode, err)
		}
		return fmt.Errorf("bad response status code - %d. Response: %s", resp.StatusCode, errResp)
	}

	return nil
}

func (c Client) sendReq(ctx context.Context, method, url string, body io.Reader, queryParams map[string]string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	c.addDefaultHeaders(req)
	c.addQueryParams(req, queryParams)

	httpResp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do: %w", err)
	}

	if err = c.checkResponseStatusCode(httpResp); err != nil {
		return nil, fmt.Errorf("c.checkResponseStatusCode: %w", err)
	}

	return httpResp.Body, nil
}
