package libconfig

import (
    "github.com/git-roll/libgitgo/pkg/types"
)

func User(opt *types.Options) (*types.User, error) {
    return with(opt).User()
}

type wrapper interface {
    User() (*types.User, error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGit2Go:
        panic("git2go doesn't support reading configurations")
    case types.PreferGoGit:
        fallthrough
    default:
        return &goGit{opt}
    }

    return nil
}
