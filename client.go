package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	gitlabclient "github.com/xanzy/go-gitlab"
)

const (
	DefaultBaseURL    = "https://code.tencent.com/"
	DefaultAPIVersion = "v3"
)

type Options struct {
	// BaseURL is the Gongfeng host URL, for example: https://code.tencent.com/
	BaseURL    string
	// APIVersion is the API version path suffix, for example: v3 or v4.
	APIVersion string
	HTTPClient *http.Client
}

type Client struct {
	*gitlabclient.Client
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

	apiURL, err := buildAPIURL(baseURL)
	if err != nil {
		return nil, err
	}

	httpClient = wrapHTTPClientForVersion(httpClient, apiVersion)

	clientOpts := []gitlabclient.ClientOptionFunc{gitlabclient.WithBaseURL(apiURL)}
	if httpClient != nil {
		clientOpts = append(clientOpts, gitlabclient.WithHTTPClient(httpClient))
	}

	cli, err := gitlabclient.NewClient(token, clientOpts...)
if err != nil {
return nil, err
}

return &Client{Client: cli}, nil
}

// Call sends a REST API request to endpoint (relative to /api/{version}/),
// optionally JSON-encodes payload as the request body, and decodes the response
// into out.
func (c *Client) Call(ctx context.Context, method, endpoint string, payload any, out any, requestOptions ...gitlabclient.RequestOptionFunc) (*gitlabclient.Response, error) {
if c == nil || c.Client == nil {
return nil, fmt.Errorf("client is nil")
}

requestPath := strings.TrimPrefix(endpoint, "/")
	req, err := c.NewRequest(method, requestPath, payload, requestOptions)
if err != nil {
return nil, err
}

if ctx != nil {
req = req.WithContext(ctx)
}

return c.Do(req, out)
}

func buildAPIURL(baseURL string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base url: %w", err)
	}

	basePath := strings.TrimSuffix(u.Path, "/")
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
