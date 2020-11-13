package libcommit

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

func Get(ref string, opt *types.Options) (*types.Commit, error) {
	return with(opt).Get(ref)
}

type CommitOptions struct {
	All       bool
	Author    *types.User
	Committer *types.User
}

func CommitStaging(message string, commitOpt *CommitOptions, opt *types.Options) (*types.Commit, error) {
	return with(opt).CommitStaging(message, commitOpt)
}

type wrapper interface {
	Get(ref string) (*types.Commit, error)
	CommitStaging(message string, commitOpt *CommitOptions) (*types.Commit, error)
}

func with(opt *types.Options) wrapper {
	switch opt.PreferredLib {
	case types.PreferGoGit:
		return &goGit{opt}
	case types.PreferGit2Go:
		return &git2go{opt}
	}

	return nil
}
