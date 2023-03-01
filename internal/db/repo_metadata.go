package db

import (
	"net/http"
	"path/filepath"
	"strings"

	"crypto/sha256"
	"encoding/hex"

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
		return nil, nil, err
	}
	log.Trace("GetLastCommitByBranch() commit_id : %s", commit_id)

	// get tree object id by commit_id
	tree_id, err := gitcmd.GetTreeIDByCommitId(repoPath, commit_id)
	if err != nil {
		return nil, nil, err
	}
	log.Trace("GetLastCommitByBranch() tree_id : %s", tree_id)

	// read tree on bare repo
	if err = gitcmd.GitReadTree(repoPath, tree_id); err != nil {
		return nil, nil, err
	}

	// get data list from repository
	data_list, err := gitcmd.GetFileDetailList(repoPath)
	if err != nil {
		return nil, nil, err
	}
	git_contents, annex_contents := gitcmd.DivideByMode(data_list)
	files := []datastruct.File{}

	// extract git/git-annex content metadat
	git_files, err := repo.ExtractGitContent(git_contents)
	if err != nil {
		return nil, nil, err
	}
	git_annex_files, err := repo.ExtractGitContent(annex_contents)
	if err != nil {
		return nil, nil, err
	}
	files = append(files, git_files...)
	files = append(files, git_annex_files...)

	return files, []datastruct.Dataset{}, nil
}

func (repo *Repository) ExtractGitContent(git_contents []gitcmd.DataDetail) ([]datastruct.File, error) {
	files := []datastruct.File{}
	for _, git_content := range git_contents {
		repoPath := repo.RepoPath()
		file_size, err := gitcmd.GetFileSizeByObjectId(repoPath, git_content.Hash)
		if err != nil {
			return nil, err
		}
		content, err := gitcmd.GetContentByObjectId(repoPath, git_content.Hash)
		if err != nil {
			return nil, err
		}
		mime_type := http.DetectContentType(content)

		if strings.Contains(mime_type, ";") {
			index := strings.Index(mime_type, ";")
			mime_type = mime_type[0:index]
		}
		r := sha256.Sum256(content)
		hash := hex.EncodeToString(r[:])

		// http://dg01.dg.rcos.nii.ac.jp/ivis-tsukioka/test_repo/src/master/WORKFLOWS/EX-WORKFLOWS/images/notebooks.diag
		log.Trace("repo.Name: %s", repo.Name)

		file := datastruct.File{
			ID:             git_content.FilePath,
			Name:           filepath.Base(git_content.FilePath),
			ContentSize:    file_size,
			EncodingFormat: mime_type,
			Sha256:         hash,
		}
		files = append(files, file)
	}
	return files, nil

}

func (repo *Repository) ExtractGitAnnexContent(git_annex_content []gitcmd.DataDetail) ([]datastruct.File, error) {
	return nil, nil
}
