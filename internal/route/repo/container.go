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
	res, err := db.GetJupyterContainerByRepoIDAndUserID(c.UserID(), c.Repo.Repository.ID)

	if err != nil {
		return err
	}
	displayRelaunch := false
	repoNames := make([]string, len(res))
	for i, repo := range res {
		rep, err2 := db.GetRepositoryByID(repo.RepoID)

		if rep != nil {
			repoNames[i] = rep.Name
		}
		if len(repo.ExperimentPackage) > 0 {
			displayRelaunch = true
		}

		if err2 != nil {
			return err2
		}
	}

	c.Data["JupyterContainer"] = res
	c.Data["repoNames"] = repoNames
	c.Data["displayRelaunch"] = displayRelaunch
	c.Success(VIEW_CONTAINER)
	return
}
