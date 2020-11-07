package remote

import (
	"github.com/git-roll/git-cli/pkg/args"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	worktree string
}

func (g git2go) Create(p args.Map) (err error) {
	libArgs := p.Git2GoWrapper()
	repo, err := git.OpenRepository(g.worktree)
	if err != nil {
		return
	}

	url, err := libArgs.MustGet(ParameterKeyURL)
	if err != nil {
		return
	}

	fetchSpec := libArgs.MayGet(ParameterKeyFetchSpec)
	if len(fetchSpec) > 0 {
		name, err := libArgs.MustGet(ParameterKeyName)
		if err != nil {
			return err
		}

		_, err = repo.Remotes.CreateWithFetchspec(name, url, fetchSpec)
		return err
	}

	name := libArgs.MayGet(ParameterKeyName)
	if len(name) == 0 {
		_, err = repo.Remotes.CreateAnonymous(url)
	} else {
		_, err = repo.Remotes.Create(name, url)
	}

	return
}

func (g git2go) List() (names []string, err error) {
	repo, err := git.OpenRepository(g.worktree)
	if err != nil {
		return
	}

	return repo.Remotes.List()
}

func (g git2go) Lookup(name string) (obj args.Object, err error) {
	repo, err := git.OpenRepository(g.worktree)
	if err != nil {
		return
	}

	remote, err := repo.Remotes.Lookup(name)
	if err != nil {
		return
	}

	obj = make(args.Object)
	obj["Name()"] = remote.Name()
	obj["Url()"] = remote.Url()
	fetchSpecs, err := remote.FetchRefspecs()
	if err != nil {
		return
	}

	obj["FetchRefspecs()"] = fetchSpecs
	return
}
