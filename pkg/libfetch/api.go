package libfetch

import (
	"github.com/git-roll/libgitgo/pkg/types"
	"golang.org/x/xerrors"
)

var (
	ErrUpToDate = xerrors.Errorf("up-to-date")
)

type Options struct {
	Git2Go Git2GoOptions
	GoGit  GoGitOptions
}

type Git2GoOptions struct {
	DownloadTags types.Git2GoDownloadTags
}

type GoGitOptions struct {
	Depth   int
	TagMode types.GoGitTagMode
}

func Remote(remote string, fetchOpt *Options, opt *types.Options) (err error) {
	return with(opt).Start("", remote, fetchOpt)
}

func Branch(branch, remote string, fetchOpt *Options, opt *types.Options) (err error) {
	return with(opt).Start(branch, remote, fetchOpt)
}

type wrapper interface {
	Start(branch, remote string, fetchOpt *Options) (err error)
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
