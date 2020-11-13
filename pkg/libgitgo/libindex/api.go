package libindex

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

func Add(paths []string, opt *types.Options) (err error) {
    return with(opt).Add(paths)
}

type wrapper interface {
    Add([]string) (err error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        return &goGit{opt}
    case types.PreferGit2Go:
        return &git2go{opt}
    }

    return nil
}
