package librebase

import (
    "github.com/git-roll/libgitgo/pkg/types"
)

type RebaseOptions struct {
    Author    *types.User
    Committer *types.User
}

func CompactPrivateCommits(upstream, messagePrefix string, rebaseOpt *RebaseOptions, opt *types.Options) (err error) {
    return with(opt).CompactPrivateCommits(upstream, messagePrefix, rebaseOpt)
}

func Start(branch, upstream, onto string, rebaseOpt *RebaseOptions, opt *types.Options) error {
    return with(opt).Start(branch, upstream, onto, rebaseOpt)
}

func Abort(opt *types.Options) error {
    return with(opt).Abort()
}

type wrapper interface {
    CompactPrivateCommits(upstream, messagePrefix string, rebaseOpt *RebaseOptions) (err error)
    Start(branch, upstream, onto string, rebaseOpt *RebaseOptions) error
    Abort() error
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
