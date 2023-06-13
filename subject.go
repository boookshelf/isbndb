package isbndb

import "context"

const (
	subjectPath      = "subject"
	subjectQueryPath = "subjects"
)

type Subject struct {
	Subject string `json:"subject,omitempty"`
	Parent  string `json:"parent,omitempty"`
}

type SubjectQuery struct {
	Total    int      `json:"total,omitempty"`
	Subjects []string `json:"subjects,omitempty"`
}

// Returns details and a list of books with subject
func (c *Client) GetSubject(ctx context.Context, name string) (Subject, error) {
	url := c.baseURL.JoinPath(subjectPath, name).String()
	var result Subject
	err := c.get(ctx, url, &result)

	return result, err
}

// This returns a list of subjects that match the given query
func (c *Client) QuerySubject(ctx context.Context, query string) (SubjectQuery, error) {
	url := c.baseURL.JoinPath(subjectQueryPath, query).String()
	var result SubjectQuery
	err := c.get(ctx, url, &result)

	return result, err
}
