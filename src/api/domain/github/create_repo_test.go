package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestTest(t *testing.T) {

	newRequest := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a golang intro repository",
		HomePage:    "https://github.com",
		Private:     false,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}

	bytes, err := json.Marshal(newRequest)

	fmt.Println(string(bytes))
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t, `{"name":"golang introduction","description":"a golang intro repository","homepage":"https://github.com","private":false,"has_issues":false,"has_projects":false,"has_wiki":false}`, string(bytes))

	err = json.Unmarshal(bytes, &newRequest)
	assert.Nil(t, err)
	assert.NotNil(t, newRequest)
	assert.EqualValues(t, newRequest.Name, "golang introduction")
}
