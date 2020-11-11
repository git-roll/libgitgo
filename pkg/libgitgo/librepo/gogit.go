package librepo

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
    "github.com/go-git/go-git/v5"
)

type goGit struct {
    workdir string
}

func (g goGit) Start() (*types.Repository, error) {
    r, err := git.PlainOpen(g.workdir)
    if err != nil {
        return nil, err
    }

    return &types.Repository{GoGit: r}, nil
}

func (g goGit) Init(bare bool) (repo *types.Repository, err error) {
    r, err := git.PlainInit(g.workdir, bare)
    if err != nil {
        return
    }

    return &types.Repository{GoGit: r}, nil
}
