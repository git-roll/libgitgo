package librepo

import "github.com/git-roll/libgitgo/pkg/libgitgo/types"

func Init(bare bool, opt *types.Options) (*types.Repository, error) {
    repo, err := with(opt).Init(bare)
    if err != nil {
        return nil, err
    }

    if opt.FollowOpenedRepo {
        opt.WithRepo(repo)
    }

    return repo, err
}

func Open(opt *types.Options) (*types.Repository, error) {
    repo, err := with(opt).Start()
    if err != nil {
        return nil, err
    }

    if opt.FollowOpenedRepo {
        opt.WithRepo(repo)
    }

    return repo, err
}

func HEAD(opt *types.Options) (*types.HEAD, error) {
    return with(opt).HEAD()
}

type wrapper interface {
    Init(bare bool) (*types.Repository, error)
    Start() (*types.Repository, error)
    HEAD() (*types.HEAD, error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        return &goGit{opt}
    case types.PreferGit2Go:
        fallthrough
    default:
        return &git2go{opt}
    }
}
