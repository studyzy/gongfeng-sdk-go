package gongfeng

import (
	"fmt"
	"time"
)

// timeLayout 是工蜂 API 返回的时间格式。
const timeLayout = "2006-01-02T15:04:05+08:00"

// Time 是工蜂 API 中使用的时间类型。
// 它封装了 time.Time 并提供了自定义的 JSON 序列化/反序列化，
// 以正确处理工蜂 API 返回的多种时间格式。
type Time struct {
	time.Time
}

// MarshalJSON 实现 json.Marshaler 接口。
func (t Time) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Time.Format(timeLayout) + `"`), nil
}

// UnmarshalJSON 实现 json.Unmarshaler 接口。
// 支持解析多种时间格式。
func (t *Time) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" || str == `""` {
		return nil
	}

	// 去除引号
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	// 尝试多种时间格式
	layouts := []string{
		"2006-01-02T15:04:05+08:00",
		"2006-01-02T15:04:05+0000",
		"2006-01-02T15:04:05-0700",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.000+08:00",
		time.RFC3339,
		"2006-01-02",
	}

	var parseErr error
	for _, layout := range layouts {
		parsed, err := time.Parse(layout, str)
		if err == nil {
			t.Time = parsed
			return nil
		}
		parseErr = err
	}

	return fmt.Errorf("gongfeng: cannot parse time %q: %w", string(data), parseErr)
}

// String 实现 fmt.Stringer 接口。
func (t Time) String() string {
	return t.Time.Format(timeLayout)
}

// ListOptions 是所有列表接口通用的分页参数。
type ListOptions struct {
	Page    int `url:"page,omitempty" json:"page,omitempty"`
	PerPage int `url:"per_page,omitempty" json:"per_page,omitempty"`
}

// AccessLevelValue 表示工蜂的权限级别。
type AccessLevelValue int

// 工蜂权限级别常量。
const (
	GuestPermission     AccessLevelValue = 10
	FollowerPermission  AccessLevelValue = 15
	ReporterPermission  AccessLevelValue = 20
	DeveloperPermission AccessLevelValue = 30
	MasterPermission    AccessLevelValue = 40
	OwnerPermission     AccessLevelValue = 50
)

// VisibilityValue 表示项目的可见级别。
type VisibilityValue int

// 项目可见级别常量。
const (
	PrivateVisibility  VisibilityValue = 0
	InternalVisibility VisibilityValue = 10
	PublicVisibility   VisibilityValue = 20
)

// Namespace 表示工蜂命名空间。
type Namespace struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Path        string `json:"path,omitempty"`
	Kind        string `json:"kind,omitempty"`
	WebURL      string `json:"web_url,omitempty"`
	Description string `json:"description,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
}

// User 表示工蜂用户信息。
type User struct {
	ID        int    `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Name      string `json:"name,omitempty"`
	State     string `json:"state,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	WebURL    string `json:"web_url,omitempty"`
}

// UserDetail 表示工蜂用户的详细信息。
type UserDetail struct {
	User
	Email            string `json:"email,omitempty"`
	Bio              string `json:"bio,omitempty"`
	CreatedAt        Time   `json:"created_at,omitempty"`
	IsAdmin          bool   `json:"is_admin,omitempty"`
	ProjectLimit     int    `json:"projects_limit,omitempty"`
	CanCreateGroup   bool   `json:"can_create_group,omitempty"`
	CanCreateProject bool   `json:"can_create_project,omitempty"`
}

// Member 表示一个组或项目的成员。
type Member struct {
	ID          int              `json:"id,omitempty"`
	Username    string           `json:"username,omitempty"`
	Name        string           `json:"name,omitempty"`
	State       string           `json:"state,omitempty"`
	AvatarURL   string           `json:"avatar_url,omitempty"`
	WebURL      string           `json:"web_url,omitempty"`
	AccessLevel AccessLevelValue `json:"access_level,omitempty"`
}

// Commit 表示一次 Git 提交。
type Commit struct {
	ID             string   `json:"id,omitempty"`
	ShortID        string   `json:"short_id,omitempty"`
	Title          string   `json:"title,omitempty"`
	Message        string   `json:"message,omitempty"`
	AuthorName     string   `json:"author_name,omitempty"`
	AuthorEmail    string   `json:"author_email,omitempty"`
	AuthoredDate   Time     `json:"authored_date,omitempty"`
	CommitterName  string   `json:"committer_name,omitempty"`
	CommitterEmail string   `json:"committer_email,omitempty"`
	CommittedDate  Time     `json:"committed_date,omitempty"`
	CreatedAt      Time     `json:"created_at,omitempty"`
	ParentIDs      []string `json:"parent_ids,omitempty"`
}

// Diff 表示一个文件的变更。
type Diff struct {
	OldPath     string `json:"old_path,omitempty"`
	NewPath     string `json:"new_path,omitempty"`
	AMode       int    `json:"a_mode,omitempty"`
	BMode       int    `json:"b_mode,omitempty"`
	Diff        string `json:"diff,omitempty"`
	NewFile     bool   `json:"new_file,omitempty"`
	RenamedFile bool   `json:"renamed_file,omitempty"`
	DeletedFile bool   `json:"deleted_file,omitempty"`
	IsTooLarge  bool   `json:"is_too_large,omitempty"`
	IsCollapse  bool   `json:"is_collapse,omitempty"`
	Additions   int    `json:"additions,omitempty"`
	Deletions   int    `json:"deletions,omitempty"`
}

// ConfigStorage 表示项目的存储配置。
type ConfigStorage struct {
	LimitLFSFileSize float64 `json:"limit_lfs_file_size,omitempty"`
	LimitSize        float64 `json:"limit_size,omitempty"`
	LimitFileSize    float64 `json:"limit_file_size,omitempty"`
	LimitLFSSize     float64 `json:"limit_lfs_size,omitempty"`
}

// Statistics 表示项目的统计信息。
type Statistics struct {
	CommitCount    int     `json:"commit_count,omitempty"`
	RepositorySize float64 `json:"repository_size,omitempty"`
}

// Ptr 返回 v 的指针，用于构造可选参数。
func Ptr[T any](v T) *T {
	return &v
}
