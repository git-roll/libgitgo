package libmerge

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

func FastForward(branch, remote string, opt *types.Options) (err error) {
    return with(opt).FastForward(branch, remote)
}

type wrapper interface {
    FastForward(branch, remote string) (err error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        panic("go-git doesn't support to merge branches")
    case types.PreferGit2Go:
        fallthrough
    default:
        return &git2go{opt}
    }
}
