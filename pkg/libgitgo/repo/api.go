package repo

import "github.com/git-roll/libgitgo/pkg/libgitgo/types"

func Init(bare bool, opt *types.Options) (*types.Repository, error) {
    return with(opt).Init(bare)
}

func Open(opt *types.Options) (*types.Repository, error) {
    return with(opt).Start()
}

type wrapper interface {
    Init(bare bool) (*types.Repository, error)
    Start() (*types.Repository, error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        return &goGit{workdir: opt.WorkDir}
    case types.PreferGit2Go:
        return &git2go{workdir: opt.WorkDir}
    }

    return nil
}
