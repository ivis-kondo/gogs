package gitcmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/NII-DG/gogs/internal/utils"
	"github.com/gogs/git-module"
	//log "unknwon.dev/clog/v2"
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

	FileDetailList := []DataDetail{}

	for _, v := range file_list {
		file_info := strings.Fields(v)
		fileDateil := DataDetail{
			Mode:     file_info[0],
			Hash:     file_info[1],
			FilePath: file_info[3],
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
