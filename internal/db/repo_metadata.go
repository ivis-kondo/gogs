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
	"github.com/NII-DG/gogs/internal/utils/const_utils"
	"github.com/unknwon/com"
	log "unknwon.dev/clog/v2"
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

	// extract git/git-annex content metadat
	git_files, err := ExtractMetaDataGitContent(repo, git_contents, branch)
	if err != nil {
		return nil, nil, datastruct.GinMonitoring{}, err
	}

	git_annex_files, err := ExtractMetaDataGitAnnexContent(repo, annex_contents, branch)
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

			//create ExperimentPackageList from datasets
			//create ParameterExperimentList datasets
			experimentPackageList, parameterExperimentList := ExtractExperimentPackageList(datasetStructure, datasets)
			gin_monitoring = datastruct.GinMonitoring{
				WorkflowIdentifier:      field,
				ContentSize:             dataSize,
				DatasetStructure:        datasetStructure,
				ExperimentPackageList:   experimentPackageList,
				ParameterExperimentList: parameterExperimentList,
			}

		}
	}

	if !isDMP {
		return nil, nil, datastruct.GinMonitoring{}, nil
	}

	return files, datasets, gin_monitoring, nil
}

func ExtractExperimentPackageList(struct_type string, datasets []datastruct.Dataset) ([]string, []string) {
	log.Error("[ExtractExperimentPackageList()] len(datasets): %d", len(datasets))
	experimentPackageList := []string{}
	parameterExperimentList := []string{}
	//create ExperimentPackageList from datasets
	for _, dataset := range datasets {
		log.Error("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		log.Error("[ExtractExperimentPackageList()] dataset.ID: %s", dataset.ID)
		if IsExperimentPackage(dataset.ID) {
			log.Error("[ExtractExperimentPackageList()] dataset.ID: %s is ExperimentPackage", dataset.ID)
			path_compoment := strings.Split(filepath.ToSlash(dataset.ID), "/")
			if len(path_compoment[:len(path_compoment)-1]) == 2 {
				isInvoled := false
				for _, v := range experimentPackageList {
					if v != dataset.ID {
						isInvoled = true
					}
				}
				if !isInvoled {
					log.Error("[ExtractExperimentPackageList()] dataset.ID: %s add to experimentPackageList", dataset.ID)
					experimentPackageList = append(experimentPackageList, dataset.ID)
				}
			}
			if struct_type == const_utils.GetForParameters() && len(path_compoment[:len(path_compoment)-1]) == 3 {
				log.Error("[ExtractExperimentPackageList()] struct_type: %s is ForParameters", struct_type)
				isInvoled := false
				for _, v := range parameterExperimentList {
					if v != dataset.ID {
						isInvoled = true
						parameterExperimentList = append(parameterExperimentList, dataset.ID)
					}
				}
				if !isInvoled && const_utils.IsParameterFolder(path_compoment[2]) {
					log.Error("[ExtractExperimentPackageList()] dataset.ID: %s add to parameterExperimentList", dataset.ID)
					parameterExperimentList = append(parameterExperimentList, dataset.ID)
				}
			}
		}
	}
	return experimentPackageList, parameterExperimentList

}

func ExtractMetaDataGitContent(repo *Repository, git_contents []gitcmd.DataDetail, branch string) ([]datastruct.File, error) {
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

		url, err := repo.CreateAccessUrlToFile(branch, git_content.FilePath)
		if err != nil {
			return nil, err
		}

		isExperimentPackageFlag := IsExperimentPackage(git_content.FilePath)

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

func ExtractMetaDataGitAnnexContent(repo *Repository, git_annex_contents []gitcmd.DataDetail, branch string) ([]datastruct.File, error) {
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

		isExperimentPackageFlag := IsExperimentPackage(git_annex_content.FilePath)

		url, err := repo.CreateAccessUrlToFile(branch, git_annex_content.FilePath)
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

func (repo *Repository) CreateAccessUrlToFile(branch, file_path string) (string, error) {
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
				dataset_url, err := repo.CreateAccessUrlToFile(branch, dataset_id)
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

func (repo *Repository) CreateAccessUrlToRepo() (string, error) {
	urlPath := repo.FullName()
	url, err := urlutil.UpdatePath(conf.Server.ExternalURL, urlPath)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (repo *Repository) ExtractRepoMetadata() (datastruct.RepositoryObject, error) {
	url, err := repo.CreateAccessUrlToRepo()
	if err != nil {
		return datastruct.RepositoryObject{}, err
	}
	return datastruct.RepositoryObject{
		ID:          url,
		Name:        repo.Name,
		Description: repo.Description,
	}, nil

}

const EXPERIMENTS = "experiments"
const GIT_KEEP = ".gitkeep"

func IsExperimentPackage(file_path string) bool {
	splited_file_path := strings.Split(filepath.ToSlash(file_path), "/")
	if splited_file_path[0] != EXPERIMENTS {
		return false
	}
	if splited_file_path[len(splited_file_path)-1] != GIT_KEEP {
		return true
	}
	return false
}
