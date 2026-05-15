package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Issue 表示一个缺陷。
type Issue struct {
	ID          int        `json:"id"`
	IID         int        `json:"iid"`
	ProjectID   int        `json:"project_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	State       string     `json:"state"`
	Labels      []string   `json:"labels"`
	Assignee    *User      `json:"assignee"`
	Author      *User      `json:"author"`
	CreatedAt   Time       `json:"created_at"`
	UpdatedAt   Time       `json:"updated_at"`
	Milestone   *Milestone `json:"milestone"`
}

// IssuesService 处理与 Issue 相关的 API 调用。
type IssuesService struct {
	client *Client
}

// CreateIssueOptions 是 CreateIssue 的可选参数。
type CreateIssueOptions struct {
	Title       *string `json:"title,omitempty" url:"title,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	AssigneeID  *int    `json:"assignee_id,omitempty" url:"assignee_id,omitempty"`
	MilestoneID *int    `json:"milestone_id,omitempty" url:"milestone_id,omitempty"`
	Labels      *string `json:"labels,omitempty" url:"labels,omitempty"`
}

// CreateIssue 在项目中新建一个缺陷。
func (s *IssuesService) CreateIssue(ctx context.Context, pid interface{}, opts *CreateIssueOptions) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/issues", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var issue Issue
	resp, err := s.client.Do(req, &issue)
	if err != nil {
		return nil, resp, err
	}

	return &issue, resp, nil
}

// UpdateIssueOptions 是 UpdateIssue 的可选参数。
type UpdateIssueOptions struct {
	Title       *string `json:"title,omitempty" url:"title,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	AssigneeID  *int    `json:"assignee_id,omitempty" url:"assignee_id,omitempty"`
	MilestoneID *int    `json:"milestone_id,omitempty" url:"milestone_id,omitempty"`
	Labels      *string `json:"labels,omitempty" url:"labels,omitempty"`
	StateEvent  *string `json:"state_event,omitempty" url:"state_event,omitempty"`
}

// UpdateIssue 编辑项目中指定的缺陷。
func (s *IssuesService) UpdateIssue(ctx context.Context, pid interface{}, issueID int, opts *UpdateIssueOptions) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d", project, issueID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var issue Issue
	resp, err := s.client.Do(req, &issue)
	if err != nil {
		return nil, resp, err
	}

	return &issue, resp, nil
}

// ListIssuesOptions 是 ListIssues 的可选参数。
type ListIssuesOptions struct {
	ListOptions
	State     *string `url:"state,omitempty" json:"state,omitempty"`
	Labels    *string `url:"labels,omitempty" json:"labels,omitempty"`
	Milestone *string `url:"milestone,omitempty" json:"milestone,omitempty"`
	OrderBy   *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort      *string `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListIssues 获取项目的缺陷列表。
func (s *IssuesService) ListIssues(ctx context.Context, pid interface{}, opts *ListIssuesOptions) ([]*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/issues", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var issues []*Issue
	resp, err := s.client.Do(req, &issues)
	if err != nil {
		return nil, resp, err
	}

	return issues, resp, nil
}

// GetIssue 获取项目中指定 ID 的缺陷。
func (s *IssuesService) GetIssue(ctx context.Context, pid interface{}, issueID int) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d", project, issueID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var issue Issue
	resp, err := s.client.Do(req, &issue)
	if err != nil {
		return nil, resp, err
	}

	return &issue, resp, nil
}

// DeleteIssue 删除项目中指定 ID 的缺陷。
func (s *IssuesService) DeleteIssue(ctx context.Context, pid interface{}, issueID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d", project, issueID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
