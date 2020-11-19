package librepo

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
    "github.com/go-git/go-git/v5"
)

type goGit struct {
    *types.Options
}

func (g goGit) HEAD() (*types.HEAD, error) {
    repo, err := g.OpenGoGitRepo()
    if err != nil {
        return nil, err
    }

    head, err := repo.Head()
    if err != nil {
        return nil, err
    }

    return &types.HEAD{GoGit: head}, nil
}

func (g goGit) Start() (*types.Repository, error) {
    r, err := g.OpenGoGitRepo()
    if err != nil {
        return nil, err
    }

    return &types.Repository{GoGit: r}, nil
}

func (g goGit) Init(bare bool) (repo *types.Repository, err error) {
    r, err := git.PlainInit(g.Options.WorkDir, bare)
    if err != nil {
        return
    }

    return &types.Repository{GoGit: r}, nil
}
