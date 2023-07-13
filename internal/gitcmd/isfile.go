package gitcmd

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/NII-DG/gogs/internal/utils"
	//constval "github.com/NII-DG/gogs/internal/utils/const"
	"github.com/gogs/git-module"
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
	if err != nil {
		return []DataDetail{}, err
	}
	reg := "\r\n|\n"
	file_list := regexp.MustCompile(reg).Split(raw_msg, -1)
	file_list = file_list[0 : len(file_list)-1]

	FileDetailList := []DataDetail{}

	for _, v := range file_list {
		file_info := strings.Fields(v)
		var file_path string
		if len(file_info) >= 5 {
			head := file_info[3]
			index := strings.Index(v, head)
			file_path = v[index:]
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
		switch v.Mode {

		case "120000": // symbolic_link
			symbolic_link_list = append(symbolic_link_list, v)
		case "100644": // file
			file_list = append(file_list, v)

		}
	}
	return file_list, symbolic_link_list
}
