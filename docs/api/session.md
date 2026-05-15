# 会话


如果使用 NAMESPACE_PATH/PROJECT_PATH，确保NAMESPACE_PATH/PROJECT_PATH URL编码过，例子： /api/v3/projects/diaspora%2Fdiaspora (/编码 %2F)

### 获取私有令牌


```
POST /api/v3/session
```


**参数:**



| 参数         | 类型       | 描述        |
| ---------- | -------- | --------- |
| login      | string   | 登录的用户     |
| email      | string   | 用户的邮箱地址   |
| password   | string   | 用户的有效密码   |


### 您现在可以使用 工蜂Git 和 LDAP 凭证登录

**返回值：**


```json
{
  "id": 11323,
  "email": "git_user1@tencent.com",
  "name": "git_user1",
  "username": "git_user1",
  "web_url": "http://git.code.tencent.com/u/git_user1",
  "current_sign_in_at": "2018-07-26T01:55:51+0000",
  "is_admin": true,
  "projects_limit": 9999999,
  "skype": "",
  "linkedin": "",
  "twitter": "",
  "theme_id": 1,
  "bio": "",
  "state": "active",
  "color_scheme_id": 1,
  "website_url": "",
  "created_at": "2018-01-16T08:29:42+0000",
  "private_token": "MbvOEGNM7P5UEcvdNQG9"
}
```
