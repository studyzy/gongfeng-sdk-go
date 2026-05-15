package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// UsersService 处理与用户相关的 API 调用。
type UsersService struct {
	client *Client
}

// ListUsersOptions 是 ListUsers 的可选参数。
type ListUsersOptions struct {
	ListOptions
	Search *string `url:"search,omitempty" json:"search,omitempty"`
}

// ListUsers 获取用户列表。
func (s *UsersService) ListUsers(ctx context.Context, opts *ListUsersOptions) ([]*User, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "users", opts)
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

// GetUser 获取指定 ID 的用户信息。
func (s *UsersService) GetUser(ctx context.Context, uid int) (*UserDetail, *Response, error) {
	path := fmt.Sprintf("users/%d", uid)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var user UserDetail
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}

// GetCurrentUser 获取当前认证用户的信息。
func (s *UsersService) GetCurrentUser(ctx context.Context) (*UserDetail, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "user", nil)
	if err != nil {
		return nil, nil, err
	}

	var user UserDetail
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}
