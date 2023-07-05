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

	c.Data["JupyterContainer"] = res
	c.Data["HasWorkflows"] = context.HasTreeInRepo(c, "/WORKFLOWS")

	c.Success(VIEW_CONTAINER)
	return
}
