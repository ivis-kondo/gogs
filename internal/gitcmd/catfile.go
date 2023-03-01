package gitcmd

import (
	"fmt"
	"regexp"

	"github.com/NII-DG/gogs/internal/utils"
	"github.com/gogs/git-module"
	log "unknwon.dev/clog/v2"
)

func GitCatFile(repoPath, hash string) (string, error) {
	cmd := git.NewCommand("cat-file", "-p", hash)
	raw_msg, err := cmd.RunInDir(repoPath)
	if err != nil {
		return "", fmt.Errorf("[%v]. exec cmd : [%v]. exec dir : [%s]", err, cmd.String(), repoPath)
	}
	return utils.BytesToString(raw_msg), nil
}

func GetTreeIDByCommitId(repoPath, commit_id string) (string, error) {
	raw_msg, err := GitCatFile(repoPath, commit_id)
	if err != nil {
		return "", err
	}
	reg := "\r\n|\n"
	msg_list := regexp.MustCompile(reg).Split(raw_msg, -1)
	log.Trace("GetTreeIDByCommitId raw_msg : %s", raw_msg)
	log.Trace("GetTreeIDByCommitId msg_list[0] : %s", msg_list[0])
	return "", nil
}
