// Package gongfeng 提供腾讯工蜂 REST API 的 Go 客户端。
package gongfeng

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	// DefaultBaseURL 是工蜂公有云的默认地址。
	DefaultBaseURL = "https://git.code.tencent.com/"

	apiVersionPath = "api/v3"
	userAgent      = "gongfeng-sdk-go"
)

// Client 是腾讯工蜂 API 的客户端入口。
type Client struct {
	// baseURL 是 API 的基础 URL，包含 /api/v3/ 路径。
	baseURL *url.URL

	// token 是用于认证的 PRIVATE-TOKEN。
	token string

	// httpClient 是执行 HTTP 请求的客户端。
	httpClient *http.Client

	// userAgent 是请求中使用的 User-Agent 标头值。
	userAgent string

	// Services

	Branches       *BranchesService
	CommitStatuses *CommitStatusService
	Commits        *CommitsService
	Forks          *ForksService
	Groups         *GroupsService
	Issues         *IssuesService
	Labels         *LabelsService
	MergeRequests  *MergeRequestsService
	Milestones     *MilestonesService
	Namespaces     *NamespacesService
	Notes          *NotesService
	Projects       *ProjectsService
	Releases       *ReleasesService
	Repositories   *RepositoriesService
	Reviews        *ReviewsService
	Session        *SessionService
	Tags           *TagsService
	Users          *UsersService
	Watchers       *WatchersService
	Webhooks       *WebhooksService
}

// ClientOptionFunc 是客户端的函数式选项。
type ClientOptionFunc func(*Client) error

// WithBaseURL 设置工蜂实例的 URL。
func WithBaseURL(urlStr string) ClientOptionFunc {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}

// WithHTTPClient 设置自定义的 HTTP 客户端。
func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}

// NewClient 创建一个工蜂 API 客户端。token 为用户的 Private Token。
func NewClient(token string, options ...ClientOptionFunc) (*Client, error) {
	if strings.TrimSpace(token) == "" {
		return nil, fmt.Errorf("gongfeng: token is required")
	}

	c := &Client{
		token:     token,
		userAgent: userAgent,
	}

	// 设置默认 BaseURL
	if err := c.setBaseURL(DefaultBaseURL); err != nil {
		return nil, err
	}

	// 应用选项
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if c.httpClient == nil {
		c.httpClient = newDefaultHTTPClient()
	}

	// 初始化所有 Service
	c.Branches = &BranchesService{client: c}
	c.CommitStatuses = &CommitStatusService{client: c}
	c.Commits = &CommitsService{client: c}
	c.Forks = &ForksService{client: c}
	c.Groups = &GroupsService{client: c}
	c.Issues = &IssuesService{client: c}
	c.Labels = &LabelsService{client: c}
	c.MergeRequests = &MergeRequestsService{client: c}
	c.Milestones = &MilestonesService{client: c}
	c.Namespaces = &NamespacesService{client: c}
	c.Notes = &NotesService{client: c}
	c.Projects = &ProjectsService{client: c}
	c.Releases = &ReleasesService{client: c}
	c.Repositories = &RepositoriesService{client: c}
	c.Reviews = &ReviewsService{client: c}
	c.Session = &SessionService{client: c}
	c.Tags = &TagsService{client: c}
	c.Users = &UsersService{client: c}
	c.Watchers = &WatchersService{client: c}
	c.Webhooks = &WebhooksService{client: c}

	return c, nil
}

// newDefaultHTTPClient 创建默认 HTTP 客户端，启用兼容旧版 TLS 密码套件，
// 以支持部分工蜂实例使用的较旧 TLS 配置。
func newDefaultHTTPClient() *http.Client {
	var suites []uint16
	for _, s := range tls.CipherSuites() {
		suites = append(suites, s.ID)
	}
	for _, s := range tls.InsecureCipherSuites() {
		suites = append(suites, s.ID)
	}
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion:   tls.VersionTLS12,
				CipherSuites: suites,
			},
		},
	}
}

// setBaseURL 解析并设置基础 URL，自动拼接 /api/v3/ 路径。
func (c *Client) setBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("gongfeng: invalid base url %q: %w", urlStr, err)
	}

	// 拼接 api/v3/ 路径
	baseURL.Path += apiVersionPath + "/"
	c.baseURL = baseURL

	return nil
}

// Response 包装 http.Response，附加工蜂 API 的分页信息。
type Response struct {
	*http.Response

	// TotalItems 是查询结果的总条目数。
	TotalItems int

	// TotalPages 是总页数。
	TotalPages int

	// ItemsPerPage 是每页条目数。
	ItemsPerPage int

	// CurrentPage 是当前页码（从 1 开始）。
	CurrentPage int

	// NextPage 是下一页页码。
	NextPage int

	// PreviousPage 是上一页页码。
	PreviousPage int
}

// newResponse 从 http.Response 创建 Response，解析分页 Header。
func newResponse(r *http.Response) *Response {
	resp := &Response{Response: r}
	resp.TotalItems = parseIntHeader(r, "X-Total")
	resp.TotalPages = parseIntHeader(r, "X-Total-Pages")
	resp.ItemsPerPage = parseIntHeader(r, "X-Per-Page")
	resp.CurrentPage = parseIntHeader(r, "X-Page")
	resp.NextPage = parseIntHeader(r, "X-Next-Page")
	resp.PreviousPage = parseIntHeader(r, "X-Prev-Page")
	return resp
}

func parseIntHeader(r *http.Response, key string) int {
	v, _ := strconv.Atoi(r.Header.Get(key))
	return v
}

// ErrorResponse 表示工蜂 API 返回的错误。
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"message"`
}

// Error 实现 error 接口。
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s %s: %d %s",
		e.Response.Request.Method,
		e.Response.Request.URL,
		e.Response.StatusCode,
		e.Message,
	)
}

// NewRequest 构造一个 API 请求。
//
// path 是相对于 /api/v3/ 的路径（不带前导 /）。
// 对于 GET 请求，body 会被编码为 URL query 参数。
// 对于其他方法，body 会被 JSON 编码为请求体。
func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// 解析相对路径并拼接到 baseURL
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	reqHeaders := make(http.Header)
	reqHeaders.Set("PRIVATE-TOKEN", c.token)
	reqHeaders.Set("User-Agent", c.userAgent)

	var reqBody io.Reader
	switch {
	case method == http.MethodGet && body != nil:
		q, err := query.Values(body)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	case body != nil:
		reqHeaders.Set("Content-Type", "application/json")
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
		reqBody = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
	if err != nil {
		return nil, err
	}

	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

// Do 执行 API 请求并将 JSON 响应解码到 v。
// 如果 v 为 nil 则不解码响应体。
// 如果 v 实现了 io.Writer，则将响应体直接写入。
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	if err := checkResponse(resp); err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return response, err
}

// checkResponse 检查 HTTP 状态码，非 2xx 返回 ErrorResponse。
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		_ = json.Unmarshal(data, errorResponse)
		if errorResponse.Message == "" {
			errorResponse.Message = string(data)
		}
	}
	// 替换 body 以便后续读取
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	return errorResponse
}

// pathEscape 对路径参数进行 URL 编码。
func pathEscape(s string) string {
	return url.PathEscape(s)
}

// parseID 将项目/资源 ID 转换为 URL 路径段。
// 支持 int（直接转字符串）和 string（URL 编码，用于 namespace/project 格式）。
func parseID(id interface{}) (string, error) {
	switch v := id.(type) {
	case int:
		return strconv.Itoa(v), nil
	case string:
		return pathEscape(v), nil
	default:
		return "", fmt.Errorf("gongfeng: invalid ID type %T, expected int or string", id)
	}
}
