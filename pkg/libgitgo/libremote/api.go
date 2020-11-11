package libremote

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

func List(opt *types.Options) ([]*types.Remote, error) {
    return with(opt).List()
}

func Create(name, url, fetchSpec string, opt *types.Options) (*types.Remote, error) {
    return with(opt).Create(name, url, fetchSpec)
}

func Lookup(name string, opt *types.Options) (*types.Remote, error) {
    return with(opt).Lookup(name)
}

type wrapper interface {
    List() ([]*types.Remote, error)
    Create(name, url, fetchSpec string) (*types.Remote, error)
    Lookup(name string) (*types.Remote, error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        return &goGit{worktree: opt.WorkDir}
    case types.PreferGit2Go:
        return &git2go{worktree: opt.WorkDir}
    }

    return nil
}
