package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// ReleaseRef 表示 Release 关联的 Tag 引用信息。
type ReleaseRef struct {
	Name    string  `json:"name,omitempty"`
	Message string  `json:"message,omitempty"`
	Commit  *Commit `json:"commit,omitempty"`
}

// Release 表示一个项目发布版本。
type Release struct {
	ID          int          `json:"id,omitempty"`
	ProjectID   int          `json:"project_id,omitempty"`
	Tag         string       `json:"tag,omitempty"`
	Title       string       `json:"title,omitempty"`
	Type        string       `json:"type,omitempty"`
	Description string       `json:"description,omitempty"`
	Attachments interface{}  `json:"attachments,omitempty"`
	Ref         *ReleaseRef  `json:"ref,omitempty"`
	CreatedAt   Time         `json:"created_at,omitempty"`
	UpdatedAt   Time         `json:"updated_at,omitempty"`
	Author      *User        `json:"author,omitempty"`
}

// ReleasesService 处理与 Release 相关的 API 调用。
type ReleasesService struct {
	client *Client
}

// ListReleasesOptions 是 ListReleases 的可选参数。
type ListReleasesOptions struct {
	ListOptions
}

// ListReleases 获取项目的所有 Release。
func (s *ReleasesService) ListReleases(ctx context.Context, pid interface{}, opts *ListReleasesOptions) ([]*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/releases", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var releases []*Release
	resp, err := s.client.Do(req, &releases)
	if err != nil {
		return nil, resp, err
	}

	return releases, resp, nil
}

// GetRelease 获取项目中指定 ID 的 Release。
func (s *ReleasesService) GetRelease(ctx context.Context, pid interface{}, releaseID int) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/releases/%d", project, releaseID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var release Release
	resp, err := s.client.Do(req, &release)
	if err != nil {
		return nil, resp, err
	}

	return &release, resp, nil
}

// CreateReleaseOptions 是 CreateRelease 的可选参数。
type CreateReleaseOptions struct {
	Tag         *string `json:"tag,omitempty" url:"tag,omitempty"`
	StartPoint  *string `json:"start_point,omitempty" url:"start_point,omitempty"`
	Title       *string `json:"title,omitempty" url:"title,omitempty"`
	Type        *string `json:"type,omitempty" url:"type,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
}

// CreateRelease 在项目中创建一个新的 Release。
func (s *ReleasesService) CreateRelease(ctx context.Context, pid interface{}, opts *CreateReleaseOptions) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/releases", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var release Release
	resp, err := s.client.Do(req, &release)
	if err != nil {
		return nil, resp, err
	}

	return &release, resp, nil
}

// UpdateReleaseOptions 是 UpdateRelease 的可选参数。
type UpdateReleaseOptions struct {
	Description *string `json:"description,omitempty" url:"description,omitempty"`
}

// UpdateRelease 更新项目中指定 ID 的 Release。
func (s *ReleasesService) UpdateRelease(ctx context.Context, pid interface{}, releaseID int, opts *UpdateReleaseOptions) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/releases/%d", project, releaseID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var release Release
	resp, err := s.client.Do(req, &release)
	if err != nil {
		return nil, resp, err
	}

	return &release, resp, nil
}

// DeleteRelease 删除项目中指定 ID 的 Release。
func (s *ReleasesService) DeleteRelease(ctx context.Context, pid interface{}, releaseID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/releases/%d", project, releaseID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
