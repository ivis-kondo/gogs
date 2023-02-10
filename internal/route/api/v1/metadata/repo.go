package metadata

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	"github.com/NII-DG/gogs/internal/urlutil"
	log "unknwon.dev/clog/v2"
)

func GetRepo(c *context.APIContext) {
	req_user := c.User
	log.Trace("API to get Repository Metadata has been done by User[ID : %d]", req_user.ID)

	// Getting repository owner information from DB
	ownerName := c.Params(":ownername")
	owner, err := db.GetUserByName(ownerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal Server Error",
		})
		log.Error("failure getting user by name from DB. User Name : %s", ownerName)
		return
	}

	// Getting repository information from DB
	repoName := c.Params(":reponame")
	repo, err := db.GetRepositoryByName(owner.ID, repoName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal Server Error",
		})
		log.Error("failure getting repository by owner name and repository name from DB.  Repository : %s, Owner ID : %d", repoName, owner.ID)
		return
	}
	branchname := c.Params(":branchname")

	// check repository has branch
	if _, err := repo.GetBranch(branchname); err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"warm": fmt.Sprintf("this repository <%s/%s> dosen't have %s baranch.", ownerName, repoName, branchname),
		})
		log.Error("this repository <%s/%s> dosen't have %s baranch.", ownerName, repoName, branchname)
		return
	}

	// check request user access repository information
	users, err := repo.GetAssignees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal Server Error",
		})
		log.Error("failure getting read only user on repository from DB, Repository : %s", repo.Name)
		return
	}

	request_user_id := req_user.ID
	accessRight := false
	for _, u := range users {
		if u.ID == request_user_id {
			accessRight = true
		}
	}

	if !accessRight {

		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"warm": fmt.Sprintf("you do not has access right to get repository <%s of %s> metadata.", repoName, ownerName),
		})
		log.Trace("user<%s> do not has access right to get repository <%s of %s> metadata.", req_user.Name, repoName, ownerName)
		return
	}

	path := fmt.Sprintf("%s/%s/archive/%s.zip", repo.Owner.Name, repoName, branchname)
	url, err := urlutil.UpdatePath(c.BaseURL, path)
	if err != nil {
		c.Errorf(err, "%v", err)
		return
	}

	updatetime := time.Unix(repo.UpdatedUnix, 0).Format("2006-01-02")
	download := DownloadMetadat{
		Url:         url,
		Description: fmt.Sprint(c.Tr("metadata.download.description", fmt.Sprintf("%s/%s", repo.Owner.Name, repoName))),
		Date:        updatetime,
	}

	// Creating Repository Metadata
	path = repo.Owner.Name + "/" + repoName
	url, err = urlutil.UpdatePath(c.BaseURL, path)
	if err != nil {
		c.Errorf(err, "%v", err)
		return
	}
	repoMatadata := RepositoryMetadata{
		Name:        repo.Name,
		Description: repo.Description,
		Url:         url,
		Download:    download,
	}
	c.JSONSuccess(repoMatadata)
}
