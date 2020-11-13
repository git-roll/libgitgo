package libindex

import (
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
  "github.com/go-git/go-git/v5"
)

type goGit struct {
  *types.Options
}

func (g goGit) Add(paths []string) (err error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return
  }

  wt, err := repo.Worktree()
  if err != nil {
    return
  }

  if len(paths) == 1 && paths[0] == "." {
    return wt.AddWithOptions(&git.AddOptions{
      All:  true,
    })
  }

  for _, path := range paths {
    err = wt.AddWithOptions(&git.AddOptions{
      Glob: path,
    })

    if err != nil {
      return
    }
  }

  return
}
