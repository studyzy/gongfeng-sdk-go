package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Note 表示一条评论。
type Note struct {
	ID           int    `json:"id,omitempty"`
	Body         string `json:"body,omitempty"`
	Attachment   string `json:"attachment,omitempty"`
	Author       *User  `json:"author,omitempty"`
	CreatedAt    Time   `json:"created_at,omitempty"`
	UpdatedAt    Time   `json:"updated_at,omitempty"`
	System       bool   `json:"system,omitempty"`
	NoteableID   int    `json:"noteable_id,omitempty"`
	NoteableType string `json:"noteable_type,omitempty"`
}

// NotesService 处理与 Note 相关的 API 调用。
type NotesService struct {
	client *Client
}

// CreateMergeRequestNoteOptions 是 CreateMergeRequestNote 的可选参数。
type CreateMergeRequestNoteOptions struct {
	Body          *string `json:"body,omitempty" url:"body,omitempty"`
	Path          *string `json:"path,omitempty" url:"path,omitempty"`
	Line          *string `json:"line,omitempty" url:"line,omitempty"`
	LineType      *string `json:"line_type,omitempty" url:"line_type,omitempty"`
	ReviewerState *string `json:"reviewer_state,omitempty" url:"reviewer_state,omitempty"`
}

// CreateMergeRequestNote 为合并请求创建一条评论。
func (s *NotesService) CreateMergeRequestNote(ctx context.Context, pid interface{}, mergeRequestID int, opts *CreateMergeRequestNoteOptions) (*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_requests/%d/notes", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var note Note
	resp, err := s.client.Do(req, &note)
	if err != nil {
		return nil, resp, err
	}

	return &note, resp, nil
}

// ListMergeRequestNotesOptions 是 ListMergeRequestNotes 的可选参数。
type ListMergeRequestNotesOptions struct {
	ListOptions
}

// ListMergeRequestNotes 获取合并请求的评论列表。
func (s *NotesService) ListMergeRequestNotes(ctx context.Context, pid interface{}, mergeRequestID int, opts *ListMergeRequestNotesOptions) ([]*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_requests/%d/notes", project, mergeRequestID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var notes []*Note
	resp, err := s.client.Do(req, &notes)
	if err != nil {
		return nil, resp, err
	}

	return notes, resp, nil
}

// GetMergeRequestNote 获取合并请求中指定 ID 的评论。
func (s *NotesService) GetMergeRequestNote(ctx context.Context, pid interface{}, mergeRequestID, noteID int) (*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_requests/%d/notes/%d", project, mergeRequestID, noteID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var note Note
	resp, err := s.client.Do(req, &note)
	if err != nil {
		return nil, resp, err
	}

	return &note, resp, nil
}

// UpdateMergeRequestNoteOptions 是 UpdateMergeRequestNote 的可选参数。
type UpdateMergeRequestNoteOptions struct {
	Body          *string `json:"body,omitempty" url:"body,omitempty"`
	ReviewerState *string `json:"reviewer_state,omitempty" url:"reviewer_state,omitempty"`
}

// UpdateMergeRequestNote 修改合并请求中指定 ID 的评论。
func (s *NotesService) UpdateMergeRequestNote(ctx context.Context, pid interface{}, mergeRequestID, noteID int, opts *UpdateMergeRequestNoteOptions) (*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/merge_requests/%d/notes/%d", project, mergeRequestID, noteID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var note Note
	resp, err := s.client.Do(req, &note)
	if err != nil {
		return nil, resp, err
	}

	return &note, resp, nil
}

// CreateIssueNoteOptions 是 CreateIssueNote 的可选参数。
type CreateIssueNoteOptions struct {
	Body *string `json:"body,omitempty" url:"body,omitempty"`
}

// CreateIssueNote 为缺陷创建一条评论。
func (s *NotesService) CreateIssueNote(ctx context.Context, pid interface{}, issueID int, opts *CreateIssueNoteOptions) (*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d/notes", project, issueID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var note Note
	resp, err := s.client.Do(req, &note)
	if err != nil {
		return nil, resp, err
	}

	return &note, resp, nil
}

// ListIssueNotesOptions 是 ListIssueNotes 的可选参数。
type ListIssueNotesOptions struct {
	ListOptions
}

// ListIssueNotes 获取缺陷的评论列表。
func (s *NotesService) ListIssueNotes(ctx context.Context, pid interface{}, issueID int, opts *ListIssueNotesOptions) ([]*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d/notes", project, issueID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var notes []*Note
	resp, err := s.client.Do(req, &notes)
	if err != nil {
		return nil, resp, err
	}

	return notes, resp, nil
}

// GetIssueNote 获取缺陷中指定 ID 的评论。
func (s *NotesService) GetIssueNote(ctx context.Context, pid interface{}, issueID, noteID int) (*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d/notes/%d", project, issueID, noteID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var note Note
	resp, err := s.client.Do(req, &note)
	if err != nil {
		return nil, resp, err
	}

	return &note, resp, nil
}

// UpdateIssueNoteOptions 是 UpdateIssueNote 的可选参数。
type UpdateIssueNoteOptions struct {
	Body *string `json:"body,omitempty" url:"body,omitempty"`
}

// UpdateIssueNote 修改缺陷中指定 ID 的评论。
func (s *NotesService) UpdateIssueNote(ctx context.Context, pid interface{}, issueID, noteID int, opts *UpdateIssueNoteOptions) (*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/issues/%d/notes/%d", project, issueID, noteID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var note Note
	resp, err := s.client.Do(req, &note)
	if err != nil {
		return nil, resp, err
	}

	return &note, resp, nil
}

// CreateReviewNoteOptions 是 CreateReviewNote 的可选参数。
type CreateReviewNoteOptions struct {
	Body          *string `json:"body,omitempty" url:"body,omitempty"`
	Path          *string `json:"path,omitempty" url:"path,omitempty"`
	Line          *string `json:"line,omitempty" url:"line,omitempty"`
	LineType      *string `json:"line_type,omitempty" url:"line_type,omitempty"`
	ReviewerState *string `json:"reviewer_state,omitempty" url:"reviewer_state,omitempty"`
}

// ListReviewNotesOptions 是 ListReviewNotes 的可选参数。
type ListReviewNotesOptions struct {
	ListOptions
}

// UpdateReviewNoteOptions 是 UpdateReviewNote 的可选参数。
type UpdateReviewNoteOptions struct {
	Body          *string `json:"body,omitempty" url:"body,omitempty"`
	ReviewerState *string `json:"reviewer_state,omitempty" url:"reviewer_state,omitempty"`
}

// CreateReviewNote 为代码评审创建一条评论。
func (s *NotesService) CreateReviewNote(ctx context.Context, pid interface{}, reviewID int, opts *CreateReviewNoteOptions) (*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/reviews/%d/notes", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var note Note
	resp, err := s.client.Do(req, &note)
	if err != nil {
		return nil, resp, err
	}

	return &note, resp, nil
}

// ListReviewNotes 获取代码评审的评论列表。
func (s *NotesService) ListReviewNotes(ctx context.Context, pid interface{}, reviewID int, opts *ListReviewNotesOptions) ([]*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/reviews/%d/notes", project, reviewID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var notes []*Note
	resp, err := s.client.Do(req, &notes)
	if err != nil {
		return nil, resp, err
	}

	return notes, resp, nil
}

// GetReviewNote 获取代码评审中指定 ID 的评论。
func (s *NotesService) GetReviewNote(ctx context.Context, pid interface{}, reviewID, noteID int) (*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/reviews/%d/notes/%d", project, reviewID, noteID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var note Note
	resp, err := s.client.Do(req, &note)
	if err != nil {
		return nil, resp, err
	}

	return &note, resp, nil
}

// UpdateReviewNote 修改代码评审中指定 ID 的评论。
func (s *NotesService) UpdateReviewNote(ctx context.Context, pid interface{}, reviewID, noteID int, opts *UpdateReviewNoteOptions) (*Note, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/reviews/%d/notes/%d", project, reviewID, noteID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var note Note
	resp, err := s.client.Do(req, &note)
	if err != nil {
		return nil, resp, err
	}

	return &note, resp, nil
}
