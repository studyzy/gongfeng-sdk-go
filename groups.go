package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Group 表示工蜂项目组。
type Group struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Path        string     `json:"path,omitempty"`
	WebURL      string     `json:"web_url,omitempty"`
	Description string     `json:"description,omitempty"`
	AvatarURL   string     `json:"avatar_url,omitempty"`
	Projects    []*Project `json:"projects,omitempty"`
}

// GroupsService 处理与工蜂项目组相关的 API。
type GroupsService struct {
	client *Client
}

// CreateGroupOptions 表示 CreateGroup 的可选参数。
type CreateGroupOptions struct {
	Name        *string `json:"name,omitempty"`
	Path        *string `json:"path,omitempty"`
	Description *string `json:"description,omitempty"`
}

// CreateGroup 创建一个新项目组。
func (s *GroupsService) CreateGroup(ctx context.Context, opts *CreateGroupOptions) (*Group, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "groups", opts)
	if err != nil {
		return nil, nil, err
	}

	var g Group
	resp, err := s.client.Do(req, &g)
	if err != nil {
		return nil, resp, err
	}

	return &g, resp, nil
}

// EditGroupOptions 表示 EditGroup 的可选参数。
type EditGroupOptions struct {
	ID          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// EditGroup 编辑项目组。
func (s *GroupsService) EditGroup(ctx context.Context, opts *EditGroupOptions) (*Group, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPut, "groups", opts)
	if err != nil {
		return nil, nil, err
	}

	var g Group
	resp, err := s.client.Do(req, &g)
	if err != nil {
		return nil, resp, err
	}

	return &g, resp, nil
}

// DeleteGroup 删除项目组。
func (s *GroupsService) DeleteGroup(ctx context.Context, gid interface{}) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s", group)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListGroupsOptions 表示 ListGroups 的可选参数。
type ListGroupsOptions struct {
	ListOptions
	Search *string `url:"search,omitempty" json:"search,omitempty"`
}

// ListGroups 获取项目组列表。
func (s *GroupsService) ListGroups(ctx context.Context, opts *ListGroupsOptions) ([]*Group, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "groups", opts)
	if err != nil {
		return nil, nil, err
	}

	var groups []*Group
	resp, err := s.client.Do(req, &groups)
	if err != nil {
		return nil, resp, err
	}

	return groups, resp, nil
}

// GetGroup 获取项目组详情（含下属项目）。
func (s *GroupsService) GetGroup(ctx context.Context, gid interface{}) (*Group, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s", group)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var g Group
	resp, err := s.client.Do(req, &g)
	if err != nil {
		return nil, resp, err
	}

	return &g, resp, nil
}

// ListGroupMembersOptions 表示 ListGroupMembers 的可选参数。
type ListGroupMembersOptions struct {
	ListOptions
}

// ListGroupMembers 获取项目组成员列表。
func (s *GroupsService) ListGroupMembers(ctx context.Context, gid interface{}, opts *ListGroupMembersOptions) ([]*Member, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/members", group)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var members []*Member
	resp, err := s.client.Do(req, &members)
	if err != nil {
		return nil, resp, err
	}

	return members, resp, nil
}

// AddGroupMemberOptions 表示 AddGroupMember 的可选参数。
type AddGroupMemberOptions struct {
	UserID      *int              `json:"user_id,omitempty"`
	AccessLevel *AccessLevelValue `json:"access_level,omitempty"`
}

// AddGroupMember 添加项目组成员。
func (s *GroupsService) AddGroupMember(ctx context.Context, gid interface{}, opts *AddGroupMemberOptions) (*Member, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/members", group)

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var m Member
	resp, err := s.client.Do(req, &m)
	if err != nil {
		return nil, resp, err
	}

	return &m, resp, nil
}

// EditGroupMemberOptions 表示 EditGroupMember 的可选参数。
type EditGroupMemberOptions struct {
	AccessLevel *AccessLevelValue `json:"access_level,omitempty"`
}

// EditGroupMember 修改项目组成员的权限。
func (s *GroupsService) EditGroupMember(ctx context.Context, gid interface{}, userID int, opts *EditGroupMemberOptions) (*Member, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/members/%d", group, userID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var m Member
	resp, err := s.client.Do(req, &m)
	if err != nil {
		return nil, resp, err
	}

	return &m, resp, nil
}

// DeleteGroupMember 移除项目组成员。
func (s *GroupsService) DeleteGroupMember(ctx context.Context, gid interface{}, userID int) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/members/%d", group, userID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
