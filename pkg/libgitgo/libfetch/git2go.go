package libfetch

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	"github.com/git-roll/libgitgo/pkg/refspec"
	git "github.com/libgit2/git2go/v31"
	"os"
	"path/filepath"
)

type git2go struct {
	*types.Options
}

func (g git2go) Start(branch, remoteName string, fetchOpt *Options) (err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	remote, err := repo.Remotes.Lookup(remoteName)
	if err != nil {
		return
	}

	refs, err := remote.FetchRefspecs()
	if err != nil {
		return err
	}

	if len(branch) > 0 {
		refs = []string{refspec.FetchBranch(branch, remoteName)}
	}

	fetchHead := filepath.Join(repo.Workdir(), ".git", "FETCH_HEAD")
	fetchHeadSt, _ := os.Lstat(fetchHead)

	err = remote.Fetch(refs, &git.FetchOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			CredentialsCallback:          g.GenGit2GoAuth,
			CertificateCheckCallback:     func(cert *git.Certificate, valid bool, hostname string) git.ErrorCode{
				return git.ErrOk
			},
		},
		UpdateFetchhead: true,
		DownloadTags:    fetchOpt.Git2Go.DownloadTags,
	}, "")

	if err != nil {
		return
	}

	newFetchHeadSt, err := os.Lstat(fetchHead)
	if err != nil {
		return
	}

	if fetchHeadSt != nil && fetchHeadSt.ModTime() == newFetchHeadSt.ModTime() {
		err = ErrUpToDate
	}

	return
}
