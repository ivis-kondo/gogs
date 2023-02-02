package metadata

import (
	"fmt"
	"net/http"

	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	"github.com/NII-DG/gogs/internal/urlutil"
)

func SearchRepo(c *context.APIContext) {

	// Getting repository owner information from DB
	ownerName := c.Params(":ownername")
	owner, err := db.GetUserByName(ownerName)
	if err != nil {
		c.NotFoundOrError(err, fmt.Sprintf("failure getting user by name from DB. User Name : %s", ownerName))
		return
	}

	// Getting repository information from DB
	repoName := c.Params(":reponame")
	repo, err := db.GetRepositoryByName(owner.ID, repoName)
	if err != nil {
		c.NotFoundOrError(err, fmt.Sprintf("failure getting repo by owner name and repository name from DB.  Repository : %s, Owner ID : %d", repoName, owner.ID))
		return
	}

	// check request user access repository information
	users, err := repo.GetAssignees()
	if err != nil {
		c.Error(err, fmt.Sprintf("failure getting read only user on repository from DB, Repository : %s", repo.Name))
	}

	request_user_id := c.User.ID

	accessRight := false
	for _, u := range users {
		if u.ID == request_user_id {
			accessRight = true
		}
	}

	if !accessRight {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"ok":    false,
			"error": "not access",
		})
		return
	}

	// Creating Repository Metadata
	path := repo.Owner.Name + "/" + repoName
	url, err := urlutil.UpdatePath(c.BaseURL, path)
	if err != nil {
		c.Errorf(err, "%v", err)
		return
	}
	repoMatadata := RepositoryMetadata{
		Name:        repo.Name,
		Description: repo.Description,
		Url:         url,
	}
	c.JSONSuccess(repoMatadata)
}
