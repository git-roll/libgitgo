package libconfig

import (
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
  "github.com/go-git/go-git/v5/config"
)

type goGit struct {
  *types.Options
}

func (g goGit) User() (user *types.User, err error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return
  }

  conf, err := repo.ConfigScoped(config.GlobalScope)
  if err != nil {
    return
  }

  user = &types.User{
    Name:  conf.User.Name,
    Email: conf.User.Email,
  }

  return
}
