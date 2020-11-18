package libmerge

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
    "golang.org/x/xerrors"
)

var (
    ErrWorktreeIsNotClean = xerrors.Errorf("worktree is not clean")
    ErrConflictAfterMerging = xerrors.Errorf("conflicts occur")
)

func FastForward(branch, remote string, opt *types.Options) (err error) {
    return with(opt).FastForward(branch, remote)
}

type MergeOptions struct {
    Author    *types.User
    Committer *types.User
}

func Start(branch, remote string, mergeOpt *MergeOptions, opt *types.Options) (err error) {
    return with(opt).Start(branch, remote, mergeOpt)
}

type wrapper interface {
    FastForward(branch, remote string) (err error)
    Start(branch, remote string, mergeOpt *MergeOptions) (err error)
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
