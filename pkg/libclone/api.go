package libclone

import (
	"context"
	"github.com/git-roll/libgitgo/pkg/types"
)

type Options struct {
	Git2Go Git2GoOptions
	GoGit  GoGitOptions
}

type Git2GoOptions struct {
	DownloadTags types.Git2GoDownloadTags
}

type GoGitOptions struct {
	// usually `origin`
	RemoteName string
	// specified branch only
	SingleBranch          bool
	NoCheckout            bool
	Depth                 int
	SubmoduleRescursivity types.GoGitSubmoduleRescursivity
	TagMode               types.GoGitTagMode
}

func Start(url string, branch string, bare bool, cloneOpt *Options, opt *types.Options) (*types.Repository, error) {
	repo, err := with(opt).Start(url, branch, bare, cloneOpt)
	if err != nil {
		return nil, err
	}

	if opt.FollowOpenedRepo {
		opt.WithRepo(repo)
	}

	return repo, err
}

type wrapper interface {
	Start(url string, branch string, bare bool, cloneOpt *Options) (*types.Repository, error)
}

func with(opt *types.Options) wrapper {
	switch opt.PreferredLib {
	case types.PreferGit2Go:
		return &git2go{workdir: opt.WorkDir}
	case types.PreferGoGit:
		fallthrough
	default:
		return &goGit{workdir: opt.WorkDir, auth: &opt.Auth, ctx: context.TODO(), progress: opt.Progress}
	}
}
