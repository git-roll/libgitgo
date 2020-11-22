package libref

import (
    "github.com/git-roll/libgitgo/pkg/types"
)

func List(opt *types.Options) ([]*types.Reference, error) {
    return with(opt).List()
}

type wrapper interface {
    List() ([]*types.Reference, error)
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
