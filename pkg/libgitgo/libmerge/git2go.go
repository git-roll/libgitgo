package libmerge

import (
	"fmt"
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	*types.Options
}

func (g git2go) FastForward(branch, remote string) (err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	var refSpec string
	if len(remote) > 0 {
		refSpec = fmt.Sprintf("refs/remotes/%s/%s", remote, branch)
	} else {
		refSpec = fmt.Sprintf("refs/heads/%s", branch)
	}

	commit, err := repo.AnnotatedCommitFromRevspec(refSpec)
	if err != nil {
		return
	}

	analysis, _, err := repo.MergeAnalysis([]*git.AnnotatedCommit{commit})
	if err != nil {
		return
	}

	if (analysis & git.MergeAnalysisUpToDate) > 0 {
		return nil
	}

	if (analysis & git.MergeAnalysisFastForward) == 0 {
		return fmt.Errorf("can't merge if not fast forward")
	}

	head, err := repo.Head()
	if err != nil {
		return
	}

	_, err = head.SetTarget(commit.Id(), "")
	return
}
