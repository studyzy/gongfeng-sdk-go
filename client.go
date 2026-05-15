package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

gitlab "github.com/xanzy/go-gitlab"
)

const (
	DefaultBaseURL    = "https://code.tencent.com/"
	DefaultAPIVersion = "v3"
)

type Options struct {
BaseURL    string
APIVersion string
HTTPClient *http.Client
}

type Client struct {
*gitlab.Client
}

func NewClient(token string, opts *Options) (*Client, error) {
if strings.TrimSpace(token) == "" {
return nil, fmt.Errorf("token is required")
}

baseURL := DefaultBaseURL
apiVersion := DefaultAPIVersion
var httpClient *http.Client
if opts != nil {
if opts.BaseURL != "" {
baseURL = opts.BaseURL
}
if opts.APIVersion != "" {
apiVersion = opts.APIVersion
}
httpClient = opts.HTTPClient
}

	apiURL, err := buildAPIURL(baseURL, apiVersion)
	if err != nil {
		return nil, err
	}

	httpClient = wrapHTTPClientForVersion(httpClient, apiVersion)

	clientOpts := []gitlab.ClientOptionFunc{gitlab.WithBaseURL(apiURL)}
	if httpClient != nil {
		clientOpts = append(clientOpts, gitlab.WithHTTPClient(httpClient))
	}

cli, err := gitlab.NewClient(token, clientOpts...)
if err != nil {
return nil, err
}

return &Client{Client: cli}, nil
}

func (c *Client) Call(ctx context.Context, method, endpoint string, payload any, out any, reqOpts ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
if c == nil || c.Client == nil {
return nil, fmt.Errorf("client is nil")
}

requestPath := strings.TrimPrefix(endpoint, "/")
	req, err := c.NewRequest(method, requestPath, payload, reqOpts)
if err != nil {
return nil, err
}

if ctx != nil {
req = req.WithContext(ctx)
}

return c.Do(req, out)
}

func buildAPIURL(baseURL, apiVersion string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base url: %w", err)
	}

	_ = apiVersion
	basePath := path.Join(strings.TrimSuffix(u.Path, "/"))
	if basePath == "." {
		basePath = ""
	}
	if basePath == "" {
		u.Path = "/"
	} else {
		u.Path = basePath + "/"
	}
	return u.String(), nil
}

func wrapHTTPClientForVersion(client *http.Client, apiVersion string) *http.Client {
	if apiVersion == "" || apiVersion == "v4" {
		return client
	}

	baseClient := client
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	clone := *baseClient
	transport := clone.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	clone.Transport = &versionRewriteTransport{
		base:       transport,
		apiVersion: strings.TrimPrefix(apiVersion, "/"),
	}

	return &clone
}

type versionRewriteTransport struct {
	base       http.RoundTripper
	apiVersion string
}

func (t *versionRewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	replaced := req.Clone(req.Context())
	replaced.URL = cloneURL(req.URL)
	replaced.URL.Path = strings.Replace(replaced.URL.Path, "/api/v4/", "/api/"+t.apiVersion+"/", 1)
	return t.base.RoundTrip(replaced)
}

func cloneURL(u *url.URL) *url.URL {
	clone := *u
	return &clone
}
