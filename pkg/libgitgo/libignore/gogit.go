package libignore

import (
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
  "github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

type goGit struct {
  *types.Options
}

func (g goGit) Check(relativePath string) (err error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return
  }

  wt, err := repo.Worktree()
  if err != nil {
    return
  }

  fi, err := wt.Filesystem.Stat(relativePath)
  if err != nil {
    return
  }

  patterns, err := gitignore.ReadPatterns(wt.Filesystem, nil)
  if err != nil {
    return
  }

  m := gitignore.NewMatcher(patterns)

  if m.Match([]string{relativePath}, fi.IsDir()) {
    return ErrPathIgnored
  }

  return
}
