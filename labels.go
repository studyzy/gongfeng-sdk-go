package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Label 表示一个项目标签。
type Label struct {
	Name                   string `json:"name,omitempty"`
	Color                  string `json:"color,omitempty"`
	Description            string `json:"description,omitempty"`
	OpenIssuesCount        int    `json:"open_issues_count,omitempty"`
	ClosedIssuesCount      int    `json:"closed_issues_count,omitempty"`
	OpenMergeRequestsCount int    `json:"open_merge_requests_count,omitempty"`
}

// LabelsService 处理与 Label 相关的 API 调用。
type LabelsService struct {
	client *Client
}

// CreateLabelOptions 是 CreateLabel 的可选参数。
type CreateLabelOptions struct {
	Name        *string `json:"name,omitempty" url:"name,omitempty"`
	Color       *string `json:"color,omitempty" url:"color,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
}

// CreateLabel 在项目中创建一个新标签。
func (s *LabelsService) CreateLabel(ctx context.Context, pid interface{}, opts *CreateLabelOptions) (*Label, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/labels", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var label Label
	resp, err := s.client.Do(req, &label)
	if err != nil {
		return nil, resp, err
	}

	return &label, resp, nil
}

// UpdateLabelOptions 是 UpdateLabel 的可选参数。
type UpdateLabelOptions struct {
	Name        *string `json:"name,omitempty" url:"name,omitempty"`
	NewName     *string `json:"new_name,omitempty" url:"new_name,omitempty"`
	Color       *string `json:"color,omitempty" url:"color,omitempty"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
}

// UpdateLabel 更新项目中的一个标签。
func (s *LabelsService) UpdateLabel(ctx context.Context, pid interface{}, opts *UpdateLabelOptions) (*Label, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/labels", project)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var label Label
	resp, err := s.client.Do(req, &label)
	if err != nil {
		return nil, resp, err
	}

	return &label, resp, nil
}

// DeleteLabelOptions 是 DeleteLabel 的可选参数。
type DeleteLabelOptions struct {
	Name *string `json:"name,omitempty" url:"name,omitempty"`
}

// DeleteLabel 删除项目中的一个标签。
func (s *LabelsService) DeleteLabel(ctx context.Context, pid interface{}, opts *DeleteLabelOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/labels", project)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListLabelsOptions 是 ListLabels 的可选参数。
type ListLabelsOptions struct {
	ListOptions
}

// ListLabels 获取项目的标签列表。
func (s *LabelsService) ListLabels(ctx context.Context, pid interface{}, opts *ListLabelsOptions) ([]*Label, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/labels", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var labels []*Label
	resp, err := s.client.Do(req, &labels)
	if err != nil {
		return nil, resp, err
	}

	return labels, resp, nil
}
