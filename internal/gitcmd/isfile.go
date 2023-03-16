package gitcmd

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/NII-DG/gogs/internal/utils"
	constval "github.com/NII-DG/gogs/internal/utils/const"
	"github.com/gogs/git-module"
	log "unknwon.dev/clog/v2"
)

/*
ARG
----------------------

option : string. ["-c", "-d", "-m", "-o", "-s"]
*/
func GitIsFile(repoPath, option string) (string, error) {
	cmd := git.NewCommand("ls-files", option)
	raw_msg, err := cmd.RunInDir(repoPath)
	if err != nil {
		return "", fmt.Errorf("[%v]. exec cmd : [%v]. exec dir : [%s]", err, cmd.String(), repoPath)
	}
	return utils.BytesToString(raw_msg), nil
}

type DataDetail struct {
	Mode     string
	Hash     string
	FilePath string
}

func GetFileDetailList(repoPath string) ([]DataDetail, error) {
	raw_msg, err := GitIsFile(repoPath, "-s")
	log.Trace("[GetFileDetailList()] raw_msg : %s", raw_msg)
	if err != nil {
		return []DataDetail{}, err
	}
	reg := "\r\n|\n"
	file_list := regexp.MustCompile(reg).Split(raw_msg, -1)
	log.Trace("[GetFileDetailList()] file_list : %v", file_list)
	file_list = file_list[0 : len(file_list)-1]
	FileDetailList := []DataDetail{}

	for _, v := range file_list {
		log.Trace("[GetFileDetailList()] v : %v", v)
		aaaa := strings.Split(v, " ")
		log.Trace("[GetFileDetailList()] len(aaaa) : %v", len(aaaa))
		log.Trace("[GetFileDetailList()] aaaa : %v", aaaa)
		for i, v := range aaaa {
			log.Trace("[GetFileDetailList()] aaaa i : %d, v : %s", i, v)
		}
		bbbb := strings.Split(v, "　")
		for i, v := range bbbb {
			log.Trace("[GetFileDetailList()] bbbb i : %d, v : %s", i, v)
		}
		log.Trace("[GetFileDetailList()] len(bbbb) : %v", len(bbbb))
		log.Trace("[GetFileDetailList()] bbbb : %v", bbbb)
		file_info := strings.Fields(v)
		log.Trace("[GetFileDetailList()] len(file_info) : %v", len(file_info))
		log.Trace("[GetFileDetailList()] file_info : %v", file_info)
		log.Trace("[GetFileDetailList()] file_info[0] : %v", file_info[0])
		log.Trace("[GetFileDetailList()] file_info[1] : %v", file_info[1])
		log.Trace("[GetFileDetailList()] file_info[3] : %v", file_info[3])
		log.Trace("[GetFileDetailList()] filepath.ToSlash(file_info[3]) : %v", filepath.ToSlash(file_info[3]))
		var file_path string
		if len(file_info) >= 5 {
			data_com := file_info[3:]
			file_path = strings.Join(data_com, " ")
		} else {
			file_path = filepath.ToSlash(file_info[3])
		}
		fileDateil := DataDetail{
			Mode:     file_info[0],
			Hash:     file_info[1],
			FilePath: filepath.ToSlash(file_path),
		}
		FileDetailList = append(FileDetailList, fileDateil)
	}
	return FileDetailList, nil
}

func DivideByMode(data_list []DataDetail) (file_list []DataDetail, symbolic_link_list []DataDetail) {

	for _, v := range data_list {
		log.Trace("[DivideByMode()] v.FilePath : %s", v.FilePath)
		switch v.Mode {

		case "120000": // symbolic_link
			symbolic_link_list = append(symbolic_link_list, v)
		case "100644": // file
			file_list = append(file_list, v)

		}
	}
	return file_list, symbolic_link_list
}

func (dd DataDetail) IsExperimentPackage(data_struct_type string) (bool, error) {
	splited_file_path := strings.Split(filepath.ToSlash(dd.FilePath), "/")
	if data_struct_type == constval.WITH_CODE {
		return IsExperimentPackageOnWithCode(splited_file_path), nil
	} else if data_struct_type == constval.FOR_PARAMETER {
		return IsExperimentPackageOnForParameter(splited_file_path), nil
	} else {
		return false, fmt.Errorf("data_struct_type[%s] is not defined", data_struct_type)
	}
}

//experiments/test_ex/source/s3/sample.ipynb
const EXPERIMENTS = "experiments"
const INPUT_DATA = "input_data"
const SOURCE = "source"
const OUTPUT_DATA = "output_data"
const PARAM = "param"
const GIT_KEEP = ".gitkeep"

func IsExperimentPackageOnWithCode(splited_file_path []string) bool {
	if splited_file_path[0] != EXPERIMENTS {
		return false
	}
	if splited_file_path[2] == INPUT_DATA || splited_file_path[2] == SOURCE || splited_file_path[2] == OUTPUT_DATA {
		return splited_file_path[len(splited_file_path)-1] != GIT_KEEP
	}
	return false
}

func IsExperimentPackageOnForParameter(splited_file_path []string) bool {
	if splited_file_path[0] != EXPERIMENTS {
		return false
	}

	if splited_file_path[2] == INPUT_DATA || splited_file_path[2] == SOURCE {
		return splited_file_path[len(splited_file_path)-1] != GIT_KEEP
	}

	if splited_file_path[3] == PARAM || splited_file_path[3] == OUTPUT_DATA {
		return splited_file_path[len(splited_file_path)-1] != GIT_KEEP
	}
	return false
}
