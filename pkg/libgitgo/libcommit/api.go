package libcommit

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

func Get(ref string, opt *types.Options) (*types.Commit, error) {
	return with(opt).Get(ref)
}

type Git2GoCommitOptions struct {
	RefName string
	Parent  string
}

type CommitOptions struct {
	All       bool
	Author    *types.User
	Committer *types.User
	Git2Go    Git2GoCommitOptions
}

func CommitStaging(message string, commitOpt *CommitOptions, opt *types.Options) (*types.Commit, error) {
	return with(opt).CommitStaging(message, commitOpt)
}

func Amend(message string, commitOpt *CommitOptions, opt *types.Options) (*types.Commit, error) {
	return with(opt).Amend(message, commitOpt)
}

func IsAncestor(ancestor, second string, opt *types.Options) (bool, error) {
	return with(opt).IsAncestor(ancestor, second)
}

type wrapper interface {
	Get(ref string) (*types.Commit, error)
	CommitStaging(message string, commitOpt *CommitOptions) (*types.Commit, error)
	Amend(message string, commitOpt *CommitOptions) (*types.Commit, error)
	IsAncestor(ancestor, second string) (bool, error)
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
