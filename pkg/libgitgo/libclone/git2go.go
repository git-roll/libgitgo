package libclone

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	workdir string
}

func (g git2go) Start(url string, branch string, bare bool, opt *Options) (*types.Repository, error) {
	r, err := git.Clone(url, g.workdir, &git.CloneOptions{
		FetchOptions:         &git.FetchOptions{
			RemoteCallbacks: git.RemoteCallbacks{
				CredentialsCallback: func(url string, username_from_url string, allowed_types git.CredType) (*git.Cred, error) {
					return git.NewCredDefault()
				},
			},
			UpdateFetchhead: true,
			DownloadTags:    git.DownloadTags(opt.Git2Go.DownloadTags),
		},
		Bare:                 bare,
		CheckoutBranch:       branch,
		RemoteCreateCallback: nil,
	})

	if err != nil {
		return nil, err
	}

	return &types.Repository{Git2Go: r}, nil
}
