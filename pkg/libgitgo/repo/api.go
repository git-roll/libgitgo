package repo

import "github.com/git-roll/git-cli/pkg/libgitgo/types"

func Init(bare bool, opt *types.Options) (*types.Repository, error) {
    return with(opt).Init(bare)
}

type wrapper interface {
    Init(bare bool) (*types.Repository, error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        return &goGit{worktree: opt.WorkDir}
    case types.PreferGit2Go:
        return &git2go{worktree: opt.WorkDir}
    }

    return nil
}
