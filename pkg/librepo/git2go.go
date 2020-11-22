package librepo

import (
    "github.com/git-roll/libgitgo/pkg/types"
    git "github.com/libgit2/git2go/v31"
)

type git2go struct {
    *types.Options
}

func (g git2go) HEAD() (h *types.HEAD, err error) {
    repo, err := g.OpenGit2GoRepo()
    if err != nil {
        return
    }

    head, err := repo.Head()
    if err != nil {
        return
    }

    return &types.HEAD{Git2Go: head}, nil
}

func (g git2go) Start() (*types.Repository, error) {
    r, err := g.OpenGit2GoRepo()
    if err != nil {
        return nil, err
    }

    return &types.Repository{Git2Go: r}, nil
}

func (g git2go) Init(bare bool) (repo *types.Repository, err error) {
    r, err := git.InitRepository(g.Options.WorkDir, bare)
    if err != nil {
        return
    }

    return &types.Repository{Git2Go: r}, nil
}
