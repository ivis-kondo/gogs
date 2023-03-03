package db

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"crypto/sha256"
	"encoding/hex"

	"encoding/json"

	"github.com/NII-DG/gogs/internal/conf"
	"github.com/NII-DG/gogs/internal/gitcmd"
	"github.com/NII-DG/gogs/internal/gitcmd/annex"
	datastruct "github.com/NII-DG/gogs/internal/route/api/v1/metadata/datastruct"
	"github.com/NII-DG/gogs/internal/urlutil"
	"github.com/NII-DG/gogs/internal/utils"
	"github.com/unknwon/com"
)

/*
RCOS Function
Extract metadata from bere Repository
*/
func (repo *Repository) ExtractMetadata(branch string) ([]datastruct.File, []datastruct.Dataset, datastruct.GinMonitoring, error) {

	// exclusive control for same repository
	pool_ID := "bere-" + com.ToStr(repo.ID)
	repoWorkingPool.CheckIn(pool_ID)
	defer repoWorkingPool.CheckOut(pool_ID)

	repoPath := repo.RepoPath()

	// get last commit ID by branch
	commit_id, err := gitcmd.GetLastCommitByBranch(repoPath, branch)
	if err != nil {
		return nil, nil, datastruct.GinMonitoring{}, err
	}

	// get tree object id by commit_id
	tree_id, err := gitcmd.GetTreeIDByCommitId(repoPath, commit_id)
	if err != nil {
		return nil, nil, datastruct.GinMonitoring{}, err
	}

	// read tree on bare repo
	if err = gitcmd.GitReadTree(repoPath, tree_id); err != nil {
		return nil, nil, datastruct.GinMonitoring{}, err
	}

	// get data list from repository
	data_list, err := gitcmd.GetFileDetailList(repoPath)
	if err != nil {
		return nil, nil, datastruct.GinMonitoring{}, err
	}
	git_contents, annex_contents := gitcmd.DivideByMode(data_list)
	files := []datastruct.File{}

	//get research policy
	isDMP := false
	var gin_monitoring datastruct.GinMonitoring
	for _, git_content := range git_contents {

		if strings.Contains(git_content.FilePath, "dmp.json") {
			isDMP = true
			dmp, err := gitcmd.GetContentByObjectId(repoPath, git_content.Hash)
			if err != nil {
				return nil, nil, datastruct.GinMonitoring{}, err
			}
			var jsonObj interface{}
			_ = json.Unmarshal(dmp, &jsonObj)
			field := jsonObj.(map[string]interface{})["workflowIdentifier"].(string)
			dataSize := jsonObj.(map[string]interface{})["contentSize"].(string)
			datasetStructure := jsonObj.(map[string]interface{})["datasetStructure"].(string)
			gin_monitoring = datastruct.GinMonitoring{
				WorkflowIdentifier: field,
				ContentSize:        dataSize,
				DatasetStructure:   datasetStructure,
			}

		}
	}

	if !isDMP {
		return nil, nil, datastruct.GinMonitoring{}, fmt.Errorf("dmp.json is not")
	}

	// extract git/git-annex content metadat
	git_files, err := ExtractMetaDataGitContent(repo, git_contents, branch, gin_monitoring.DatasetStructure)
	if err != nil {
		return nil, nil, datastruct.GinMonitoring{}, err
	}
	git_annex_files, err := ExtractMetaDataGitAnnexContent(repo, annex_contents, branch, gin_monitoring.DatasetStructure)
	if err != nil {
		return nil, nil, datastruct.GinMonitoring{}, err
	}
	files = append(files, git_files...)
	files = append(files, git_annex_files...)

	//create Dataset
	datasets, err := CreateFilesToDatasets(repo, files, branch)
	if err != nil {
		return nil, nil, datastruct.GinMonitoring{}, err
	}
	return files, datasets, gin_monitoring, nil
}

func ExtractMetaDataGitContent(repo *Repository, git_contents []gitcmd.DataDetail, branch, data_struct_type string) ([]datastruct.File, error) {
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

		url, err := repo.CreateAccessUrl(branch, git_content.FilePath)
		if err != nil {
			return nil, err
		}

		isExperimentPackageFlag, err := git_content.IsExperimentPackage(data_struct_type)
		if err != nil {
			return nil, err
		}
		file := datastruct.File{
			ID:                    git_content.FilePath,
			Name:                  filepath.Base(git_content.FilePath),
			ContentSize:           file_size,
			EncodingFormat:        mime_type,
			Sha256:                hash,
			Url:                   url,
			ExperimentPackageFlag: isExperimentPackageFlag,
		}
		files = append(files, file)
	}
	return files, nil

}

func ExtractMetaDataGitAnnexContent(repo *Repository, git_annex_contents []gitcmd.DataDetail, branch, data_struct_type string) ([]datastruct.File, error) {
	files := []datastruct.File{}
	for _, git_annex_content := range git_annex_contents {
		object_id := git_annex_content.Hash
		repoPath := repo.RepoPath()
		content, err := gitcmd.GetContentByObjectId(repoPath, object_id)
		if err != nil {
			return nil, err
		}
		annex_key := filepath.Base(utils.BytesToString(content))
		annex_key = strings.ReplaceAll(annex_key, "&c", ":")
		annex_key = strings.ReplaceAll(annex_key, "%", "/")
		annex_key = strings.ReplaceAll(annex_key, "&a", "&")
		annex_key = strings.ReplaceAll(annex_key, "&s", "%")
		field, err := annex.GetFieldsFromMetadata(repoPath, annex_key)
		if err != nil {
			return nil, err
		}
		isExperimentPackageFlag, err := git_annex_content.IsExperimentPackage(data_struct_type)
		if err != nil {
			return nil, err
		}

		url, err := repo.CreateAccessUrl(branch, git_annex_content.FilePath)
		if err != nil {
			return nil, err
		}

		file := datastruct.File{
			ID:                    git_annex_content.FilePath,
			Name:                  filepath.Base(git_annex_content.FilePath),
			ContentSize:           field.ContentSize,
			EncodingFormat:        field.EncodingFormat,
			Sha256:                field.Sha256,
			Url:                   url,
			SdDatePublished:       field.SdDatePublished,
			ExperimentPackageFlag: isExperimentPackageFlag,
		}
		files = append(files, file)
	}
	return files, nil
}

func (repo *Repository) CreateAccessUrl(branch, file_path string) (string, error) {
	urlPath := fmt.Sprintf("%s/src/%s/%s", repo.FullName(), branch, file_path)
	url, err := urlutil.UpdatePath(conf.Server.ExternalURL, urlPath)
	if err != nil {
		return "", err
	}
	return url, nil
}

func CreateFilesToDatasets(repo *Repository, files []datastruct.File, branch string) ([]datastruct.Dataset, error) {
	datasets := []datastruct.Dataset{}

	tmp_data := map[string]datastruct.Dataset{}

	for _, file := range files {
		path := file.ID
		splited_file_path := strings.Split(path, "/")
		splited_file_path = splited_file_path[0 : len(splited_file_path)-1]
		dataset_id := ""
		for _, element := range splited_file_path {
			dataset_id = dataset_id + element + "/"
			if _, ok := tmp_data[dataset_id]; !ok {
				dataset_url, err := repo.CreateAccessUrl(branch, dataset_id)
				if err != nil {
					return nil, err
				}
				dataset := datastruct.Dataset{
					ID:   dataset_id,
					Name: element,
					Url:  dataset_url,
				}
				tmp_data[dataset_id] = dataset
			}
		}
	}

	for _, v := range tmp_data {
		datasets = append(datasets, v)
	}
	return datasets, nil
}
