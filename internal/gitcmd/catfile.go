package gitcmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/NII-DG/gogs/internal/utils"
	"github.com/gogs/git-module"
)

func GitCatFile(repoPath, option, hash string) ([]byte, error) {
	cmd := git.NewCommand("cat-file", option, hash)
	raw_msg, err := cmd.RunInDir(repoPath)
	if err != nil {
		return nil, fmt.Errorf("[%v]. exec cmd : [%v]. exec dir : [%s]", err, cmd.String(), repoPath)
	}
	return raw_msg, nil
}

func GetTreeIDByCommitId(repoPath, commit_id string) (string, error) {
	raw_msg, err := GitCatFile(repoPath, "-p", commit_id)
	if err != nil {
		return "", err
	}
	reg := "\r\n|\n"
	tree_id := strings.Split(regexp.MustCompile(reg).Split(utils.BytesToString(raw_msg), -1)[0], " ")[1]
	return tree_id, nil
}

func GetFileSizeByObjectId(repoPath, object_id string) (string, error) {
	raw_msg, err := GitCatFile(repoPath, "-s", object_id)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(utils.BytesToString(raw_msg), "\n", ""), nil
}

func GetContentByObjectId(repoPath, object_id string) ([]byte, error) {
	raw_msg, err := GitCatFile(repoPath, "-p", object_id)
	if err != nil {
		return nil, err
	}
	return raw_msg, nil
}
