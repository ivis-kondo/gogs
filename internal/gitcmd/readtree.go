package gitcmd

import (
	"fmt"

	"github.com/gogs/git-module"
)

func GitReadTree(repoPath, tree_id string) error {
	cmd := git.NewCommand("read-tree", tree_id)
	_, err := cmd.RunInDir(repoPath)
	if err != nil {
		return fmt.Errorf("[%v]. exec cmd : [%v]. exec dir : [%s]", err, cmd.String(), repoPath)
	}
	return nil
}
