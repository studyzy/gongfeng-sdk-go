package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// WatchersService 处理与项目关注者相关的 API 调用。
type WatchersService struct {
	client *Client
}

// ListWatchersOptions 是 ListWatchers 的可选参数。
type ListWatchersOptions struct {
	ListOptions
}

// ListWatchers 获取项目的关注者列表。
func (s *WatchersService) ListWatchers(ctx context.Context, pid interface{}, opts *ListWatchersOptions) ([]*User, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/watchers", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

// WatchProject 关注一个项目。
func (s *WatchersService) WatchProject(ctx context.Context, pid interface{}) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/watchers", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnwatchProject 取消关注一个项目。
func (s *WatchersService) UnwatchProject(ctx context.Context, pid interface{}) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/watchers", project)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
