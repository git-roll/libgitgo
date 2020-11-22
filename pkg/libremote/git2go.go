package libremote

import (
    "github.com/git-roll/libgitgo/pkg/types"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	*types.Options
}

func (g git2go) Create(name, url, fetchSpec string) (remote *types.Remote, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	if len(url) == 0 {
		panic(name)
	}

	var r *git.Remote
	if len(fetchSpec) > 0 {
		if len(name) == 0 {
			panic(fetchSpec)
		}

		r, err = repo.Remotes.CreateWithFetchspec(name, url, fetchSpec)
	} else {
		if len(name) == 0 {
			r, err = repo.Remotes.CreateAnonymous(url)
		} else {
			r, err = repo.Remotes.Create(name, url)
		}
	}

	if err != nil {
		return
	}

	return &types.Remote{Git2Go: r}, err
}

func (g git2go) List() (remotes []*types.Remote, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	names, err := repo.Remotes.List()
	if err != nil {
		return
	}

	for _, name := range names {
		remote, err := repo.Remotes.Lookup(name)
		if err != nil {
			return nil, err
		}

		remotes = append(remotes, &types.Remote{Git2Go: remote})
	}

	return
}

func (g git2go) Lookup(name string) (remote *types.Remote, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	r, err := repo.Remotes.Lookup(name)
	if err != nil {
		return
	}

	remote = &types.Remote{Git2Go: r}
	return
}
