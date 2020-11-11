package libremote

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type goGit struct {
	*types.Options
}

func (g goGit) List() (remotes []*types.Remote, err error) {
	repo, err := g.Options.OpenGoGitRepo()
	if err != nil {
		return
	}

	remoteObjs, err := repo.Remotes()
	if err != nil {
		return
	}

	for _, remote := range remoteObjs {
		remotes = append(remotes, &types.Remote{GoGit: remote})
	}

	return
}

func (g goGit) Create(name, url, fetchSpec string) (remote *types.Remote, err error) {
	repo, err := g.Options.OpenGoGitRepo()
	if err != nil {
		return
	}

	if len(url) == 0 {
		panic(name)
	}

	conf := &config.RemoteConfig{}
	conf.URLs = append(conf.URLs, url)

	if len(fetchSpec) > 0 {
		conf.Fetch = append(conf.Fetch, config.RefSpec(fetchSpec))
	}

	var r *git.Remote
	if len(name) == 0 {
		r, err = repo.CreateRemoteAnonymous(conf)
	} else {
		r, err = repo.CreateRemote(conf)
	}

	if err != nil {
		return
	}

	return &types.Remote{GoGit: r}, err
}

func (g goGit) Lookup(name string) (remote *types.Remote, err error) {
	repo, err := g.Options.OpenGoGitRepo()
	if err != nil {
		return nil, err
	}

	r, err := repo.Remote(name)
	if err != nil {
		return nil, err
	}

	return &types.Remote{GoGit: r}, err
}
