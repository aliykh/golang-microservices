package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HomePage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	FullName string    `json:"full_name"`
	Owner    RepoOwner `json:"owner"`
}

/*
  "owner": {
    "login": "aliykh",
    "id": 22681649,
    "node_id": "MDQ6VXNlcjIyNjgxNjQ5",
    "avatar_url": "https://avatars2.githubusercontent.com/u/22681649?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/aliykh",
    "html_url": "https://github.com/aliykh",
    "followers_url": "https://api.github.com/users/aliykh/followers",
    "following_url": "https://api.github.com/users/aliykh/following{/other_user}",
    "gists_url": "https://api.github.com/users/aliykh/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/aliykh/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/aliykh/subscriptions",
    "organizations_url": "https://api.github.com/users/aliykh/orgs",
    "repos_url": "https://api.github.com/users/aliykh/repos",
    "events_url": "https://api.github.com/users/aliykh/events{/privacy}",
    "received_events_url": "https://api.github.com/users/aliykh/received_events",
    "type": "User",
    "site_admin": false
  },
*/

type RepoOwner struct {
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
	Id      int64  `json:"id"`
}
