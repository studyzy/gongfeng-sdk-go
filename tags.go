package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Tag 表示一个 Git 标签。
type Tag struct {
	Name    string  `json:"name,omitempty"`
	Message string  `json:"message,omitempty"`
	Commit  *Commit `json:"commit,omitempty"`
}

// TagsService 处理与 Tag 相关的 API 调用。
type TagsService struct {
	client *Client
}

// ListTagsOptions 是 ListTags 的可选参数。
type ListTagsOptions struct {
	ListOptions
}

// ListTags 获取项目的所有 Tag。
func (s *TagsService) ListTags(ctx context.Context, pid interface{}, opts *ListTagsOptions) ([]*Tag, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/tags", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var tags []*Tag
	resp, err := s.client.Do(req, &tags)
	if err != nil {
		return nil, resp, err
	}

	return tags, resp, nil
}

// GetTag 获取项目中指定名称的 Tag。
func (s *TagsService) GetTag(ctx context.Context, pid interface{}, tagName string) (*Tag, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/tags/%s", project, pathEscape(tagName))

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var tag Tag
	resp, err := s.client.Do(req, &tag)
	if err != nil {
		return nil, resp, err
	}

	return &tag, resp, nil
}

// CreateTagOptions 是 CreateTag 的可选参数。
type CreateTagOptions struct {
	TagName *string `json:"tag_name,omitempty" url:"tag_name,omitempty"`
	Ref     *string `json:"ref,omitempty" url:"ref,omitempty"`
	Message *string `json:"message,omitempty" url:"message,omitempty"`
}

// CreateTag 在项目中创建一个新的 Tag。
func (s *TagsService) CreateTag(ctx context.Context, pid interface{}, opts *CreateTagOptions) (*Tag, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/tags", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var tag Tag
	resp, err := s.client.Do(req, &tag)
	if err != nil {
		return nil, resp, err
	}

	return &tag, resp, nil
}

// DeleteTag 删除项目中指定名称的 Tag。
func (s *TagsService) DeleteTag(ctx context.Context, pid interface{}, tagName string) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/tags/%s", project, pathEscape(tagName))

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
