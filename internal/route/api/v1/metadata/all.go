package metadata

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	"github.com/NII-DG/gogs/internal/urlutil"
	"github.com/NII-DG/gogs/internal/utils/regex"
	log "unknwon.dev/clog/v2"
)

func GetAllMetadata(c *context.APIContext) {
	repoid_str := c.Params(":repoid")
	if !regex.CheckNumeric(repoid_str) {
		c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"warm": "Repository ID is not Numeric.",
		})
		return
	}
	branch := c.Params(":branch")
	req_user := c.User
	log.Trace("API to get Research All Metadata[Repository ID : %s, Branch : %s] has been done by User[ID : %d]", repoid_str, branch, req_user.ID)

	// Getting repository information from DB
	repoid, _ := strconv.Atoi(repoid_str)
	repo, err := db.GetRepositoryByID(int64(repoid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal Server Error",
		})
		log.Error("failure getting repository by owner name and repository name from DB.  Repository ID : %s", repoid_str)
		return
	}

	// check repository has branch
	if _, err := repo.GetBranch(branch); err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"warm": fmt.Sprintf("this repository <ID : %s> dosen't have %s baranch.", repoid_str, branch),
		})
		log.Error("this repository <ID : %s> dosen't have %s baranch.", repoid_str, branch)
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
			"warm": fmt.Sprintf("you do not has access right to get repository <ID : %s> metadata.", repoid_str),
		})
		log.Trace("user<%s> do not has access right to get repository <ID : %s> metadata.", req_user.Name, repoid_str)
		return
	}

	//Create Metadata Structure
	baseUrl, err := urlutil.UpdatePath(c.BaseURL, "")
	if err != nil {
		c.Errorf(err, "%v", err)
		return
	}
	prefixPath := fmt.Sprintf("%s/%s/src/%s", repo.Owner.Name, repo.Name, branch)
	svc := ServiceMetadata{
		Name:                "gin-fork",
		BaseUrl:             baseUrl,
		DataAccessUrlPrefix: prefixPath,
	}

	r_pj := ResearchProjectMetadata{
		Name:        repo.ProtectName,
		Description: repo.ProjectDescription,
	}

	path := fmt.Sprintf("%s/%s/archive/%s.zip", repo.Owner.Name, repo.Name, branch)
	download_url, err := urlutil.UpdatePath(c.BaseURL, path)
	if err != nil {
		c.Errorf(err, "%v", err)
		return
	}

	updatetime := time.Unix(repo.UpdatedUnix, 0).Format("2006-01-02")
	download := DownloadMetadat{
		Url:         download_url,
		Description: fmt.Sprint(c.Tr("metadata.download.description", fmt.Sprintf("%s/%s", repo.Owner.Name, repo.Name))),
		Date:        updatetime,
	}

	url, err := urlutil.UpdatePath(c.BaseURL, fmt.Sprintf("%s/%s", repo.Owner.Name, repo.Name))
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
	user_list := []UserMatadata{}
	for _, u := range users {
		org := UserOrgMetadata{
			Name:        u.Affiliation,
			Url:         u.AffiliationURL,
			AliasName:   u.AffiliationAlias,
			Description: u.AffiliationDescription,
		}
		personalUrl := ""
		if len(u.PersonalURL) > 0 {
			personalUrl = u.PersonalURL
		} else {
			url, err := urlutil.UpdatePath(c.BaseURL, u.Name)
			if err != nil {
				c.Errorf(err, "%v", err)
				return
			}
			personalUrl = url
		}
		user := UserMatadata{
			UserName:    u.Name,
			Url:         personalUrl,
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			AliasName:   u.AliasName,
			EMail:       u.Email,
			Telephone:   u.Telephone,
			ERadNumber:  u.ERadResearcherNumber,
			Affiliation: org,
		}
		user_list = append(user_list, user)
	}

	wm := WholeMetadata{
		Service:         svc,
		ResearchProject: r_pj,
		Repository:      repository,
		Users:           user_list,
	}

	c.JSONSuccess(wm)
}
