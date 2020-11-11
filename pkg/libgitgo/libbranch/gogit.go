package libbranch

import (
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
  "github.com/go-git/go-git/v5"
  "github.com/go-git/go-git/v5/config"
  "github.com/go-git/go-git/v5/plumbing"
)

type goGit struct {
  workdir string
}

func (g goGit) Create(name string, _ *Git2GoCreateOption, opt *GoGitCreateOption) (br *types.Branch, err error) {
  r, err := git.PlainOpen(g.workdir)
  if err != nil {
    return nil, err
  }

  br = &types.Branch{GoGit: &config.Branch{
    Name:   name,
    Remote: opt.Remote,
    Merge:  plumbing.ReferenceName(opt.Merge),
    Rebase: opt.Rebase,
  }}

  err = r.CreateBranch(br.GoGit)
  return
}

func (g goGit) List(_ *Git2GoListOption) (brs []*types.Branch, err error) {
  r, err := git.PlainOpen(g.workdir)
  if err != nil {
    return nil, err
  }

  // gogit filters all branch-typed refs.
  it, err := r.Branches()
  if err != nil {
    return nil, err
  }

  err = it.ForEach(func(ref *plumbing.Reference) error{
    // When lookup branch, gogit search branches in .git/config.
    // Since most branches created by other git clients would not state in the file, the r.Branch almost fails always.
    br, err := r.Branch(ref.Name().Short())
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
