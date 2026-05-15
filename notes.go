package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Note 表示一条评论。
type Note struct {
	ID           int    `json:"id"`
	Body         string `json:"body"`
	Attachment   string `json:"attachment"`
	Author       *User  `json:"author"`
	CreatedAt    Time   `json:"created_at"`
	UpdatedAt    Time   `json:"updated_at"`
	System       bool   `json:"system"`
	NoteableID   int    `json:"noteable_id"`
	NoteableType string `json:"noteable_type"`
}

// NotesService 处理与 Note 相关的 API 调用。
type NotesService struct {
	client *Client
}

// CreateMergeRequestNoteOptions 是 CreateMergeRequestNote 的可选参数。
type CreateMergeRequestNoteOptions struct {
	Body *string `json:"body,omitempty" url:"body,omitempty"`
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
	Body *string `json:"body,omitempty" url:"body,omitempty"`
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
