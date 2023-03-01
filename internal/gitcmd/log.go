package gitcmd

import (
	"fmt"

	"github.com/NII-DG/gogs/internal/utils"
	"github.com/gogs/git-module"
)

/*
Function : Getting Git log

git log --pretty=format:<format : string> -<count : int> <branch : string>

ARG

-------------------
repoPath :Repository　Dir

branch : branch name

format : output format (ex: "%h - %an, %ar : %s" --> ca82a6d - Scott Chacon, 6 years ago : changed the version number). ref: https://git-scm.com/docs/git-log

count : Number of output
*/
func GitLog(repoPath, branch, format string, count int) (string, error) {
	pretty := fmt.Sprintf("--pretty=format:\"%s\"", format)
	count_num := fmt.Sprintf("-%d", count)
	cmd := git.NewCommand("log", pretty, count_num, branch)
	raw_msg, err := cmd.RunInDir(repoPath)
	if err != nil {
		return "", fmt.Errorf("error msg : [%v]. exec cmd : [%v]", err, cmd.String())
	}
	return utils.BytesToString(raw_msg), nil
}

func GetLastCommitByBranch(repoPath, branch string) (string, error) {
	return GitLog(repoPath, branch, "%H", 1)
}
