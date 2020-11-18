package libignore

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
    "golang.org/x/xerrors"
)

var ErrPathIgnored = xerrors.Errorf("path ignored")

func Check(relativePath string, opt *types.Options) (err error) {
    return with(opt).Check(relativePath)
}

type wrapper interface {
    Check(relativePath string) (err error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        return &goGit{opt}
    case types.PreferGit2Go:
        return &git2go{opt}
    default:
        panic(opt.PreferredLib)
    }

    return nil
}
