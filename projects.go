package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Project 表示工蜂项目。
type Project struct {
	ID                   int            `json:"id,omitempty"`
	Description          string         `json:"description,omitempty"`
	Public               bool           `json:"public,omitempty"`
	Archived             bool           `json:"archived,omitempty"`
	VisibilityLevel      int            `json:"visibility_level,omitempty"`
	Namespace            *Namespace     `json:"namespace,omitempty"`
	Name                 string         `json:"name,omitempty"`
	NameWithNamespace    string         `json:"name_with_namespace,omitempty"`
	Path                 string         `json:"path,omitempty"`
	PathWithNamespace    string         `json:"path_with_namespace,omitempty"`
	DefaultBranch        string         `json:"default_branch,omitempty"`
	SSHURLToRepo         string         `json:"ssh_url_to_repo,omitempty"`
	HTTPURLToRepo        string         `json:"http_url_to_repo,omitempty"`
	HTTPSURLToRepo       string         `json:"https_url_to_repo,omitempty"`
	WebURL               string         `json:"web_url,omitempty"`
	TagList              []string       `json:"tag_list,omitempty"`
	IssuesEnabled        bool           `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled bool           `json:"merge_requests_enabled,omitempty"`
	WikiEnabled          bool           `json:"wiki_enabled,omitempty"`
	SnippetsEnabled      bool           `json:"snippets_enabled,omitempty"`
	ReviewEnabled        bool           `json:"review_enabled,omitempty"`
	ForkEnabled          bool           `json:"fork_enabled,omitempty"`
	CreatedAt            Time           `json:"created_at,omitempty"`
	LastActivityAt       Time           `json:"last_activity_at,omitempty"`
	CreatorID            int            `json:"creator_id,omitempty"`
	AvatarURL            string         `json:"avatar_url,omitempty"`
	WatchsCount          int            `json:"watchs_count,omitempty"`
	StarsCount           int            `json:"stars_count,omitempty"`
	ForksCount           int            `json:"forks_count,omitempty"`
	ConfigStorage        *ConfigStorage `json:"config_storage,omitempty"`
	Statistics           *Statistics    `json:"statistics,omitempty"`
}

// ProjectShare 表示项目共享到组的关系。
type ProjectShare struct {
	ProjectID   int  `json:"project_id,omitempty"`
	GroupID     int  `json:"group_id,omitempty"`
	GroupAccess int  `json:"group_access,omitempty"`
	CreatedAt   Time `json:"created_at,omitempty"`
	UpdatedAt   Time `json:"updated_at,omitempty"`
}

// ProjectEvent 表示项目事件。
type ProjectEvent struct {
	Title          string                 `json:"title,omitempty"`
	ProjectID      int                    `json:"project_id,omitempty"`
	ActionName     string                 `json:"action_name,omitempty"`
	TargetID       int                    `json:"target_id,omitempty"`
	TargetType     string                 `json:"target_type,omitempty"`
	AuthorID       int                    `json:"author_id,omitempty"`
	AuthorUsername string                 `json:"author_username,omitempty"`
	CreatedAt      Time                   `json:"created_at,omitempty"`
	TargetTitle    string                 `json:"target_title,omitempty"`
	Data           map[string]interface{} `json:"data,omitempty"`
}

// ProjectStar 表示项目标星关系。
type ProjectStar struct {
	ProjectID int   `json:"project_id,omitempty"`
	User      *User `json:"user,omitempty"`
}

// ProjectsService 处理与工蜂项目相关的 API。
type ProjectsService struct {
	client *Client
}

// ListProjectsOptions 表示 ListProjects 的可选参数。
type ListProjectsOptions struct {
	ListOptions
	Search           *string `url:"search,omitempty" json:"search,omitempty"`
	Archived         *bool   `url:"archived,omitempty" json:"archived,omitempty"`
	WithArchived     *bool   `url:"with_archived,omitempty" json:"with_archived,omitempty"`
	WithPush         *bool   `url:"with_push,omitempty" json:"with_push,omitempty"`
	Abandoned        *bool   `url:"abandoned,omitempty" json:"abandoned,omitempty"`
	VisibilityLevels *string `url:"visibility_levels,omitempty" json:"visibility_levels,omitempty"`
	OrderBy          *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort             *string `url:"sort,omitempty" json:"sort,omitempty"`
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
	ForkEnabled          *bool   `json:"fork_enabled,omitempty"`
	VisibilityLevel      *int    `json:"visibility_level,omitempty"`
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
	Query *string `url:"query,omitempty" json:"query,omitempty"`
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

// GetProjectMember 获取项目中的单个成员。
func (s *ProjectsService) GetProjectMember(ctx context.Context, pid interface{}, userID int) (*Member, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/members/%d", project, userID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var member Member
	resp, err := s.client.Do(req, &member)
	if err != nil {
		return nil, resp, err
	}

	return &member, resp, nil
}

// UpdateProjectOptions 表示 UpdateProject 的可选参数。
type UpdateProjectOptions struct {
	Name                 *string  `json:"name,omitempty"`
	Description          *string  `json:"description,omitempty"`
	DefaultBranch        *string  `json:"default_branch,omitempty"`
	LimitFileSize        *float64 `json:"limit_file_size,omitempty"`
	LimitLFSFileSize     *float64 `json:"limit_lfs_file_size,omitempty"`
	IssuesEnabled        *bool    `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled *bool    `json:"merge_requests_enabled,omitempty"`
	WikiEnabled          *bool    `json:"wiki_enabled,omitempty"`
	ReviewEnabled        *bool    `json:"review_enabled,omitempty"`
	ForkEnabled          *bool    `json:"fork_enabled,omitempty"`
	TagNameRegex         *string  `json:"tag_name_regex,omitempty"`
	TagCreatePushLevel   *int     `json:"tag_create_push_level,omitempty"`
	VisibilityLevel      *int     `json:"visibility_level,omitempty"`
}

// UpdateProject 修改项目设置。
func (s *ProjectsService) UpdateProject(ctx context.Context, pid interface{}, opts *UpdateProjectOptions) (*Project, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s", project)

	req, err := s.client.NewRequest(ctx, http.MethodPut, u, opts)
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

// ListOwnedProjects 获取当前用户拥有的项目列表。
func (s *ProjectsService) ListOwnedProjects(ctx context.Context, opts *ListProjectsOptions) ([]*Project, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "projects/owned", opts)
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

// DeleteProject 删除项目。
func (s *ProjectsService) DeleteProject(ctx context.Context, pid interface{}) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s", project)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ShareProjectOptions 表示 ShareProject 的可选参数。
type ShareProjectOptions struct {
	GroupID     *int `json:"group_id,omitempty"`
	GroupAccess *int `json:"group_access,omitempty"`
}

// ShareProject 将项目共享给组。
func (s *ProjectsService) ShareProject(ctx context.Context, pid interface{}, opts *ShareProjectOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/share", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListProjectShares 获取项目共享的组列表。
func (s *ProjectsService) ListProjectShares(ctx context.Context, pid interface{}) ([]*ProjectShare, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/shares", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var shares []*ProjectShare
	resp, err := s.client.Do(req, &shares)
	if err != nil {
		return nil, resp, err
	}

	return shares, resp, nil
}

// DeleteProjectShare 删除项目共享关系。
func (s *ProjectsService) DeleteProjectShare(ctx context.Context, pid interface{}, groupID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/share/%d", project, groupID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListProjectEventsOptions 表示 ListProjectEvents 的可选参数。
type ListProjectEventsOptions struct {
	ListOptions
	UserIDOrName interface{}
}

// ListProjectEvents 获取项目事件列表。
func (s *ProjectsService) ListProjectEvents(ctx context.Context, pid interface{}, opts *ListProjectEventsOptions) ([]*ProjectEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/events", project)

	var reqOpts *struct {
		ListOptions
		UserIDOrName *string `url:"user_id_or_name,omitempty" json:"user_id_or_name,omitempty"`
	}
	if opts != nil {
		reqOpts = &struct {
			ListOptions
			UserIDOrName *string `url:"user_id_or_name,omitempty" json:"user_id_or_name,omitempty"`
		}{
			ListOptions: opts.ListOptions,
		}
		if opts.UserIDOrName != nil {
			id, err := parseID(opts.UserIDOrName)
			if err != nil {
				return nil, nil, err
			}
			reqOpts.UserIDOrName = Ptr(id)
		}
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, reqOpts)
	if err != nil {
		return nil, nil, err
	}

	var events []*ProjectEvent
	resp, err := s.client.Do(req, &events)
	if err != nil {
		return nil, resp, err
	}

	return events, resp, nil
}

// StarProject 对项目标星。
func (s *ProjectsService) StarProject(ctx context.Context, pid interface{}) (bool, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return false, nil, err
	}
	u := fmt.Sprintf("projects/%s/star", project)

	req, err := s.client.NewRequest(ctx, http.MethodPut, u, nil)
	if err != nil {
		return false, nil, err
	}

	var starred bool
	resp, err := s.client.Do(req, &starred)
	if err != nil {
		return false, resp, err
	}

	return starred, resp, nil
}

// UnstarProject 取消项目标星。
func (s *ProjectsService) UnstarProject(ctx context.Context, pid interface{}) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/star", project)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// GetStarStatus 查询项目标星状态。
func (s *ProjectsService) GetStarStatus(ctx context.Context, pid interface{}) (bool, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return false, nil, err
	}
	u := fmt.Sprintf("projects/%s/star", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return false, nil, err
	}

	var starred bool
	resp, err := s.client.Do(req, &starred)
	if err != nil {
		return false, resp, err
	}

	return starred, resp, nil
}

// ListProjectStarsOptions 表示 ListProjectStars 的可选参数。
type ListProjectStarsOptions struct {
	ListOptions
}

// ListProjectStars 获取项目标星用户列表。
func (s *ProjectsService) ListProjectStars(ctx context.Context, pid interface{}, opts *ListProjectStarsOptions) ([]*ProjectStar, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/stars", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var stars []*ProjectStar
	resp, err := s.client.Do(req, &stars)
	if err != nil {
		return nil, resp, err
	}

	return stars, resp, nil
}
