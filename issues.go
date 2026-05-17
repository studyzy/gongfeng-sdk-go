package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Issue 表示一个缺陷。
type Issue struct {
	ID           int        `json:"id,omitempty"`
	IID          int        `json:"iid,omitempty"`
	ProjectID    int        `json:"project_id,omitempty"`
	Title        string     `json:"title,omitempty"`
	Description  string     `json:"description,omitempty"`
	State        string     `json:"state,omitempty"`
	ResolveState string     `json:"resolve_state,omitempty"`
	Grade        *int       `json:"grade,omitempty"`
	Labels       []string   `json:"labels,omitempty"`
	Assignees    []*User    `json:"assignees,omitempty"`
	Assignee     *User      `json:"assignee,omitempty"`
	Author       *User      `json:"author,omitempty"`
	CreatedAt    Time       `json:"created_at,omitempty"`
	UpdatedAt    Time       `json:"updated_at,omitempty"`
	Milestone    *Milestone `json:"milestone,omitempty"`
}

// IssuesService 处理与 Issue 相关的 API 调用。
type IssuesService struct {
	client *Client
}

// CreateIssueOptions 是 CreateIssue 的可选参数。
type CreateIssueOptions struct {
	Title       *string `json:"title,omitempty" url:"title,omitempty"`
	Grade       *int    `json:"grade,omitempty" url:"grade,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	AssigneeIDs *string `json:"assignee_ids,omitempty" url:"assignee_ids,omitempty"`
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
	Title        *string `json:"title,omitempty" url:"title,omitempty"`
	ResolveState *string `json:"resolve_state,omitempty" url:"resolve_state,omitempty"`
	Grade        *int    `json:"grade,omitempty" url:"grade,omitempty"`
	Description  *string `json:"description,omitempty" url:"description,omitempty"`
	AssigneeIDs  *string `json:"assignee_ids,omitempty" url:"assignee_ids,omitempty"`
	MilestoneID  *int    `json:"milestone_id,omitempty" url:"milestone_id,omitempty"`
	Labels       *string `json:"labels,omitempty" url:"labels,omitempty"`
	StateEvent   *string `json:"state_event,omitempty" url:"state_event,omitempty"`
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
	IID           *int    `url:"iid,omitempty" json:"iid,omitempty"`
	ResolveState  *string `url:"resolve_state,omitempty" json:"resolve_state,omitempty"`
	Grade         *int    `url:"grade,omitempty" json:"grade,omitempty"`
	State         *string `url:"state,omitempty" json:"state,omitempty"`
	Labels        *string `url:"labels,omitempty" json:"labels,omitempty"`
	Milestone     *string `url:"milestone,omitempty" json:"milestone,omitempty"`
	OrderBy       *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort          *string `url:"sort,omitempty" json:"sort,omitempty"`
	CreatedAfter  *string `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore *string `url:"created_before,omitempty" json:"created_before,omitempty"`
}

// ListUserIssues 获取当前用户创建的缺陷列表。
func (s *IssuesService) ListUserIssues(ctx context.Context, opts *ListIssuesOptions) ([]*Issue, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "issues", opts)
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

// GetIssueSubscription 查询缺陷订阅状态。
func (s *IssuesService) GetIssueSubscription(ctx context.Context, pid interface{}, issueID int) (bool, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return false, nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d/subscribe", project, issueID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return false, nil, err
	}

	var subscribed bool
	resp, err := s.client.Do(req, &subscribed)
	if err != nil {
		return false, resp, err
	}

	return subscribed, resp, nil
}

// SubscribeIssue 订阅缺陷。
func (s *IssuesService) SubscribeIssue(ctx context.Context, pid interface{}, issueID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d/subscribe", project, issueID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnsubscribeIssue 取消订阅缺陷。
func (s *IssuesService) UnsubscribeIssue(ctx context.Context, pid interface{}, issueID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d/unsubscribe", project, issueID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
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
