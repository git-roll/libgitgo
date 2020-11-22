package libdiff

import (
    "github.com/git-roll/libgitgo/pkg/types"
)

func HeadToWorkDir(opt *types.Options) (*types.Diff, error) {
    return with(opt).HeadToWorkDir()
}

type wrapper interface {
    HeadToWorkDir() (*types.Diff, error)
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
