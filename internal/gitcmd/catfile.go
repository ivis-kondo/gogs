package gitcmd

import (
	"fmt"

	"github.com/NII-DG/gogs/internal/utils"
	"github.com/gogs/git-module"
)

func GitCatFile(repoPath, hash string) (string, error) {
	cmd := git.NewCommand("cat-file", "-p", hash)
	raw_msg, err := cmd.RunInDir(repoPath)
	if err != nil {
		return "", fmt.Errorf("error msg : [%v]. exec cmd : [%v]", err, cmd.String())
	}

	return utils.BytesToString(raw_msg), nil
}
