package metadata

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	datastruct "github.com/NII-DG/gogs/internal/route/api/v1/metadata/datastruct"
	"github.com/NII-DG/gogs/internal/urlutil"
	"github.com/NII-DG/gogs/internal/utils/regex"
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

func GetAllMetadataByRepoIDAndBranch(c *context.APIContext) {
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

	// Create ResearchProject
	Research_pj := datastruct.ResearchProject{
		Name:        repo.ProtectName,
		Description: repo.ProjectDescription,
	}
	log.Info("%s, %s", Research_pj.Name, Research_pj.Description)
	log.Info("repo.LocalCopyPath : %s", repo.LocalCopyPath())
	log.Info("repo.RepoPath : %s", repo.RepoPath())

	// Create Persons
	persons := []datastruct.Person{}
	for _, u := range users {

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

		person := datastruct.Person{
			ID:    u.IDStr(),
			Url:   personalUrl,
			Name:  u.FullName,
			Alias: u.AliasName,
			//Affiliation:          "",
			Email:                u.Email,
			Telephone:            u.Telephone,
			ERadResearcherNumber: u.ERadResearcherNumber,
		}
		persons = append(persons, person)
	}

	// Create Files and GinMonitoring
	files, dataset, gin_monitoring, err := repo.ExtractMetadata(branch)
	if err != nil {
		log.Error("failure extracting metadata from repository <ID : %s>. err msg : %v", repoid_str, err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal Server Error",
		})
		return
	}

	//TODO : Create dmps
	//TODO : Create funder_orgs
	//TODO : Create research_orgs
	//TODO : Create licenses
	//TODO : Create data_downloads
	//TODO : Create repository_objs
	//TODO : Create hosting_institutions

	// Create Metadata
	metadata := datastruct.Metadata{
		ResearchProject:     Research_pj,
		FunderOrgs:          []datastruct.FunderOrg{},
		ResearchOrgs:        []datastruct.ResearchOrg{},
		Licenses:            []datastruct.License{},
		DataDownloads:       []datastruct.DataDownload{},
		RepositoryObjects:   []datastruct.RepositoryObject{},
		HostingInstitutions: []datastruct.HostingInstitution{},
		Persons:             persons,
		Files:               files,
		Datasets:            dataset,
		GinMonitorings:      []datastruct.GinMonitoring{gin_monitoring},
		Dmps:                []datastruct.Dmp{},
	}

	c.JSONSuccess(metadata)
}
