package types

import (
	"context"
	"fmt"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	gitgo "github.com/libgit2/git2go/v31"
	sshCli "golang.org/x/crypto/ssh"
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
	Context context.Context
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

func (opt *Options) WithContext(ctx context.Context) *Options {
	newOpt := *opt
	newOpt.Context = ctx
	return &newOpt
}

type Auth struct {
	User       string
	Password   string
	SSHId      string
	Passphrase string
}

func (a Auth) GenGoGitAuth(url string) (transport.AuthMethod, error) {
	ep, err := transport.NewEndpoint(url)
	if err != nil {
		return nil, err
	}

	switch ep.Protocol {
	case "ssh":
		if len(ep.User) > 0 && len(ep.Password) > 0 {
			return &ssh.Password{
				User:                  ep.User,
				Password:              ep.Password,
				HostKeyCallbackHelper: ssh.HostKeyCallbackHelper{HostKeyCallback: sshCli.InsecureIgnoreHostKey()},
			}, nil
		} else if len(a.User) > 0 {
			return &ssh.Password{
				User:                  a.User,
				Password:              a.Password,
				HostKeyCallbackHelper: ssh.HostKeyCallbackHelper{HostKeyCallback: sshCli.InsecureIgnoreHostKey()},
			}, nil
		} else if len(a.SSHId) > 0 {
			return ssh.NewPublicKeysFromFile(ssh.DefaultUsername, a.SSHId, a.Passphrase)
		}
	case "file", "http", "https":
		if len(ep.User) > 0 && len(ep.Password) > 0 {
			return &http.BasicAuth{
				Username: ep.User,
				Password: ep.Password,
			}, nil
		}
	default:
		return nil, fmt.Errorf("unsupported protocol %s", ep.Protocol)
	}

	return nil, nil
}

func (a Auth) GenGit2GoAuth(url string, username string, allowedTypes gitgo.CredType) (*gitgo.Cred, error) {
	if (allowedTypes & gitgo.CredTypeUserpassPlaintext) > 0 {
		if len(username) > 0 && username != a.User {
			return nil, fmt.Errorf("username not matched")
		}

		return gitgo.NewCredUserpassPlaintext(a.User, a.Password)
	}

	if (allowedTypes & gitgo.CredTypeSshKey) > 0 {
		return gitgo.NewCredSshKey(username, a.SSHId+".pub", a.SSHId, a.Passphrase)
	}

	if (allowedTypes & gitgo.CredTypeDefault) > 0 {
		return gitgo.NewCredDefault()
	}

	return nil, nil
}
