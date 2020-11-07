package remote

import (
    "github.com/git-roll/git-cli/pkg/args"
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/config"
)

type goGit struct {
    worktree string
}

func (g goGit) List() (names []string, err error) {
    repo, err := git.PlainOpen(g.worktree)
    if err != nil {
        return
    }

    remotes, err := repo.Remotes()
    if err != nil {
        return
    }

    for _, remote := range remotes {
        names = append(names, remote.Config().Name)
    }

    return
}

func (g goGit) Create(a args.Map) (err error) {
    repo, err := git.PlainOpen(g.worktree)
    if err != nil {
        return
    }

    conf := &config.RemoteConfig{}

    libArgs := a.Git2GoWrapper()
    url, err := libArgs.MustGet(ParameterKeyURL)
    if err != nil {
        return
    }

    conf.URLs = append(conf.URLs, url)

    fetchSpec := libArgs.MayGet(ParameterKeyFetchSpec)
    if len(fetchSpec) > 0 {
        conf.Fetch = append(conf.Fetch, config.RefSpec(fetchSpec))
    }

    name := libArgs.MayGet(ParameterKeyName)
    if len(name) == 0 {
        _, err = repo.CreateRemoteAnonymous(conf)
    } else {
        _, err = repo.CreateRemote(conf)
    }

    return
}

func (g goGit) Lookup(name string) (args.Object, error) {
    repo, err := git.PlainOpen(g.worktree)
    if err != nil {
        return nil, err
    }

    remote, err := repo.Remote(name)
    if err != nil {
        return nil, err
    }

    return args.Normalize(remote.Config()), err
}
