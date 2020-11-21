package libclone

import (
	"context"
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	gogit "github.com/go-git/go-git/v5"
	gitgo "github.com/libgit2/git2go/v31"
)

type Options struct {
	Git2Go Git2GoOptions
	GoGit  GoGitOptions
}

type Git2GoDownloadTags gitgo.DownloadTags

const (
	DownloadTagsUnspecified = Git2GoDownloadTags(gitgo.DownloadTagsUnspecified)
	DownloadTagsAuto        = Git2GoDownloadTags(gitgo.DownloadTagsAuto)
	DownloadTagsNone        = Git2GoDownloadTags(gitgo.DownloadTagsNone)
	DownloadTagsAll         = Git2GoDownloadTags(gitgo.DownloadTagsAll)
)

type Git2GoOptions struct {
	DownloadTags Git2GoDownloadTags
}

type GoGitSubmoduleRescursivity gogit.SubmoduleRescursivity

const (
	NoRecurseSubmodules            = GoGitSubmoduleRescursivity(gogit.NoRecurseSubmodules)
	DefaultSubmoduleRecursionDepth = GoGitSubmoduleRescursivity(gogit.DefaultSubmoduleRecursionDepth)
)

type GoGitTagMode gogit.TagMode

const (
	InvalidTagMode = GoGitTagMode(gogit.InvalidTagMode)
	TagFollowing   = GoGitTagMode(gogit.TagFollowing)
	AllTags        = GoGitTagMode(gogit.AllTags)
	NoTags         = GoGitTagMode(gogit.NoTags)
)

type GoGitOptions struct {
	// usually `origin`
	RemoteName string
	// specified branch only
	SingleBranch          bool
	NoCheckout            bool
	Depth                 int
	SubmoduleRescursivity GoGitSubmoduleRescursivity
	TagMode               GoGitTagMode
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
