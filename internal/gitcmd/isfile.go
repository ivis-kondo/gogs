package gitcmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/NII-DG/gogs/internal/utils"
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

type FileDetail struct {
	Mode     string
	Hash     string
	FilePath string
}

func GetFileDetailList(repoPath string) (string, error) {
	raw_msg, err := GitIsFile(repoPath, "-s")
	if err != nil {
		return "", err
	}
	reg := "\r\n|\n"
	file_list := regexp.MustCompile(reg).Split(raw_msg, -1)

	//FileDetailList := []FileDetail{}

	for _, v := range file_list {
		file_info := strings.Split(v, " ")
		for index, va := range file_info {
			log.Trace("index : %d, value : %s", index, va)
		}
	}
	return "", nil
}
