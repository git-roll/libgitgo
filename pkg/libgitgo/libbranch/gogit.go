package libbranch

import (
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
  "github.com/go-git/go-git/v5"
  "github.com/go-git/go-git/v5/config"
  "github.com/go-git/go-git/v5/plumbing"
)

type goGit struct {
  *types.Options
}

func (g goGit) Create(name string, _ *Git2GoCreateOption, opt *GoGitCreateOption) (br *types.Branch, err error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return nil, err
  }

  br = &types.Branch{GoGit: &config.Branch{
    Name:   name,
    Remote: opt.Remote,
    Merge:  plumbing.ReferenceName(opt.Merge),
    Rebase: opt.Rebase,
  }}

  err = repo.CreateBranch(br.GoGit)
  return
}

func (g goGit) List(_ *Git2GoListOption) (brs []*types.Branch, err error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return nil, err
  }

  // gogit filters all branch-typed refs.
  it, err := repo.Branches()
  if err != nil {
    return nil, err
  }

  err = it.ForEach(func(ref *plumbing.Reference) error{
    // When lookup branch, gogit search branches in .git/config.
    // Since most branches created by other git clients would not state in the file, the r.Branch almost fails always.
    br, err := repo.Branch(ref.Name().Short())
    if err != nil {
      if err == git.ErrBranchNotFound {
        brs = append(brs, &types.Branch{GoGit: &config.Branch{Name: ref.Name().String()}})
        return nil
      }

      panic(err)
    }

    brs = append(brs, &types.Branch{GoGit: br})
    return nil
  })

  return
}
