package metadata

import (
	"fmt"

	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	"github.com/NII-DG/gogs/internal/urlutil"
)

func SearchRepo(c *context.APIContext) {

	// Getting repository owner information from DB
	ownerName := c.Params(":ownername")
	owner, err := db.GetUserByName(ownerName)
	if err != nil {
		c.NotFoundOrError(err, "get user by name")
		return
	}

	// Getting repository information from DB
	repoName := c.Params(":reponame")
	repo, err := db.GetRepositoryByName(owner.ID, repoName)
	if err != nil {
		c.NotFoundOrError(err, "get repo by owner name and repository name")
		return
	}

	// check request user access repository information
	users, err := repo.GetAssignees()
	if err != nil {
		c.Error(err, "failure connect to DB")
	}

	request_user_id := c.User.ID

	accessRight := false
	for _, u := range users {
		if u.ID == request_user_id {
			accessRight = true
		}
	}

	if !accessRight {
		c.Error(fmt.Errorf("your not accessRight"), "failure connect to DB")
	}

	// Creating Repository Metadata
	path := repo.Owner.Name + "/" + repoName
	url, err := urlutil.UpdatePath(c.BaseURL, path)
	if err != nil {
		c.Errorf(err, "%v", err)
	}
	repoMatadata := RepositoryMetadata{
		Name:        repo.Name,
		Description: repo.Description,
		Url:         url,
	}
	c.JSONSuccess(repoMatadata)
}
