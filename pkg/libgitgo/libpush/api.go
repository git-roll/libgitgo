package libpush

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

func CurBranch(remote string, force bool, opt *types.Options) (remoteOut string, err error) {
    return with(opt).Start([]string{""}, remote, force)
}

func Branch(branch, remote string, force bool, opt *types.Options) (remoteOut string, err error) {
    return with(opt).Start([]string{branch}, remote, force)
}

func AllBranches(remote string, force bool, opt *types.Options) (remoteOut string, err error) {
    return with(opt).Start(nil, remote, force)
}

type wrapper interface {
    Start(branch []string, remote string, force bool) (remoteOut string, err error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGit2Go:
        return &git2go{opt}
    case types.PreferGoGit:
        fallthrough
    default:
        return &goGit{opt}
    }
}
