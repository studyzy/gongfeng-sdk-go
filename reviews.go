package gongfeng

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// Review 表示一个 Commit 评审。
type Review struct {
	ID                    int         `json:"id,omitempty"`
	ProjectID             int         `json:"project_id,omitempty"`
	Title                 string      `json:"title,omitempty"`
	Description           string      `json:"description,omitempty"`
	State                 string      `json:"state,omitempty"`
	CreatedAt             Time        `json:"created_at,omitempty"`
	UpdatedAt             Time        `json:"updated_at,omitempty"`
	Author                *User       `json:"author,omitempty"`
	Reviewers             []*Reviewer `json:"reviewers,omitempty"`
	ReviewableID          int         `json:"reviewable_id,omitempty"`
	ReviewableType        string      `json:"reviewable_type,omitempty"`
	ApproverRule          int         `json:"approver_rule,omitempty"`
	NecessaryApproverRule int         `json:"necessary_approver_rule,omitempty"`
	PushResetEnabled      bool        `json:"push_reset_enabled,omitempty"`
}

// Reviewer 表示一个评审人。
type Reviewer struct {
	ID          int    `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	Name        string `json:"name,omitempty"`
	State       string `json:"state,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	Type        string `json:"type,omitempty"`
	ReviewState string `json:"review_state,omitempty"`
	CreatedAt   Time   `json:"created_at,omitempty"`
	UpdatedAt   Time   `json:"updated_at,omitempty"`
}

// ReviewsService 处理与 MR 评审和 Commit 评审相关的 API 调用。
type ReviewsService struct {
	client *Client
}

// --- MR 评审方法 ---

// InviteMRReviewerOptions 是 InviteMRReviewer 的可选参数。
type InviteMRReviewerOptions struct {
	ReviewerID          *int `json:"reviewer_id,omitempty" url:"reviewer_id,omitempty"`
	NecessaryReviewerID *int `json:"necessary_reviewer_id,omitempty" url:"necessary_reviewer_id,omitempty"`
}

// InviteMRReviewer 邀请评审人参与 MR 评审。
func (s *ReviewsService) InviteMRReviewer(ctx context.Context, pid interface{}, mergeRequestID int, opts *InviteMRReviewerOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_request/%d/review/invite", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RemoveMRReviewerOptions 是 RemoveMRReviewer 的可选参数。
type RemoveMRReviewerOptions struct {
	ReviewerID *int `json:"reviewer_id,omitempty" url:"reviewer_id,omitempty"`
}

// RemoveMRReviewer 移除 MR 评审人。
func (s *ReviewsService) RemoveMRReviewer(ctx context.Context, pid interface{}, mergeRequestID int, opts *RemoveMRReviewerOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_request/%d/review/dismissals", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// GetMRReview 获取 MR 评审信息。
func (s *ReviewsService) GetMRReview(ctx context.Context, pid interface{}, mergeRequestID int) (*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_request/%d/review", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var review Review
	resp, err := s.client.Do(req, &review)
	if err != nil {
		return nil, resp, err
	}

	return &review, resp, nil
}

// SubmitMRReviewSummaryOptions 是 SubmitMRReviewSummary 的可选参数。
type SubmitMRReviewSummaryOptions struct {
	ReviewerEvent *string `json:"reviewer_event,omitempty" url:"reviewer_event,omitempty"`
	Summary       *string `json:"summary,omitempty" url:"summary,omitempty"`
}

// SubmitMRReviewSummary 发表 MR 评审意见。
func (s *ReviewsService) SubmitMRReviewSummary(ctx context.Context, pid interface{}, mergeRequestID int, opts *SubmitMRReviewSummaryOptions) (*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_request/%d/reviewer/summary", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var review Review
	resp, err := s.client.Do(req, &review)
	if err != nil {
		return nil, resp, err
	}

	return &review, resp, nil
}

// ReopenMRReview 重置 MR 评审状态。
func (s *ReviewsService) ReopenMRReview(ctx context.Context, pid interface{}, mergeRequestID int) (*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_request/%d/review/reopen", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var review Review
	resp, err := s.client.Do(req, &review)
	if err != nil {
		return nil, resp, err
	}

	return &review, resp, nil
}

// CancelMRReview 取消 MR 评审。
func (s *ReviewsService) CancelMRReview(ctx context.Context, pid interface{}, mergeRequestID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_request/%d/review/cancel", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// --- Commit 评审方法 ---

// CreateCommitReviewOptions 是 CreateCommitReview 的请求参数。
type CreateCommitReviewOptions struct {
	Title                 *string `json:"title,omitempty" url:"title,omitempty"`
	SourceBranch          *string `json:"source_branch,omitempty" url:"source_branch,omitempty"`
	TargetBranch          *string `json:"target_branch,omitempty" url:"target_branch,omitempty"`
	Description           *string `json:"description,omitempty" url:"description,omitempty"`
	SourceCommit          *string `json:"source_commit,omitempty" url:"source_commit,omitempty"`
	TargetCommit          *string `json:"target_commit,omitempty" url:"target_commit,omitempty"`
	TargetProjectID       *int    `json:"target_project_id,omitempty" url:"target_project_id,omitempty"`
	ReviewerIDs           *string `json:"reviewer_ids,omitempty" url:"reviewer_ids,omitempty"`
	NecessaryReviewerIDs  *string `json:"necessary_reviewer_ids,omitempty" url:"necessary_reviewer_ids,omitempty"`
	ApproverRule          *int    `json:"approver_rule,omitempty" url:"approver_rule,omitempty"`
	NecessaryApproverRule *int    `json:"necessary_approver_rule,omitempty" url:"necessary_approver_rule,omitempty"`
}

// CreateCommitReview 新建一个 Commit 评审。
func (s *ReviewsService) CreateCommitReview(ctx context.Context, pid interface{}, opts *CreateCommitReviewOptions) (*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/review", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var review Review
	resp, err := s.client.Do(req, &review)
	if err != nil {
		return nil, resp, err
	}

	return &review, resp, nil
}

// ListCommitReviewsOptions 是 ListCommitReviews 的可选参数。
type ListCommitReviewsOptions struct {
	ListOptions
	AuthorID *int    `url:"author_id,omitempty" json:"author_id,omitempty"`
	State    *string `url:"state,omitempty" json:"state,omitempty"`
	OrderBy  *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort     *string `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListCommitReviews 获取项目的 Commit 评审列表。
func (s *ReviewsService) ListCommitReviews(ctx context.Context, pid interface{}, opts *ListCommitReviewsOptions) ([]*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/reviews", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var reviews []*Review
	resp, err := s.client.Do(req, &reviews)
	if err != nil {
		return nil, resp, err
	}

	return reviews, resp, nil
}

// GetCommitReview 获取单个 Commit 评审。
func (s *ReviewsService) GetCommitReview(ctx context.Context, pid interface{}, reviewID int) (*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/review/%d", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var review Review
	resp, err := s.client.Do(req, &review)
	if err != nil {
		return nil, resp, err
	}

	return &review, resp, nil
}

// InviteCommitReviewerOptions 是 InviteCommitReviewer 的可选参数。
type InviteCommitReviewerOptions struct {
	ReviewerID          *int `json:"reviewer_id,omitempty" url:"reviewer_id,omitempty"`
	NecessaryReviewerID *int `json:"necessary_reviewer_id,omitempty" url:"necessary_reviewer_id,omitempty"`
}

// InviteCommitReviewer 邀请评审人参与 Commit 评审。
func (s *ReviewsService) InviteCommitReviewer(ctx context.Context, pid interface{}, reviewID int, opts *InviteCommitReviewerOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/review/%d/invite", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RemoveCommitReviewerOptions 是 RemoveCommitReviewer 的可选参数。
type RemoveCommitReviewerOptions struct {
	ReviewerID *int `json:"reviewer_id,omitempty" url:"reviewer_id,omitempty"`
}

// RemoveCommitReviewer 移除 Commit 评审人。
func (s *ReviewsService) RemoveCommitReviewer(ctx context.Context, pid interface{}, reviewID int, opts *RemoveCommitReviewerOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/review/%d/dismissals", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// SubmitCommitReviewSummaryOptions 是 SubmitCommitReviewSummary 的可选参数。
type SubmitCommitReviewSummaryOptions struct {
	ReviewerEvent *string `json:"reviewer_event,omitempty" url:"reviewer_event,omitempty"`
	Summary       *string `json:"summary,omitempty" url:"summary,omitempty"`
}

// SubmitCommitReviewSummary 发表 Commit 评审意见。
func (s *ReviewsService) SubmitCommitReviewSummary(ctx context.Context, pid interface{}, reviewID int, opts *SubmitCommitReviewSummaryOptions) (*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/review/%d/reviewer/summary", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var review Review
	resp, err := s.client.Do(req, &review)
	if err != nil {
		return nil, resp, err
	}

	return &review, resp, nil
}

// ReopenCommitReview 重置 Commit 评审状态。
func (s *ReviewsService) ReopenCommitReview(ctx context.Context, pid interface{}, reviewID int) (*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/review/%d/reopen", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var review Review
	resp, err := s.client.Do(req, &review)
	if err != nil {
		return nil, resp, err
	}

	return &review, resp, nil
}

// UpdateCommitReviewOptions 是 UpdateCommitReview 的可选参数。
type UpdateCommitReviewOptions struct {
	Title       *string `json:"title,omitempty" url:"title,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
}

// UpdateCommitReview 更新 Commit 评审。
func (s *ReviewsService) UpdateCommitReview(ctx context.Context, pid interface{}, reviewID int, opts *UpdateCommitReviewOptions) (*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/review/%d", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var review Review
	resp, err := s.client.Do(req, &review)
	if err != nil {
		return nil, resp, err
	}

	return &review, resp, nil
}

// DownloadCommitReviewChangedFiles 下载 Commit Review 差异文件集。
func (s *ReviewsService) DownloadCommitReviewChangedFiles(ctx context.Context, pid interface{}, reviewID int, w io.Writer) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/review/%d/changed_files", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, w)
}
