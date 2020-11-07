package repo

import (
    "github.com/git-roll/git-cli/pkg/libgitgo/types"
    git "github.com/libgit2/git2go/v31"
)

type git2go struct {
    worktree string
}

func (g git2go) Init(bare bool) (repo *types.Repository, err error) {
    r, err := git.InitRepository(g.worktree, bare)
    if err != nil {
        return
    }

    return &types.Repository{Git2Go: r}, nil
}
