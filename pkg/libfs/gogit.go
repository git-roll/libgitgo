package libfs

import (
  "github.com/git-roll/libgitgo/pkg/types"
  "io/ioutil"
  "os"
)

type goGit struct {
  *types.Options
}

func (g goGit) WriteFile(name, content string) (err error) {
  repo, err := g.OpenGoGitRepo()
  if err != nil {
    return
  }

  wt, err := repo.Worktree()
  if err != nil {
    return
  }

  f, err := wt.Filesystem.Create(name)
  if err != nil {
    return
  }

  n, err := f.Write([]byte(content))
  if err != nil {
    return
  }

  if n != len(content) {
    panic(n)
  }

  return
}

func (g goGit) Stat(name string) (os.FileInfo, error) {
  repo, err := g.OpenGoGitRepo()
  if err != nil {
    return nil, err
  }

  wt, err := repo.Worktree()
  if err != nil {
    return nil, err
  }

  return wt.Filesystem.Stat(name)
}

func (g goGit) ReadFile(name string) ([]byte, error) {
  repo, err := g.OpenGoGitRepo()
  if err != nil {
    return nil, err
  }

  wt, err := repo.Worktree()
  if err != nil {
    return nil, err
  }

  f, err := wt.Filesystem.Open(name)
  if err != nil {
    return nil, err
  }

  defer func() {
    if err := f.Close(); err != nil {
      panic(err)
    }
  }()

  return ioutil.ReadAll(f)
}
