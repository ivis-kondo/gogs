package db

import (
	"strings"

	"github.com/NII-DG/gogs/internal/gitcmd"
	datastruct "github.com/NII-DG/gogs/internal/route/api/v1/metadata/datastruct"
	"github.com/unknwon/com"
	log "unknwon.dev/clog/v2"
)

/*
RCOS Function
Extract metadata from bere Repository
*/
func (repo *Repository) ExtractMetadata(branch string) ([]datastruct.File, []datastruct.Dataset, error) {

	// exclusive control for same repository
	pool_ID := "bere-" + com.ToStr(repo.ID)
	repoWorkingPool.CheckIn(pool_ID)
	defer repoWorkingPool.CheckOut(pool_ID)

	commit_id, err := gitcmd.GetLastCommitByBranch(repo.RepoPath(), branch)
	if err != nil {
		return []datastruct.File{}, []datastruct.Dataset{}, err
	}
	log.Trace("GetLastCommitByBranch() commit_id : %s", commit_id)
	log.Trace("GetLastCommitByBranch() commit_id : %s", strings.Replace(commit_id, "\"", "", -1))

	return []datastruct.File{}, []datastruct.Dataset{}, nil
}
