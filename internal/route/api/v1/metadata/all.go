package metadata

import (
	"fmt"
	"net/http"

	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	log "unknwon.dev/clog/v2"
)

func GetAllMetadata(c *context.APIContext, form Repository) {
	req_user := c.User
	log.Trace("API to get Research All Metadata[Repository : %s/%s, Branch : %s] has been done by User[ID : %d]", form.OwnerName, form.RepoName, form.BranchName, req_user.ID)

	// Getting repository owner information from DB
	ownerName := form.OwnerName
	owner, err := db.GetUserByName(ownerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal Server Error",
		})
		log.Error("failure getting user by name from DB. User Name : %s", ownerName)
		return
	}

	// Getting repository information from DB
	repoName := form.RepoName
	repo, err := db.GetRepositoryByName(owner.ID, repoName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal Server Error",
		})
		log.Error("failure getting repository by owner name and repository name from DB.  Repository : %s, Owner ID : %d", repoName, owner.ID)
		return
	}

	// check repository has branch
	if _, err := repo.GetBranch(form.BranchName); err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"warm": fmt.Sprintf("this repository <%s/%s> dosen't have %s baranch.", ownerName, repoName, form.BranchName),
		})
		log.Error("this repository <%s/%s> dosen't have %s baranch.", ownerName, repoName, form.BranchName)
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
			"warm": fmt.Sprintf("you do not has access right to get Reaserch Project <%s of %s> metadata.", repoName, ownerName),
		})
		log.Trace("user<%s> do not has access right to get Reaserch Project <%s of %s> metadata.", req_user.Name, repoName, ownerName)
		return
	}

	c.JSONSuccess(WholeMetadata{})
}
