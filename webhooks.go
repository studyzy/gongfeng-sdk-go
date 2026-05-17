package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Webhook 表示一个项目的回调钩子。
type Webhook struct {
	ID                    int    `json:"id,omitempty"`
	URL                   string `json:"url,omitempty"`
	ProjectID             int    `json:"project_id,omitempty"`
	PushEvents            bool   `json:"push_events,omitempty"`
	IssuesEvents          bool   `json:"issues_events,omitempty"`
	MergeRequestsEvents   bool   `json:"merge_requests_events,omitempty"`
	TagPushEvents         bool   `json:"tag_push_events,omitempty"`
	NoteEvents            bool   `json:"note_events,omitempty"`
	EnableSSLVerification bool   `json:"enable_ssl_verification,omitempty"`
	CreatedAt             Time   `json:"created_at,omitempty"`
}

// WebhooksService 处理与项目 Webhook 相关的 API 调用。
type WebhooksService struct {
	client *Client
}

// AddWebhookOptions 是 AddWebhook 的可选参数。
type AddWebhookOptions struct {
	URL                   *string `json:"url,omitempty" url:"url,omitempty"`
	PushEvents            *bool   `json:"push_events,omitempty" url:"push_events,omitempty"`
	IssuesEvents          *bool   `json:"issues_events,omitempty" url:"issues_events,omitempty"`
	MergeRequestsEvents   *bool   `json:"merge_requests_events,omitempty" url:"merge_requests_events,omitempty"`
	TagPushEvents         *bool   `json:"tag_push_events,omitempty" url:"tag_push_events,omitempty"`
	NoteEvents            *bool   `json:"note_events,omitempty" url:"note_events,omitempty"`
	EnableSSLVerification *bool   `json:"enable_ssl_verification,omitempty" url:"enable_ssl_verification,omitempty"`
}

// AddWebhook 为项目添加一个回调钩子。
func (s *WebhooksService) AddWebhook(ctx context.Context, pid interface{}, opts *AddWebhookOptions) (*Webhook, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/hooks", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var hook Webhook
	resp, err := s.client.Do(req, &hook)
	if err != nil {
		return nil, resp, err
	}

	return &hook, resp, nil
}

// ListWebhooksOptions 是 ListWebhooks 的可选参数。
type ListWebhooksOptions struct {
	ListOptions
}

// ListWebhooks 获取项目的所有 Webhook。
func (s *WebhooksService) ListWebhooks(ctx context.Context, pid interface{}, opts *ListWebhooksOptions) ([]*Webhook, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/hooks", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var hooks []*Webhook
	resp, err := s.client.Do(req, &hooks)
	if err != nil {
		return nil, resp, err
	}

	return hooks, resp, nil
}

// GetWebhook 获取项目中指定 ID 的 Webhook。
func (s *WebhooksService) GetWebhook(ctx context.Context, pid interface{}, hookID int) (*Webhook, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/hooks/%d", project, hookID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var hook Webhook
	resp, err := s.client.Do(req, &hook)
	if err != nil {
		return nil, resp, err
	}

	return &hook, resp, nil
}

// EditWebhookOptions 是 EditWebhook 的可选参数。
type EditWebhookOptions struct {
	URL                   *string `json:"url,omitempty" url:"url,omitempty"`
	PushEvents            *bool   `json:"push_events,omitempty" url:"push_events,omitempty"`
	IssuesEvents          *bool   `json:"issues_events,omitempty" url:"issues_events,omitempty"`
	MergeRequestsEvents   *bool   `json:"merge_requests_events,omitempty" url:"merge_requests_events,omitempty"`
	TagPushEvents         *bool   `json:"tag_push_events,omitempty" url:"tag_push_events,omitempty"`
	NoteEvents            *bool   `json:"note_events,omitempty" url:"note_events,omitempty"`
	EnableSSLVerification *bool   `json:"enable_ssl_verification,omitempty" url:"enable_ssl_verification,omitempty"`
}

// EditWebhook 编辑项目中指定 ID 的 Webhook。
func (s *WebhooksService) EditWebhook(ctx context.Context, pid interface{}, hookID int, opts *EditWebhookOptions) (*Webhook, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/hooks/%d", project, hookID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var hook Webhook
	resp, err := s.client.Do(req, &hook)
	if err != nil {
		return nil, resp, err
	}

	return &hook, resp, nil
}

// DeleteWebhook 删除项目中指定 ID 的 Webhook。
func (s *WebhooksService) DeleteWebhook(ctx context.Context, pid interface{}, hookID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/hooks/%d", project, hookID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
