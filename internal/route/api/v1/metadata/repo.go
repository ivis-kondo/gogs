package metadata

import (
	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
)

func SearchRepo(c *context.APIContext) {
	ownerName := c.Params(":ownername")
	owner, err := db.GetUserByName(ownerName)
	if err != nil {
		c.NotFoundOrError(err, "get user by name")
		return
	}
	repoName := c.Params(":reponame")
	repo, err := db.GetRepositoryByName(owner.ID, repoName)
	if err != nil {
		c.NotFoundOrError(err, "get repo by owner name and repository name")
		return
	}

	repoUrl := c.BaseURL + "/" + repo.Owner.Name + "/" + repoName
	repoMatadata := RepositoryMetadata{
		Name:        repo.Name,
		Description: repo.Description,
		Url:         repoUrl,
	}
	c.JSONSuccess(repoMatadata)
}
