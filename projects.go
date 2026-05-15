package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Project 表示工蜂项目。
type Project struct {
	ID                   int            `json:"id"`
	Description          string         `json:"description"`
	Public               bool           `json:"public"`
	Archived             bool           `json:"archived"`
	VisibilityLevel      int            `json:"visibility_level"`
	Namespace            *Namespace     `json:"namespace"`
	Name                 string         `json:"name"`
	NameWithNamespace    string         `json:"name_with_namespace"`
	Path                 string         `json:"path"`
	PathWithNamespace    string         `json:"path_with_namespace"`
	DefaultBranch        string         `json:"default_branch"`
	SSHURLToRepo         string         `json:"ssh_url_to_repo"`
	HTTPURLToRepo        string         `json:"http_url_to_repo"`
	HTTPSURLToRepo       string         `json:"https_url_to_repo"`
	WebURL               string         `json:"web_url"`
	TagList              []string       `json:"tag_list"`
	IssuesEnabled        bool           `json:"issues_enabled"`
	MergeRequestsEnabled bool           `json:"merge_requests_enabled"`
	WikiEnabled          bool           `json:"wiki_enabled"`
	SnippetsEnabled      bool           `json:"snippets_enabled"`
	ReviewEnabled        bool           `json:"review_enabled"`
	ForkEnabled          bool           `json:"fork_enabled"`
	CreatedAt            Time           `json:"created_at"`
	LastActivityAt       Time           `json:"last_activity_at"`
	CreatorID            int            `json:"creator_id"`
	AvatarURL            string         `json:"avatar_url"`
	WatchsCount          int            `json:"watchs_count"`
	StarsCount           int            `json:"stars_count"`
	ForksCount           int            `json:"forks_count"`
	ConfigStorage        *ConfigStorage `json:"config_storage"`
	Statistics           *Statistics    `json:"statistics"`
}

// ProjectsService 处理与工蜂项目相关的 API。
type ProjectsService struct {
	client *Client
}

// ListProjectsOptions 表示 ListProjects 的可选参数。
type ListProjectsOptions struct {
	ListOptions
	Search  *string `url:"search,omitempty" json:"search,omitempty"`
	OrderBy *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort    *string `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListProjects 获取项目列表。
func (s *ProjectsService) ListProjects(ctx context.Context, opts *ListProjectsOptions) ([]*Project, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "projects", opts)
	if err != nil {
		return nil, nil, err
	}

	var projects []*Project
	resp, err := s.client.Do(req, &projects)
	if err != nil {
		return nil, resp, err
	}

	return projects, resp, nil
}

// GetProject 获取单个项目的详情。
func (s *ProjectsService) GetProject(ctx context.Context, pid interface{}) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
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

// CreateProjectOptions 表示 CreateProject 的可选参数。
type CreateProjectOptions struct {
	Name                 *string `json:"name,omitempty"`
	Path                 *string `json:"path,omitempty"`
	Description          *string `json:"description,omitempty"`
	NamespaceID          *int    `json:"namespace_id,omitempty"`
	IssuesEnabled        *bool   `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled *bool   `json:"merge_requests_enabled,omitempty"`
	WikiEnabled          *bool   `json:"wiki_enabled,omitempty"`
	Public               *bool   `json:"public,omitempty"`
}

// CreateProject 创建一个新项目。
func (s *ProjectsService) CreateProject(ctx context.Context, opts *CreateProjectOptions) (*Project, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "projects", opts)
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

// SearchProjects 按关键词搜索项目。
func (s *ProjectsService) SearchProjects(ctx context.Context, query string) ([]*Project, *Response, error) {
	u := fmt.Sprintf("projects/search/%s", pathEscape(query))

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var projects []*Project
	resp, err := s.client.Do(req, &projects)
	if err != nil {
		return nil, resp, err
	}

	return projects, resp, nil
}

// ListProjectMembersOptions 表示 ListProjectMembers 的可选参数。
type ListProjectMembersOptions struct {
	ListOptions
}

// ListProjectMembers 获取项目成员列表。
func (s *ProjectsService) ListProjectMembers(ctx context.Context, pid interface{}, opts *ListProjectMembersOptions) ([]*Member, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/members", project)

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

// AddProjectMemberOptions 表示 AddProjectMember 的可选参数。
type AddProjectMemberOptions struct {
	UserID      *int              `json:"user_id,omitempty"`
	AccessLevel *AccessLevelValue `json:"access_level,omitempty"`
}

// AddProjectMember 添加项目成员。
func (s *ProjectsService) AddProjectMember(ctx context.Context, pid interface{}, opts *AddProjectMemberOptions) (*Member, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/members", project)

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

// EditProjectMemberOptions 表示 EditProjectMember 的可选参数。
type EditProjectMemberOptions struct {
	AccessLevel *AccessLevelValue `json:"access_level,omitempty"`
}

// EditProjectMember 修改项目成员的权限。
func (s *ProjectsService) EditProjectMember(ctx context.Context, pid interface{}, userID int, opts *EditProjectMemberOptions) (*Member, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/members/%d", project, userID)

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

// DeleteProjectMember 移除项目成员。
func (s *ProjectsService) DeleteProjectMember(ctx context.Context, pid interface{}, userID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/members/%d", project, userID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
