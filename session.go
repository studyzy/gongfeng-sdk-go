package gongfeng

import (
	"context"
	"net/http"
)

// Session 表示通过登录获取的用户会话信息。
type Session struct {
	ID           int    `json:"id,omitempty"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email,omitempty"`
	Name         string `json:"name,omitempty"`
	PrivateToken string `json:"private_token,omitempty"`
	State        string `json:"state,omitempty"`
	AvatarURL    string `json:"avatar_url,omitempty"`
}

// SessionService 处理与用户会话相关的 API 调用。
type SessionService struct {
	client *Client
}

// GetSessionOptions 是 GetSession 的请求参数。
type GetSessionOptions struct {
	Login    *string `json:"login,omitempty" url:"login,omitempty"`
	Email    *string `json:"email,omitempty" url:"email,omitempty"`
	Password *string `json:"password,omitempty" url:"password,omitempty"`
}

// GetSession 通过登录凭据获取用户的私有令牌。
func (s *SessionService) GetSession(ctx context.Context, opts *GetSessionOptions) (*Session, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "session", opts)
	if err != nil {
		return nil, nil, err
	}

	var session Session
	resp, err := s.client.Do(req, &session)
	if err != nil {
		return nil, resp, err
	}

	return &session, resp, nil
}
