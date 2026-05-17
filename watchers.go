package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Watcher 表示项目关注关系。
type Watcher struct {
	ProjectID int   `json:"project_id,omitempty"`
	Mute      bool  `json:"mute,omitempty"`
	User      *User `json:"user,omitempty"`
}

// WatchersService 处理与项目关注者相关的 API 调用。
type WatchersService struct {
	client *Client
}

// ListWatchersOptions 是 ListWatchers 的可选参数。
type ListWatchersOptions struct {
	ListOptions
}

// ListWatchers 获取项目的关注者列表。
func (s *WatchersService) ListWatchers(ctx context.Context, pid interface{}, opts *ListWatchersOptions) ([]*Watcher, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/watchers", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var watchers []*Watcher
	resp, err := s.client.Do(req, &watchers)
	if err != nil {
		return nil, resp, err
	}

	return watchers, resp, nil
}

// WatchProjectOptions 是 WatchProject 的可选参数。
type WatchProjectOptions struct {
	Mute *bool `json:"mute,omitempty" url:"mute,omitempty"`
}

// GetWatchStatus 获取当前用户是否关注项目。
func (s *WatchersService) GetWatchStatus(ctx context.Context, pid interface{}) (bool, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return false, nil, err
	}
	path := fmt.Sprintf("projects/%s/watch", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return false, nil, err
	}

	var watched bool
	resp, err := s.client.Do(req, &watched)
	if err != nil {
		return false, resp, err
	}

	return watched, resp, nil
}

// WatchProject 关注一个项目。
func (s *WatchersService) WatchProject(ctx context.Context, pid interface{}, opts *WatchProjectOptions) (*Watcher, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/watch", project)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var watcher Watcher
	resp, err := s.client.Do(req, &watcher)
	if err != nil {
		return nil, resp, err
	}

	return &watcher, resp, nil
}

// UnwatchProject 取消关注一个项目。
func (s *WatchersService) UnwatchProject(ctx context.Context, pid interface{}) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/watch", project)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
