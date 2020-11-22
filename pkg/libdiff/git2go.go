package libdiff

import (
	"github.com/git-roll/libgitgo/pkg/types"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	*types.Options
}

func (g git2go) HeadToWorkDir() (*types.Diff, error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return nil, err
	}

	head, err := repo.Head()
	if err != nil {
		return nil, err
	}

	headCommit, err := repo.LookupCommit(head.Target())
	if err != nil {
		return nil, err
	}

	headTree, err := headCommit.Tree()
	if err != nil {
		return nil, err
	}

	opt, err := git.DefaultDiffOptions()
	if err != nil {
		return nil, err
	}

	diff, err := repo.DiffTreeToWorkdir(headTree, &opt)
	if err != nil {
		return nil, err
	}

	return &types.Diff{Git2Go: diff}, nil
}
