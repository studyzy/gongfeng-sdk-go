package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Branch 表示工蜂代码分支。
type Branch struct {
	Name               string  `json:"name,omitempty"`
	Protected          bool    `json:"protected,omitempty"`
	DevelopersCanPush  bool    `json:"developers_can_push,omitempty"`
	DevelopersCanMerge bool    `json:"developers_can_merge,omitempty"`
	Commit             *Commit `json:"commit,omitempty"`
}

// BranchesService 处理与工蜂分支相关的 API。
type BranchesService struct {
	client *Client
}

// CreateBranchOptions 表示 CreateBranch 的可选参数。
type CreateBranchOptions struct {
	BranchName *string `json:"branch_name,omitempty"`
	Ref        *string `json:"ref,omitempty"`
}

// CreateBranch 创建一个新分支。
func (s *BranchesService) CreateBranch(ctx context.Context, pid interface{}, opts *CreateBranchOptions) (*Branch, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/branches", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var b Branch
	resp, err := s.client.Do(req, &b)
	if err != nil {
		return nil, resp, err
	}

	return &b, resp, nil
}

// DeleteBranch 删除分支。
func (s *BranchesService) DeleteBranch(ctx context.Context, pid interface{}, branch string) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/branches/%s", project, pathEscape(branch))

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListBranchesOptions 表示 ListBranches 的可选参数。
type ListBranchesOptions struct {
	ListOptions
}

// ListBranches 获取项目分支列表。
func (s *BranchesService) ListBranches(ctx context.Context, pid interface{}, opts *ListBranchesOptions) ([]*Branch, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/branches", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var branches []*Branch
	resp, err := s.client.Do(req, &branches)
	if err != nil {
		return nil, resp, err
	}

	return branches, resp, nil
}

// GetBranch 获取单个分支的详情。
func (s *BranchesService) GetBranch(ctx context.Context, pid interface{}, branch string) (*Branch, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/branches/%s", project, pathEscape(branch))

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var b Branch
	resp, err := s.client.Do(req, &b)
	if err != nil {
		return nil, resp, err
	}

	return &b, resp, nil
}

// ProtectedBranch 表示保护分支的详情。
type ProtectedBranch struct {
	Name               string `json:"name,omitempty"`
	DevelopersCanPush  bool   `json:"developers_can_push,omitempty"`
	DevelopersCanMerge bool   `json:"developers_can_merge,omitempty"`
}

// GetProtectedBranch 获取保护分支详情。
func (s *BranchesService) GetProtectedBranch(ctx context.Context, pid interface{}, branch string) (*ProtectedBranch, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/repository/branches/%s/protect", project, pathEscape(branch))

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var pb ProtectedBranch
	resp, err := s.client.Do(req, &pb)
	if err != nil {
		return nil, resp, err
	}

	return &pb, resp, nil
}
