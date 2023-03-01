package db

import (
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

	repoPath := repo.RepoPath()

	// get last commit ID by branch
	commit_id, err := gitcmd.GetLastCommitByBranch(repoPath, branch)
	if err != nil {
		return []datastruct.File{}, []datastruct.Dataset{}, err
	}
	log.Trace("GetLastCommitByBranch() commit_id : %s", commit_id)

	// get tree object id by commit_id
	tree_id, err := gitcmd.GetTreeIDByCommitId(repoPath, commit_id)
	if err != nil {
		return []datastruct.File{}, []datastruct.Dataset{}, err
	}
	log.Trace("GetLastCommitByBranch() tree_id : %s", tree_id)

	// read tree on bare repo
	if err = gitcmd.GitReadTree(repoPath, tree_id); err != nil {
		return []datastruct.File{}, []datastruct.Dataset{}, err
	}

	return []datastruct.File{}, []datastruct.Dataset{}, nil
}
