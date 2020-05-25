package github

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	Field    string `json:"field"`
}

type GithubErrorResponse struct {
	Message          string        `json:"message"`
	DocumentationUrl string        `json:"documentation_url"`
	Errors           []GithubError `json:"errors"`
	StatusCode       int           `json:"status_code"`
}
