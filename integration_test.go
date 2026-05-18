package gongfeng

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

// integrationProject 是用于真实环境测试的项目路径。
const integrationProject = "studyzy/gongfeng-sdk-go"

// newIntegrationClient 读取环境变量 GONGFENG_TOKEN，若未设置则跳过测试。
// 返回一个连接到 DefaultBaseURL 的真实 Client。
func newIntegrationClient(t *testing.T) *Client {
	t.Helper()
	token := strings.TrimSpace(os.Getenv("GONGFENG_TOKEN"))
	if token == "" {
		t.Skip("GONGFENG_TOKEN not set, skipping integration test")
	}
	c, err := NewClient(token)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	return c
}

// uniqueName 生成带纳秒时间戳后缀的资源名，避免多次运行命名冲突。
func uniqueName(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}

// integrationCtx 返回带 60s 超时的上下文。
func integrationCtx(t *testing.T) (context.Context, context.CancelFunc) {
	t.Helper()
	return context.WithTimeout(context.Background(), 60*time.Second)
}

// resolveDefaultBranch 获取目标项目的默认分支名（兜底 master）。
func resolveDefaultBranch(t *testing.T, c *Client) string {
	t.Helper()
	ctx, cancel := integrationCtx(t)
	defer cancel()
	p, _, err := c.Projects.GetProject(ctx, integrationProject)
	if err != nil {
		t.Fatalf("GetProject: %v", err)
	}
	if p.DefaultBranch != "" {
		return p.DefaultBranch
	}
	return "master"
}

// --- 只读测试 ---

func TestIntegration_CurrentUser(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	user, _, err := c.Users.GetCurrentUser(ctx)
	if err != nil {
		t.Fatalf("GetCurrentUser: %v", err)
	}
	if user.ID == 0 || user.Username == "" {
		t.Fatalf("unexpected user: %+v", user)
	}
	t.Logf("current user: id=%d username=%s", user.ID, user.Username)
}

func TestIntegration_GetProject(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	p, _, err := c.Projects.GetProject(ctx, integrationProject)
	if err != nil {
		t.Fatalf("GetProject: %v", err)
	}
	if p.ID == 0 {
		t.Fatalf("unexpected project: %+v", p)
	}
	if p.PathWithNamespace != "" && p.PathWithNamespace != integrationProject {
		t.Fatalf("path_with_namespace mismatch: got %q want %q", p.PathWithNamespace, integrationProject)
	}
	t.Logf("project id=%d default_branch=%s", p.ID, p.DefaultBranch)
}

func TestIntegration_ListBranches(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	branches, _, err := c.Branches.ListBranches(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListBranches: %v", err)
	}
	if len(branches) == 0 {
		t.Fatal("expected at least one branch")
	}
	t.Logf("got %d branches, first: %s", len(branches), branches[0].Name)
}

func TestIntegration_GetDefaultBranch(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	branchName := resolveDefaultBranch(t, c)
	b, _, err := c.Branches.GetBranch(ctx, integrationProject, branchName)
	if err != nil {
		t.Fatalf("GetBranch(%s): %v", branchName, err)
	}
	if b.Name != branchName {
		t.Fatalf("branch name mismatch: got %q want %q", b.Name, branchName)
	}
}

func TestIntegration_ListTags(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	tags, _, err := c.Tags.ListTags(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListTags: %v", err)
	}
	t.Logf("got %d tags", len(tags))
}

func TestIntegration_ListCommits(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	commits, _, err := c.Commits.ListCommits(ctx, integrationProject, &ListCommitsOptions{
		ListOptions: ListOptions{PerPage: 5},
	})
	if err != nil {
		t.Fatalf("ListCommits: %v", err)
	}
	if len(commits) == 0 {
		t.Fatal("expected at least one commit")
	}
	t.Logf("got %d commits, latest: %s", len(commits), commits[0].ID)
}

func TestIntegration_GetRepositoryTree(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	nodes, _, err := c.Repositories.ListTree(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListTree: %v", err)
	}
	if len(nodes) == 0 {
		t.Fatal("expected non-empty tree")
	}
	t.Logf("got %d tree nodes", len(nodes))
}

func TestIntegration_ListNamespaces(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	nss, _, err := c.Namespaces.ListNamespaces(ctx, nil)
	if err != nil {
		t.Fatalf("ListNamespaces: %v", err)
	}
	t.Logf("got %d namespaces", len(nss))
}

func TestIntegration_ListProjectMembers(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	members, _, err := c.Projects.ListProjectMembers(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListProjectMembers: %v", err)
	}
	t.Logf("got %d members", len(members))
}

func TestIntegration_ListProjectEvents(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	events, _, err := c.Projects.ListProjectEvents(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListProjectEvents: %v", err)
	}
	t.Logf("got %d events", len(events))
}

func TestIntegration_ListMergeRequests(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	mrs, _, err := c.MergeRequests.ListMergeRequests(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListMergeRequests: %v", err)
	}
	t.Logf("got %d merge requests", len(mrs))
}

func TestIntegration_ListIssues(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	issues, _, err := c.Issues.ListIssues(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListIssues: %v", err)
	}
	t.Logf("got %d issues", len(issues))
}

func TestIntegration_ListLabels(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	labels, _, err := c.Labels.ListLabels(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListLabels: %v", err)
	}
	t.Logf("got %d labels", len(labels))
}

func TestIntegration_ListWebhooks(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	hooks, _, err := c.Webhooks.ListWebhooks(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListWebhooks: %v", err)
	}
	t.Logf("got %d webhooks", len(hooks))
}

func TestIntegration_ListReleases(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	releases, _, err := c.Releases.ListReleases(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListReleases: %v", err)
	}
	t.Logf("got %d releases", len(releases))
}

// --- 可逆写操作（创建后立即注册清理） ---

func TestIntegration_BranchLifecycle(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	defaultBranch := resolveDefaultBranch(t, c)
	branchName := uniqueName("sdktest-branch")

	created, _, err := c.Branches.CreateBranch(ctx, integrationProject, &CreateBranchOptions{
		BranchName: Ptr(branchName),
		Ref:        Ptr(defaultBranch),
	})
	if err != nil {
		t.Fatalf("CreateBranch(%s from %s): %v", branchName, defaultBranch, err)
	}
	if created.Name != branchName {
		t.Fatalf("created branch name mismatch: got %q want %q", created.Name, branchName)
	}
	// 立刻注册清理，避免后续断言失败时残留。
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if _, err := c.Branches.DeleteBranch(cleanupCtx, integrationProject, branchName); err != nil {
			t.Logf("cleanup DeleteBranch(%s): %v", branchName, err)
		}
	})

	got, _, err := c.Branches.GetBranch(ctx, integrationProject, branchName)
	if err != nil {
		t.Fatalf("GetBranch(%s): %v", branchName, err)
	}
	if got.Name != branchName {
		t.Fatalf("GetBranch name mismatch: got %q want %q", got.Name, branchName)
	}
}

func TestIntegration_TagLifecycle(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	defaultBranch := resolveDefaultBranch(t, c)
	tagName := uniqueName("sdktest-tag")

	created, _, err := c.Tags.CreateTag(ctx, integrationProject, &CreateTagOptions{
		TagName: Ptr(tagName),
		Ref:     Ptr(defaultBranch),
		Message: Ptr("sdk integration test tag"),
	})
	if err != nil {
		t.Fatalf("CreateTag(%s from %s): %v", tagName, defaultBranch, err)
	}
	if created.Name != tagName {
		t.Fatalf("created tag name mismatch: got %q want %q", created.Name, tagName)
	}
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if _, err := c.Tags.DeleteTag(cleanupCtx, integrationProject, tagName); err != nil {
			t.Logf("cleanup DeleteTag(%s): %v", tagName, err)
		}
	})

	got, _, err := c.Tags.GetTag(ctx, integrationProject, tagName)
	if err != nil {
		t.Fatalf("GetTag(%s): %v", tagName, err)
	}
	if got.Name != tagName {
		t.Fatalf("GetTag name mismatch: got %q want %q", got.Name, tagName)
	}
}

func TestIntegration_LabelLifecycle(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	labelName := uniqueName("sdktest-label")

	created, _, err := c.Labels.CreateLabel(ctx, integrationProject, &CreateLabelOptions{
		Name:        Ptr(labelName),
		Color:       Ptr("#FF0000"),
		Description: Ptr("created by gongfeng-sdk-go integration test"),
	})
	if err != nil {
		t.Fatalf("CreateLabel(%s): %v", labelName, err)
	}
	if created.Name != labelName {
		t.Fatalf("created label name mismatch: got %q want %q", created.Name, labelName)
	}
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if _, err := c.Labels.DeleteLabel(cleanupCtx, integrationProject, &DeleteLabelOptions{
			Name: Ptr(labelName),
		}); err != nil {
			t.Logf("cleanup DeleteLabel(%s): %v", labelName, err)
		}
	})

	labels, _, err := c.Labels.ListLabels(ctx, integrationProject, nil)
	if err != nil {
		t.Fatalf("ListLabels: %v", err)
	}
	var found bool
	for _, l := range labels {
		if l.Name == labelName {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("created label %s not found in list", labelName)
	}
}

func TestIntegration_WebhookLifecycle(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	hookURL := fmt.Sprintf("https://example.com/%s", uniqueName("sdktest-hook"))

	created, _, err := c.Webhooks.AddWebhook(ctx, integrationProject, &AddWebhookOptions{
		URL:        Ptr(hookURL),
		PushEvents: Ptr(true),
	})
	if err != nil {
		t.Fatalf("AddWebhook(%s): %v", hookURL, err)
	}
	if created.ID == 0 {
		t.Fatalf("created webhook missing id: %+v", created)
	}
	hookID := created.ID
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if _, err := c.Webhooks.DeleteWebhook(cleanupCtx, integrationProject, hookID); err != nil {
			t.Logf("cleanup DeleteWebhook(%d): %v", hookID, err)
		}
	})

	got, _, err := c.Webhooks.GetWebhook(ctx, integrationProject, hookID)
	if err != nil {
		t.Fatalf("GetWebhook(%d): %v", hookID, err)
	}
	if got.ID != hookID {
		t.Fatalf("GetWebhook id mismatch: got %d want %d", got.ID, hookID)
	}
	if got.URL != hookURL {
		t.Fatalf("GetWebhook url mismatch: got %q want %q", got.URL, hookURL)
	}
}

func TestIntegration_IssueLifecycle(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	title := uniqueName("sdktest-issue")

	created, _, err := c.Issues.CreateIssue(ctx, integrationProject, &CreateIssueOptions{
		Title:       Ptr(title),
		Description: Ptr("created by gongfeng-sdk-go integration test"),
	})
	if err != nil {
		t.Fatalf("CreateIssue(%s): %v", title, err)
	}
	if created.ID == 0 {
		t.Fatalf("created issue missing id: %+v", created)
	}

	// 工蜂 GET/PUT /projects/:pid/issues/:id 走全局主键 id，而非项目内 iid。
	// 工蜂 API 不支持删除 Issue，cleanup 只能关闭。
	issueID := created.ID
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if _, _, err := c.Issues.UpdateIssue(cleanupCtx, integrationProject, issueID, &UpdateIssueOptions{
			StateEvent: Ptr("close"),
		}); err != nil {
			t.Logf("cleanup UpdateIssue close(%d): %v", issueID, err)
		}
	})

	got, _, err := c.Issues.GetIssue(ctx, integrationProject, issueID)
	if err != nil {
		t.Fatalf("GetIssue(%d): %v", issueID, err)
	}
	if got.Title != title {
		t.Fatalf("GetIssue title mismatch: got %q want %q", got.Title, title)
	}
}

func TestIntegration_StarLifecycle(t *testing.T) {
	c := newIntegrationClient(t)
	ctx, cancel := integrationCtx(t)
	defer cancel()

	// 记录初始状态，结束时尽量恢复。
	initial, _, err := c.Projects.GetStarStatus(ctx, integrationProject)
	if err != nil {
		t.Fatalf("GetStarStatus: %v", err)
	}

	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		current, _, err := c.Projects.GetStarStatus(cleanupCtx, integrationProject)
		if err != nil {
			t.Logf("cleanup GetStarStatus: %v", err)
			return
		}
		if current == initial {
			return
		}
		if initial {
			// 之前是 starred，现在不是 → 重新 star。
			if _, _, err := c.Projects.StarProject(cleanupCtx, integrationProject); err != nil {
				t.Logf("cleanup StarProject: %v", err)
			}
		} else {
			// 之前未 star，现在 star 了 → unstar。
			if _, err := c.Projects.UnstarProject(cleanupCtx, integrationProject); err != nil {
				t.Logf("cleanup UnstarProject: %v", err)
			}
		}
	})

	if initial {
		// 已 starred：先 unstar，再 star。
		if _, err := c.Projects.UnstarProject(ctx, integrationProject); err != nil {
			t.Fatalf("UnstarProject: %v", err)
		}
		status, _, err := c.Projects.GetStarStatus(ctx, integrationProject)
		if err != nil {
			t.Fatalf("GetStarStatus after unstar: %v", err)
		}
		if status {
			t.Fatal("expected star status false after UnstarProject")
		}
		if _, _, err := c.Projects.StarProject(ctx, integrationProject); err != nil {
			t.Fatalf("StarProject: %v", err)
		}
	} else {
		// 未 starred：先 star，再 unstar。
		if _, _, err := c.Projects.StarProject(ctx, integrationProject); err != nil {
			t.Fatalf("StarProject: %v", err)
		}
		status, _, err := c.Projects.GetStarStatus(ctx, integrationProject)
		if err != nil {
			t.Fatalf("GetStarStatus after star: %v", err)
		}
		if !status {
			t.Fatal("expected star status true after StarProject")
		}
		if _, err := c.Projects.UnstarProject(ctx, integrationProject); err != nil {
			t.Fatalf("UnstarProject: %v", err)
		}
	}
}
