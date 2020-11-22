package libindex

import (
    "github.com/git-roll/libgitgo/pkg/types"
)

func AddAll(opt *types.Options) (err error) {
    return with(opt).Add([]string{"."})
}

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
        fallthrough
    default:
        return &git2go{opt}
    }
}
