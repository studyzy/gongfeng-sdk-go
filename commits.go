package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// CommitComment 表示提交的评论。
type CommitComment struct {
	Note     string `json:"note"`
	Author   *User  `json:"author"`
	Path     string `json:"path"`
	Line     int    `json:"line"`
	LineType string `json:"line_type"`
}

// CreateCommitCommentOptions 表示 CreateCommitComment 的可选参数。
type CreateCommitCommentOptions struct {
	Note     *string `json:"note,omitempty" url:"note,omitempty"`
	Path     *string `json:"path,omitempty" url:"path,omitempty"`
	Line     *int    `json:"line,omitempty" url:"line,omitempty"`
	LineType *string `json:"line_type,omitempty" url:"line_type,omitempty"`
}

// CommitRef 表示提交对应的分支或 tag 引用。
type CommitRef struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// CommitsService 处理与工蜂提交相关的 API。
type CommitsService struct {
	client *Client
}

// GetCommit 获取单个提交的详情。
func (s *CommitsService) GetCommit(ctx context.Context, pid interface{}, sha string) (*Commit, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s", project, pathEscape(sha))

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var c Commit
	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return &c, resp, nil
}

// GetCommitDiffOptions 表示 GetCommitDiff 的可选参数。
type GetCommitDiffOptions struct {
	Path             *string `url:"path,omitempty" json:"path,omitempty"`
	IgnoreWhiteSpace *bool   `url:"ignore_white_space,omitempty" json:"ignore_white_space,omitempty"`
}

// GetCommitDiff 获取提交的 Diff。
func (s *CommitsService) GetCommitDiff(ctx context.Context, pid interface{}, sha string, opts *GetCommitDiffOptions) ([]*Diff, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/diff", project, pathEscape(sha))

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var diffs []*Diff
	resp, err := s.client.Do(req, &diffs)
	if err != nil {
		return nil, resp, err
	}

	return diffs, resp, nil
}

// CreateCommitComment 为提交添加评论。
func (s *CommitsService) CreateCommitComment(ctx context.Context, pid interface{}, sha string, opts *CreateCommitCommentOptions) (*CommitComment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/comments", project, pathEscape(sha))

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var comment CommitComment
	resp, err := s.client.Do(req, &comment)
	if err != nil {
		return nil, resp, err
	}

	return &comment, resp, nil
}

// ListCommitCommentsOptions 表示 ListCommitComments 的可选参数。
type ListCommitCommentsOptions struct {
	ListOptions
}

// ListCommitComments 获取提交的评论列表。
func (s *CommitsService) ListCommitComments(ctx context.Context, pid interface{}, sha string, opts *ListCommitCommentsOptions) ([]*CommitComment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/comments", project, pathEscape(sha))

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var comments []*CommitComment
	resp, err := s.client.Do(req, &comments)
	if err != nil {
		return nil, resp, err
	}

	return comments, resp, nil
}

// ListCommitsOptions 表示 ListCommits 的可选参数。
type ListCommitsOptions struct {
	ListOptions
	RefName *string `url:"ref_name,omitempty" json:"ref_name,omitempty"`
	Path    *string `url:"path,omitempty" json:"path,omitempty"`
	Since   *string `url:"since,omitempty" json:"since,omitempty"`
	Until   *string `url:"until,omitempty" json:"until,omitempty"`
}

// ListCommits 获取项目的提交列表。
func (s *CommitsService) ListCommits(ctx context.Context, pid interface{}, opts *ListCommitsOptions) ([]*Commit, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits", project)

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

// ListCommitRefsOptions 表示 ListCommitRefs 的可选参数。
type ListCommitRefsOptions struct {
	ListOptions
	Type *string `url:"type,omitempty" json:"type,omitempty"`
}

// ListCommitRefs 获取提交对应的分支和 tag 引用。
func (s *CommitsService) ListCommitRefs(ctx context.Context, pid interface{}, sha string, opts *ListCommitRefsOptions) ([]*CommitRef, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/commits/%s/refs", project, pathEscape(sha))

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var refs []*CommitRef
	resp, err := s.client.Do(req, &refs)
	if err != nil {
		return nil, resp, err
	}

	return refs, resp, nil
}
