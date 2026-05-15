package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// ForksService 处理与 Fork 相关的 API 调用。
type ForksService struct {
	client *Client
}

// ForkProject 将项目 Fork 到当前用户的命名空间。
func (s *ForksService) ForkProject(ctx context.Context, pid interface{}) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/fork/%s", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var p Project
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return &p, resp, nil
}

// CreateForkRelation 在项目之间创建 Fork 关系。
func (s *ForksService) CreateForkRelation(ctx context.Context, pid interface{}, forkedFromID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/fork/%d", project, forkedFromID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteForkRelation 取消项目的 Fork 关系。
func (s *ForksService) DeleteForkRelation(ctx context.Context, pid interface{}) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/fork", project)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
