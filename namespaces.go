package gongfeng

import (
	"context"
	"net/http"
)

// NamespacesService 处理与命名空间相关的 API 调用。
type NamespacesService struct {
	client *Client
}

// ListNamespacesOptions 是 ListNamespaces 的可选参数。
type ListNamespacesOptions struct {
	ListOptions
	Search *string `url:"search,omitempty" json:"search,omitempty"`
}

// ListNamespaces 获取命名空间列表。
func (s *NamespacesService) ListNamespaces(ctx context.Context, opts *ListNamespacesOptions) ([]*Namespace, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "namespaces", opts)
	if err != nil {
		return nil, nil, err
	}

	var namespaces []*Namespace
	resp, err := s.client.Do(req, &namespaces)
	if err != nil {
		return nil, resp, err
	}

	return namespaces, resp, nil
}
