package types

import (
	gogit "github.com/go-git/go-git/v5"
	gitgo "github.com/libgit2/git2go/v31"
	"io"
)

type Options struct {
	Progress io.Writer
	WorkDir string
	PreferredLib
	Auth

	// Opened Repository will be saved in the Options if enabled
	FollowOpenedRepo bool
	CtxRepo          *Repository
}

func (opt *Options) WithRepo(repo *Repository) {
	if opt.CtxRepo != nil {
		panic("an opened repo existed")
	}

	if opt.FollowOpenedRepo {
		opt.CtxRepo = repo
	}
}

func (opt *Options) OpenGoGitRepo() (repo *gogit.Repository, err error) {
	if opt.CtxRepo != nil && opt.CtxRepo.GoGit != nil {
		repo = opt.CtxRepo.GoGit
		return
	}

	repo, err = gogit.PlainOpen(opt.WorkDir)
	if err != nil {
		return nil, err
	}

	if opt.CtxRepo != nil {
		opt.CtxRepo.GoGit = repo
		return
	}

	opt.WithRepo(&Repository{GoGit: repo})
	return
}

func (opt *Options) OpenGit2GoRepo() (repo *gitgo.Repository, err error) {
	if opt.CtxRepo != nil && opt.CtxRepo.Git2Go != nil {
		repo = opt.CtxRepo.Git2Go
		return
	}

	repo, err = gitgo.OpenRepository(opt.WorkDir)
	if err != nil {
		return nil, err
	}

	if opt.CtxRepo != nil {
		opt.CtxRepo.Git2Go = repo
		return
	}

	opt.WithRepo(&Repository{Git2Go: repo})
	return
}

type Auth struct {
	User       string
	Password   string
	SSHId      string
	Passphrase string
}
