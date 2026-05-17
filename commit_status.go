package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// CommitStatus 表示一个提交的检测状态。
type CommitStatus struct {
	ID          int    `json:"id,omitempty"`
	SHA         string `json:"sha,omitempty"`
	Ref         string `json:"ref,omitempty"`
	Status      string `json:"status,omitempty"`
	Name        string `json:"name,omitempty"`
	TargetURL   string `json:"target_url,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   Time   `json:"created_at,omitempty"`
	Author      *User  `json:"author,omitempty"`
}

// CommitStatusResult 表示一个提交的检测组合结果。
type CommitStatusResult struct {
	SHA      string          `json:"sha,omitempty"`
	Status   string          `json:"status,omitempty"`
	Statuses []*CommitStatus `json:"statuses,omitempty"`
}

// CommitStatusService 处理与提交检测状态相关的 API 调用。
type CommitStatusService struct {
	client *Client
}

// CreateCommitStatusOptions 是 CreateCommitStatus 的可选参数。
type CreateCommitStatusOptions struct {
	State       *string `json:"state,omitempty" url:"state,omitempty"`
	Ref         *string `json:"ref,omitempty" url:"ref,omitempty"`
	Name        *string `json:"name,omitempty" url:"name,omitempty"`
	TargetURL   *string `json:"target_url,omitempty" url:"target_url,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	Context     *string `json:"context,omitempty" url:"context,omitempty"`
}

// CreateCommitStatus 为指定提交新建一个检测结果。
func (s *CommitStatusService) CreateCommitStatus(ctx context.Context, pid interface{}, sha string, opts *CreateCommitStatusOptions) (*CommitStatus, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/commit/%s/statuses", project, pathEscape(sha))

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var status CommitStatus
	resp, err := s.client.Do(req, &status)
	if err != nil {
		return nil, resp, err
	}

	return &status, resp, nil
}

// ListCommitStatusesOptions 是 ListCommitStatuses 的可选参数。
type ListCommitStatusesOptions struct {
	ListOptions
	Ref   *string `url:"ref,omitempty" json:"ref,omitempty"`
	Stage *string `url:"stage,omitempty" json:"stage,omitempty"`
	Name  *string `url:"name,omitempty" json:"name,omitempty"`
	All   *bool   `url:"all,omitempty" json:"all,omitempty"`
}

// ListCommitStatuses 查询指定提交的检测结果列表。
func (s *CommitStatusService) ListCommitStatuses(ctx context.Context, pid interface{}, sha string, opts *ListCommitStatusesOptions) ([]*CommitStatus, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/commits/%s/statuses", project, pathEscape(sha))

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var statuses []*CommitStatus
	resp, err := s.client.Do(req, &statuses)
	if err != nil {
		return nil, resp, err
	}

	return statuses, resp, nil
}

// GetCommitStatusResult 查询指定提交的检测组合结果。
func (s *CommitStatusService) GetCommitStatusResult(ctx context.Context, pid interface{}, ref string) (*CommitStatusResult, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/commits/%s/status", project, pathEscape(ref))

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var result CommitStatusResult
	resp, err := s.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
