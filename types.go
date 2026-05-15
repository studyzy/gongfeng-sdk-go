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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Kind        string `json:"kind"`
	WebURL      string `json:"web_url"`
	Description string `json:"description"`
	AvatarURL   string `json:"avatar_url"`
}

// User 表示工蜂用户信息。
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

// UserDetail 表示工蜂用户的详细信息。
type UserDetail struct {
	User
	Email            string `json:"email"`
	Bio              string `json:"bio"`
	CreatedAt        Time   `json:"created_at"`
	IsAdmin          bool   `json:"is_admin"`
	ProjectLimit     int    `json:"projects_limit"`
	CanCreateGroup   bool   `json:"can_create_group"`
	CanCreateProject bool   `json:"can_create_project"`
}

// Member 表示一个组或项目的成员。
type Member struct {
	ID          int              `json:"id"`
	Username    string           `json:"username"`
	Name        string           `json:"name"`
	State       string           `json:"state"`
	AvatarURL   string           `json:"avatar_url"`
	WebURL      string           `json:"web_url"`
	AccessLevel AccessLevelValue `json:"access_level"`
}

// Commit 表示一次 Git 提交。
type Commit struct {
	ID             string   `json:"id"`
	ShortID        string   `json:"short_id"`
	Title          string   `json:"title"`
	Message        string   `json:"message"`
	AuthorName     string   `json:"author_name"`
	AuthorEmail    string   `json:"author_email"`
	AuthoredDate   Time     `json:"authored_date"`
	CommitterName  string   `json:"committer_name"`
	CommitterEmail string   `json:"committer_email"`
	CommittedDate  Time     `json:"committed_date"`
	CreatedAt      Time     `json:"created_at"`
	ParentIDs      []string `json:"parent_ids"`
}

// Diff 表示一个文件的变更。
type Diff struct {
	OldPath     string `json:"old_path"`
	NewPath     string `json:"new_path"`
	AMode       int    `json:"a_mode"`
	BMode       int    `json:"b_mode"`
	Diff        string `json:"diff"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
	IsTooLarge  bool   `json:"is_too_large"`
	IsCollapse  bool   `json:"is_collapse"`
	Additions   int    `json:"additions"`
	Deletions   int    `json:"deletions"`
}

// ConfigStorage 表示项目的存储配置。
type ConfigStorage struct {
	LimitLFSFileSize int `json:"limit_lfs_file_size"`
	LimitSize        int `json:"limit_size"`
	LimitFileSize    int `json:"limit_file_size"`
	LimitLFSSize     int `json:"limit_lfs_size"`
}

// Statistics 表示项目的统计信息。
type Statistics struct {
	CommitCount    int     `json:"commit_count"`
	RepositorySize float64 `json:"repository_size"`
}

// Ptr 返回 v 的指针，用于构造可选参数。
func Ptr[T any](v T) *T {
	return &v
}
