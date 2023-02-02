package metadata

import (
	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	"github.com/NII-DG/gogs/internal/urlutil"
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
