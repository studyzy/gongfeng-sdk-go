package gongfeng

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// MergeRequest 表示工蜂合并请求。
type MergeRequest struct {
	ID              int        `json:"id,omitempty"`
	IID             int        `json:"iid,omitempty"`
	Title           string     `json:"title,omitempty"`
	Description     string     `json:"description,omitempty"`
	State           string     `json:"state,omitempty"`
	TargetBranch    string     `json:"target_branch,omitempty"`
	SourceBranch    string     `json:"source_branch,omitempty"`
	TargetProjectID int        `json:"target_project_id,omitempty"`
	SourceProjectID int        `json:"source_project_id,omitempty"`
	Assignee        *User      `json:"assignee,omitempty"`
	Author          *User      `json:"author,omitempty"`
	Milestone       *Milestone `json:"milestone,omitempty"`
	ProjectID       int        `json:"project_id,omitempty"`
	WorkInProgress  bool       `json:"work_in_progress,omitempty"`
	Labels          []string   `json:"labels,omitempty"`
	CreatedAt       Time       `json:"created_at,omitempty"`
	UpdatedAt       Time       `json:"updated_at,omitempty"`
	Upvotes         int        `json:"upvotes,omitempty"`
	Downvotes       int        `json:"downvotes,omitempty"`
}

// MergeRequestChanges 表示合并请求的详情及其代码变更。
type MergeRequestChanges struct {
	MergeRequest
	Files []*Diff `json:"files,omitempty"`
}

// MRComment 表示合并请求的评论。
type MRComment struct {
	ID         int    `json:"id,omitempty"`
	Body       string `json:"body,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	Author     *User  `json:"author,omitempty"`
	CreatedAt  Time   `json:"created_at,omitempty"`
	System     bool   `json:"system,omitempty"`
}

// MRSubscription 表示合并请求的订阅状态。
type MRSubscription struct {
	Subscribed bool `json:"subscribed,omitempty"`
}

// MergeRequestsService 处理与工蜂合并请求相关的 API。
type MergeRequestsService struct {
	client *Client
}

// CreateMergeRequestOptions 表示 CreateMergeRequest 的可选参数。
type CreateMergeRequestOptions struct {
	SourceBranch *string `json:"source_branch,omitempty"`
	TargetBranch *string `json:"target_branch,omitempty"`
	Title        *string `json:"title,omitempty"`
	AssigneeID   *int    `json:"assignee_id,omitempty"`
	Description  *string `json:"description,omitempty"`
	Reviewers    *string `json:"reviewers,omitempty"`
	ApproverRule *string `json:"approver_rule,omitempty"`
}

// CreateMergeRequest 新增合并请求。
func (s *MergeRequestsService) CreateMergeRequest(ctx context.Context, pid interface{}, opts *CreateMergeRequestOptions) (*MergeRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var mr MergeRequest
	resp, err := s.client.Do(req, &mr)
	if err != nil {
		return nil, resp, err
	}

	return &mr, resp, nil
}

// AcceptMergeRequestOptions 表示 AcceptMergeRequest 的可选参数。
type AcceptMergeRequestOptions struct {
	MergeCommitMessage *string `json:"merge_commit_message,omitempty"`
}

// AcceptMergeRequest 合并一个合并请求。
func (s *MergeRequestsService) AcceptMergeRequest(ctx context.Context, pid interface{}, mergeRequestID int, opts *AcceptMergeRequestOptions) (*MergeRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_request/%d/merge", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var mr MergeRequest
	resp, err := s.client.Do(req, &mr)
	if err != nil {
		return nil, resp, err
	}

	return &mr, resp, nil
}

// ListMergeRequestCommitsOptions 表示 ListMergeRequestCommits 的可选参数。
type ListMergeRequestCommitsOptions struct {
	ListOptions
}

// ListMergeRequestCommits 获取合并请求中的提交列表。
func (s *MergeRequestsService) ListMergeRequestCommits(ctx context.Context, pid interface{}, mergeRequestID int, opts *ListMergeRequestCommitsOptions) ([]*Commit, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/commits", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var commits []*Commit
	resp, err := s.client.Do(req, &commits)
	if err != nil {
		return nil, resp, err
	}

	return commits, resp, nil
}

// CreateMRCommentOptions 表示 CreateMRComment 的可选参数。
type CreateMRCommentOptions struct {
	Body *string `json:"body,omitempty"`
}

// CreateMRComment 添加合并请求评论。
func (s *MergeRequestsService) CreateMRComment(ctx context.Context, pid interface{}, mergeRequestID int, opts *CreateMRCommentOptions) (*MRComment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_request/%d/comments", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var c MRComment
	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return &c, resp, nil
}

// ListMRCommentsOptions 表示 ListMRComments 的可选参数。
type ListMRCommentsOptions struct {
	ListOptions
}

// ListMRComments 获取合并请求的评论列表。
func (s *MergeRequestsService) ListMRComments(ctx context.Context, pid interface{}, mergeRequestID int, opts *ListMRCommentsOptions) ([]*MRComment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_request/%d/comments", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var comments []*MRComment
	resp, err := s.client.Do(req, &comments)
	if err != nil {
		return nil, resp, err
	}

	return comments, resp, nil
}

// UpdateMergeRequestOptions 表示 UpdateMergeRequest 的可选参数。
type UpdateMergeRequestOptions struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	TargetBranch *string `json:"target_branch,omitempty"`
	AssigneeID   *int    `json:"assignee_id,omitempty"`
	StateEvent   *string `json:"state_event,omitempty"`
}

// UpdateMergeRequest 更新合并请求。
func (s *MergeRequestsService) UpdateMergeRequest(ctx context.Context, pid interface{}, mergeRequestID int, opts *UpdateMergeRequestOptions) (*MergeRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_request/%d", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var mr MergeRequest
	resp, err := s.client.Do(req, &mr)
	if err != nil {
		return nil, resp, err
	}

	return &mr, resp, nil
}

// ListMergeRequestsOptions 表示 ListMergeRequests 的可选参数。
type ListMergeRequestsOptions struct {
	ListOptions
	State   *string `url:"state,omitempty" json:"state,omitempty"`
	OrderBy *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort    *string `url:"sort,omitempty" json:"sort,omitempty"`
	IID     *int    `url:"iid,omitempty" json:"iid,omitempty"`
}

// ListMergeRequests 获取项目的合并请求列表。
func (s *MergeRequestsService) ListMergeRequests(ctx context.Context, pid interface{}, opts *ListMergeRequestsOptions) ([]*MergeRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var mrs []*MergeRequest
	resp, err := s.client.Do(req, &mrs)
	if err != nil {
		return nil, resp, err
	}

	return mrs, resp, nil
}

// GetMergeRequestChanges 查询合并请求的代码变更。
func (s *MergeRequestsService) GetMergeRequestChanges(ctx context.Context, pid interface{}, mergeRequestID int) (*MergeRequestChanges, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_request/%d/changes", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var changes MergeRequestChanges
	resp, err := s.client.Do(req, &changes)
	if err != nil {
		return nil, resp, err
	}

	return &changes, resp, nil
}

// GetMergeRequest 查询单个合并请求。
func (s *MergeRequestsService) GetMergeRequest(ctx context.Context, pid interface{}, mergeRequestID int) (*MergeRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_request/%d", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var mr MergeRequest
	resp, err := s.client.Do(req, &mr)
	if err != nil {
		return nil, resp, err
	}

	return &mr, resp, nil
}

// GetMRSubscription 查询是否订阅了合并请求。
func (s *MergeRequestsService) GetMRSubscription(ctx context.Context, pid interface{}, mergeRequestID int) (*MRSubscription, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_request/%d/subscribe", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var sub MRSubscription
	resp, err := s.client.Do(req, &sub)
	if err != nil {
		return nil, resp, err
	}

	return &sub, resp, nil
}

// SubscribeMR 订阅合并请求。
func (s *MergeRequestsService) SubscribeMR(ctx context.Context, pid interface{}, mergeRequestID int) (*MRSubscription, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_request/%d/subscribe", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var sub MRSubscription
	resp, err := s.client.Do(req, &sub)
	if err != nil {
		return nil, resp, err
	}

	return &sub, resp, nil
}

// UnsubscribeMR 取消订阅合并请求。
func (s *MergeRequestsService) UnsubscribeMR(ctx context.Context, pid interface{}, mergeRequestID int) (*MRSubscription, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_request/%d/unsubscribe", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var sub MRSubscription
	resp, err := s.client.Do(req, &sub)
	if err != nil {
		return nil, resp, err
	}

	return &sub, resp, nil
}

// DownloadMergeRequestChangedFiles 下载合并请求的差异文件集。
func (s *MergeRequestsService) DownloadMergeRequestChangedFiles(ctx context.Context, pid interface{}, mergeRequestID int, w io.Writer) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/changed_files", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, w)
}
