package repo

import (
    "github.com/git-roll/git-cli/pkg/libgitgo/types"
    "github.com/go-git/go-git/v5"
)

type goGit struct {
    worktree string
}

func (g goGit) Init(bare bool) (repo *types.Repository, err error) {
    r, err := git.PlainInit(g.worktree, bare)
    if err != nil {
        return
    }

    return &types.Repository{GoGit: r}, nil
}
