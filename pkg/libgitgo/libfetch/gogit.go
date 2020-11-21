package libfetch

import (
	"fmt"
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

type goGit struct {
	*types.Options
}

func (g goGit) Start(branch, remoteName string, fetchOpt *Options) (err error) {
	repo, err := g.OpenGoGitRepo()
	if err != nil {
		return
	}

	remote, err := repo.Remote(remoteName)
	if err != nil {
		return
	}

	var refs []config.RefSpec
	if len(branch) == 0 {
		refs = remote.Config().Fetch
	} else {
		for _, fetch := range remote.Config().Fetch {
			if fetch.Match(plumbing.NewBranchReferenceName(branch)) {
				refs = append(refs, fetch)
				break
			}
		}

		if len(refs) == 0 {
			return fmt.Errorf("no refspec matches branch %s", branch)
		}
	}

	auth, err := g.Options.Auth.GenGoGitAuth(remote.Config().URLs[0])
	if err != nil {
		return
	}

	err = repo.Fetch(&git.FetchOptions{
		RemoteName: remoteName,
		RefSpecs:   refs,
		Auth:       auth,
		Progress:   g.Progress,
		Tags:       fetchOpt.GoGit.TagMode,
		Depth:      fetchOpt.GoGit.Depth,
	})

	if err == git.NoErrAlreadyUpToDate {
		err = ErrUpToDate
	}

	return
}
