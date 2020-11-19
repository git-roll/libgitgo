package libstatus

import (
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

type goGit struct {
  *types.Options
}

func (g goGit) List() (list *types.Status, err error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return
  }

  wt, err := repo.Worktree()
  if err != nil {
    return
  }

  st, err := wt.Status()
  if err != nil {
    return
  }

  list = &types.Status{GoGit: st}
  return
}
