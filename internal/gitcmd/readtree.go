package gitcmd

import (
	"fmt"

	"github.com/NII-DG/gogs/internal/utils"
	"github.com/gogs/git-module"
	log "unknwon.dev/clog/v2"
)

func GitReadTree(repoPath, tree_id string) error {
	cmd := git.NewCommand("read-tree-file", tree_id)
	raw_msg, err := cmd.RunInDir(repoPath)
	if err != nil {
		return fmt.Errorf("[%v]. exec cmd : [%v]. exec dir : [%s]", err, cmd.String(), repoPath)
	}
	log.Trace("GitReadTree : raw_msg : %s", utils.BytesToString(raw_msg))

	return nil
}
