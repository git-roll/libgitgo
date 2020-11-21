package libref

import (
  "github.com/git-roll/libgitgo/pkg/types"
  "io"
)

type goGit struct {
  *types.Options
}

func (g goGit) List() (refs []*types.Reference, err error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return
  }

  i, err := repo.References()
  if err != nil {
    return
  }

  for ref, err := i.Next(); err == nil; ref, err = i.Next()  {
    refs = append(refs, &types.Reference{GoGit: ref})
  }

  if err == io.EOF {
    err = nil
  }

  return
}
