package repo

import (
	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
)

const (
	VIEW_CONTAINER = "repo/container"
)

func ViewContainer(c *context.Context) (err error) {
	c.Data["PageIsContainer"] = true
	res, err := db.GetJupyterContainer(c.Repo.Repository.ID, c.UserID())

	if err != nil {
		return err
	}
	repoNames := make([]string, len(res))
	for i, repo := range res {
		rep, err2 := db.GetRepositoryByID(repo.RepoID)

		if rep != nil {
			repoNames[i] = rep.Name
		}

		if err2 != nil {
			return err2
		}
	}

	c.Data["JupyterContainer"] = res
	c.Data["repoNames"] = repoNames
	c.Success(VIEW_CONTAINER)
	return
}
