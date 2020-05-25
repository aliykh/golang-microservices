package repositories


type CreateRepoRequest struct {
	Name        string `form:"name" xml:"name" json:"name" binding:"required"`
	Description string `form:"description" xml:"description" json:"description" binding:"required"`
}

type CreateRepoResponse struct {
	Id int64 `json:"id"`
	Owner string `json:"owner"`
	Name string `json:"name"`
}