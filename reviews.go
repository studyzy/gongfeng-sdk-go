package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Review 表示一个 Commit 评审。
type Review struct {
	ID          int    `json:"id"`
	ProjectID   int    `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string `json:"state"`
	CreatedAt   Time   `json:"created_at"`
	UpdatedAt   Time   `json:"updated_at"`
	Author      *User  `json:"author"`
}

// Reviewer 表示一个评审人。
type Reviewer struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	State       string `json:"state"`
	AvatarURL   string `json:"avatar_url"`
	ReviewState string `json:"review_state"`
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
	path := fmt.Sprintf("projects/%s/merge_request/%d/reviewers", project, mergeRequestID)

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
	path := fmt.Sprintf("projects/%s/merge_request/%d/reviewers", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ApproveMROptions 是 ApproveMR 的可选参数。
type ApproveMROptions struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`
}

// ApproveMR 通过 MR 评审。
func (s *ReviewsService) ApproveMR(ctx context.Context, pid interface{}, mergeRequestID int, opts *ApproveMROptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_request/%d/review/approve", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RejectMROptions 是 RejectMR 的可选参数。
type RejectMROptions struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`
}

// RejectMR 拒绝 MR 评审。
func (s *ReviewsService) RejectMR(ctx context.Context, pid interface{}, mergeRequestID int, opts *RejectMROptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_request/%d/review/reject", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// --- Commit 评审方法 ---

// CreateCommitReviewOptions 是 CreateCommitReview 的请求参数。
type CreateCommitReviewOptions struct {
	Commit                *string `json:"commit,omitempty" url:"commit,omitempty"`
	Title                 *string `json:"title,omitempty" url:"title,omitempty"`
	Description           *string `json:"description,omitempty" url:"description,omitempty"`
	Reviewers             []int   `json:"reviewers,omitempty" url:"reviewers,omitempty"`
	NecessaryReviewers    []int   `json:"necessary_reviewers,omitempty" url:"necessary_reviewers,omitempty"`
	ApproverRule          *int    `json:"approver_rule,omitempty" url:"approver_rule,omitempty"`
	NecessaryApproverRule *int    `json:"necessary_approver_rule,omitempty" url:"necessary_approver_rule,omitempty"`
}

// CreateCommitReview 新建一个 Commit 评审。
func (s *ReviewsService) CreateCommitReview(ctx context.Context, pid interface{}, opts *CreateCommitReviewOptions) (*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/commit_review", project)

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
	State   *string `url:"state,omitempty" json:"state,omitempty"`
	OrderBy *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort    *string `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListCommitReviews 获取项目的 Commit 评审列表。
func (s *ReviewsService) ListCommitReviews(ctx context.Context, pid interface{}, opts *ListCommitReviewsOptions) ([]*Review, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/commit_reviews", project)

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
	path := fmt.Sprintf("projects/%s/commit_reviews/%d", project, reviewID)

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
	path := fmt.Sprintf("projects/%s/commit_reviews/%d/reviewers", project, reviewID)

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
	path := fmt.Sprintf("projects/%s/commit_reviews/%d/reviewers", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ApproveCommitReviewOptions 是 ApproveCommitReview 的可选参数。
type ApproveCommitReviewOptions struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`
}

// ApproveCommitReview 通过 Commit 评审。
func (s *ReviewsService) ApproveCommitReview(ctx context.Context, pid interface{}, reviewID int, opts *ApproveCommitReviewOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/commit_reviews/%d/review/approve", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RejectCommitReviewOptions 是 RejectCommitReview 的可选参数。
type RejectCommitReviewOptions struct {
	Message *string `json:"message,omitempty" url:"message,omitempty"`
}

// RejectCommitReview 拒绝 Commit 评审。
func (s *ReviewsService) RejectCommitReview(ctx context.Context, pid interface{}, reviewID int, opts *RejectCommitReviewOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/commit_reviews/%d/review/reject", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
