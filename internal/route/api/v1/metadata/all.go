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

	//Create Metadata Structure
	baseUrl, err := urlutil.UpdatePath(c.BaseURL, "")
	if err != nil {
		c.Errorf(err, "%v", err)
		return
	}
	prefixPath := fmt.Sprintf("%s/%s/src/%s", ownerName, repoName, form.BranchName)
	svc := ServiceMetadata{
		Name:                "gin-fork",
		BaseUrl:             baseUrl,
		DataAccessUrlPrefix: prefixPath,
	}

	r_pj := ResearchProjectMetadata{
		Name:        repo.ProtectName,
		Description: repo.ProjectDescription,
	}

	path := fmt.Sprintf("%s/%s/archive/%s.zip", repo.Owner.Name, repoName, form.BranchName)
	download_url, err := urlutil.UpdatePath(c.BaseURL, path)
	if err != nil {
		c.Errorf(err, "%v", err)
		return
	}

	updatetime := time.Unix(repo.UpdatedUnix, 0).Format("2006-01-02")
	download := DownloadMetadat{
		Url:         download_url,
		Description: fmt.Sprint(c.Tr("metadata.download.description", fmt.Sprintf("%s/%s", repo.Owner.Name, repoName))),
		Date:        updatetime,
	}

	url, err := urlutil.UpdatePath(c.BaseURL, fmt.Sprintf("%s/%s", ownerName, repoName))
	if err != nil {
		c.Errorf(err, "%v", err)
		return
	}

	repository := RepositoryMetadata{
		Name:        repo.Name,
		Description: repo.Description,
		Url:         url,
		Download:    download,
	}

	assignees, _ := repo.GetAssignees()
	for _, u := range assignees {
		log.Trace("assignees User : %s", u.Name)
	}
	writers, _ := repo.GetWriters()
	for _, u := range writers {
		log.Trace("writers User : %s", u.Name)
	}

	wm := WholeMetadata{
		Service:         svc,
		ResearchProject: r_pj,
		Repository:      repository,
	}

	c.JSONSuccess(wm)
}
