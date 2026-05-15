package gongfeng

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// TreeNode 表示仓库文件树中的一个节点。
type TreeNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Mode string `json:"mode"`
}

// RepositoryFile 表示仓库中的一个文件。
type RepositoryFile struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	Size     int    `json:"size"`
	Encoding string `json:"encoding"`
	Content  string `json:"content"`
	Ref      string `json:"ref"`
	BlobID   string `json:"blob_id"`
	CommitID string `json:"commit_id"`
}

// CompareResult 表示两个分支/Tag/SHA 的比较结果。
type CompareResult struct {
	Commit         *Commit   `json:"commit"`
	Commits        []*Commit `json:"commits"`
	Diffs          []*Diff   `json:"diffs"`
	CompareTimeout bool      `json:"compare_timeout"`
	CompareSameRef bool      `json:"compare_same_ref"`
}

// RepositoriesService 处理与仓库相关的 API 调用。
type RepositoriesService struct {
	client *Client
}

// ArchiveOptions 是 Archive 的可选参数。
type ArchiveOptions struct {
	SHA *string `url:"sha,omitempty" json:"sha,omitempty"`
}

// Archive 下载项目仓库的存档，将内容写入 w。
func (s *RepositoriesService) Archive(ctx context.Context, pid interface{}, w io.Writer, opts *ArchiveOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/archive", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, w)
}

// ListTreeOptions 是 ListTree 的可选参数。
type ListTreeOptions struct {
	ListOptions
	Path    *string `url:"path,omitempty" json:"path,omitempty"`
	RefName *string `url:"ref_name,omitempty" json:"ref_name,omitempty"`
}

// ListTree 获取项目仓库的文件树。
func (s *RepositoriesService) ListTree(ctx context.Context, pid interface{}, opts *ListTreeOptions) ([]*TreeNode, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/tree", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var nodes []*TreeNode
	resp, err := s.client.Do(req, &nodes)
	if err != nil {
		return nil, resp, err
	}

	return nodes, resp, nil
}

// GetFileOptions 是 GetFile 的可选参数。
type GetFileOptions struct {
	FilePath *string `url:"file_path,omitempty" json:"file_path,omitempty"`
	Ref      *string `url:"ref,omitempty" json:"ref,omitempty"`
}

// GetFile 获取仓库中指定文件的内容。
func (s *RepositoriesService) GetFile(ctx context.Context, pid interface{}, opts *GetFileOptions) (*RepositoryFile, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/files", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var file RepositoryFile
	resp, err := s.client.Do(req, &file)
	if err != nil {
		return nil, resp, err
	}

	return &file, resp, nil
}

// CreateFileOptions 是 CreateFile 的可选参数。
type CreateFileOptions struct {
	FilePath      *string `json:"file_path,omitempty" url:"file_path,omitempty"`
	BranchName    *string `json:"branch_name,omitempty" url:"branch_name,omitempty"`
	Content       *string `json:"content,omitempty" url:"content,omitempty"`
	CommitMessage *string `json:"commit_message,omitempty" url:"commit_message,omitempty"`
}

// CreateFile 在仓库中创建一个新文件。
func (s *RepositoriesService) CreateFile(ctx context.Context, pid interface{}, opts *CreateFileOptions) (*RepositoryFile, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/files", project)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var file RepositoryFile
	resp, err := s.client.Do(req, &file)
	if err != nil {
		return nil, resp, err
	}

	return &file, resp, nil
}

// UpdateFileOptions 是 UpdateFile 的可选参数。
type UpdateFileOptions struct {
	FilePath      *string `json:"file_path,omitempty" url:"file_path,omitempty"`
	BranchName    *string `json:"branch_name,omitempty" url:"branch_name,omitempty"`
	Content       *string `json:"content,omitempty" url:"content,omitempty"`
	CommitMessage *string `json:"commit_message,omitempty" url:"commit_message,omitempty"`
}

// UpdateFile 修改仓库中的一个文件。
func (s *RepositoriesService) UpdateFile(ctx context.Context, pid interface{}, opts *UpdateFileOptions) (*RepositoryFile, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/files", project)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var file RepositoryFile
	resp, err := s.client.Do(req, &file)
	if err != nil {
		return nil, resp, err
	}

	return &file, resp, nil
}

// DeleteFileOptions 是 DeleteFile 的可选参数。
type DeleteFileOptions struct {
	FilePath      *string `json:"file_path,omitempty" url:"file_path,omitempty"`
	BranchName    *string `json:"branch_name,omitempty" url:"branch_name,omitempty"`
	CommitMessage *string `json:"commit_message,omitempty" url:"commit_message,omitempty"`
}

// DeleteFile 删除仓库中的一个文件。
func (s *RepositoriesService) DeleteFile(ctx context.Context, pid interface{}, opts *DeleteFileOptions) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/files", project)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, opts)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// CompareOptions 是 Compare 的可选参数。
type CompareOptions struct {
	From *string `url:"from,omitempty" json:"from,omitempty"`
	To   *string `url:"to,omitempty" json:"to,omitempty"`
}

// Compare 比较项目中两个分支、Tag 或 SHA。
func (s *RepositoriesService) Compare(ctx context.Context, pid interface{}, opts *CompareOptions) (*CompareResult, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("projects/%s/repository/compare", project)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, opts)
	if err != nil {
		return nil, nil, err
	}

	var result CompareResult
	resp, err := s.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
