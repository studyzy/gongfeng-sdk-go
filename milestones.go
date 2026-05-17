package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// Milestone 表示工蜂里程碑。
type Milestone struct {
	ID          int    `json:"id,omitempty"`
	IID         int    `json:"iid,omitempty"`
	ProjectID   int    `json:"project_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
	DueDate     Time   `json:"due_date,omitempty"`
	CreatedAt   Time   `json:"created_at,omitempty"`
	UpdatedAt   Time   `json:"updated_at,omitempty"`
}

// MilestonesService 处理与工蜂里程碑相关的 API。
type MilestonesService struct {
	client *Client
}

// CreateMilestoneOptions 表示 CreateMilestone 的可选参数。
type CreateMilestoneOptions struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	DueDate     *string `json:"due_date,omitempty"`
	StateEvent  *string `json:"state_event,omitempty"`
}

// CreateMilestone 新增里程碑。
func (s *MilestonesService) CreateMilestone(ctx context.Context, pid interface{}, opts *CreateMilestoneOptions) (*Milestone, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/milestones", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var m Milestone
	resp, err := s.client.Do(req, &m)
	if err != nil {
		return nil, resp, err
	}

	return &m, resp, nil
}

// EditMilestoneOptions 表示 EditMilestone 的可选参数。
type EditMilestoneOptions struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	DueDate     *string `json:"due_date,omitempty"`
	StateEvent  *string `json:"state_event,omitempty"`
}

// EditMilestone 编辑里程碑。
func (s *MilestonesService) EditMilestone(ctx context.Context, pid interface{}, milestoneID int, opts *EditMilestoneOptions) (*Milestone, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/milestones/%d", project, milestoneID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var m Milestone
	resp, err := s.client.Do(req, &m)
	if err != nil {
		return nil, resp, err
	}

	return &m, resp, nil
}

// ListMilestonesOptions 表示 ListMilestones 的可选参数。
type ListMilestonesOptions struct {
	ListOptions
}

// ListMilestones 获取项目的里程碑列表。
func (s *MilestonesService) ListMilestones(ctx context.Context, pid interface{}, opts *ListMilestonesOptions) ([]*Milestone, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/milestones", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var milestones []*Milestone
	resp, err := s.client.Do(req, &milestones)
	if err != nil {
		return nil, resp, err
	}

	return milestones, resp, nil
}

// GetMilestone 获取单个里程碑。
func (s *MilestonesService) GetMilestone(ctx context.Context, pid interface{}, milestoneID int) (*Milestone, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/milestones/%d", project, milestoneID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var m Milestone
	resp, err := s.client.Do(req, &m)
	if err != nil {
		return nil, resp, err
	}

	return &m, resp, nil
}

// DeleteMilestone 删除里程碑。
func (s *MilestonesService) DeleteMilestone(ctx context.Context, pid interface{}, milestoneID int) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/milestones/%d", project, milestoneID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// MilestoneIssue 表示里程碑下的缺陷。
type MilestoneIssue struct {
	ID          int      `json:"id,omitempty"`
	IID         int      `json:"iid,omitempty"`
	ProjectID   int      `json:"project_id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	State       string   `json:"state,omitempty"`
	CreatedAt   Time     `json:"created_at,omitempty"`
	UpdatedAt   Time     `json:"updated_at,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Author      *User    `json:"author,omitempty"`
	Assignee    *User    `json:"assignee,omitempty"`
}

// ListMilestoneIssuesOptions 表示 ListMilestoneIssues 的可选参数。
type ListMilestoneIssuesOptions struct {
	ListOptions
}

// ListMilestoneIssues 获取里程碑下的缺陷列表。
func (s *MilestonesService) ListMilestoneIssues(ctx context.Context, pid interface{}, milestoneID int, opts *ListMilestoneIssuesOptions) ([]*MilestoneIssue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/milestones/%d/issues", project, milestoneID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var issues []*MilestoneIssue
	resp, err := s.client.Do(req, &issues)
	if err != nil {
		return nil, resp, err
	}

	return issues, resp, nil
}
