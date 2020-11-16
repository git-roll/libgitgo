package libbranch

import (
  "fmt"
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
  "github.com/go-git/go-git/v5"
  "github.com/go-git/go-git/v5/config"
  "github.com/go-git/go-git/v5/plumbing"
  "golang.org/x/xerrors"
  "strings"
)

type goGit struct {
  *types.Options
}

func (g goGit) Current() (br *types.Branch, err error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return
  }

  head, err := repo.Head()
  if err != nil {
    return
  }

  if !head.Name().IsBranch() {
    return nil, xerrors.Errorf("HEAD is not a branch")
  }

  branch, err := repo.Branch(head.Name().Short())
  if err != nil {
    return
  }

  br = &types.Branch{GoGit: branch}
  return
}

func (g goGit) DeleteAll(brs []string) error {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return err
  }

  for _, br := range brs {
    if err := repo.DeleteBranch(br); err != nil {
      return err
    }
  }

  return nil
}

func (g goGit) Checkout(name string) error {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return err
  }

  return g.checkoutBranch(repo, name, false)
}

func (g goGit) CheckoutNew(name string) (*types.Branch, error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return nil, err
  }

  if err = g.checkoutBranch(repo, name, true); err != nil {
    return nil, err
  }

  br, err := repo.Branch(name)
  if err != nil {
    return nil, err
  }

  return &types.Branch{GoGit: br}, nil
}

func (g goGit) checkoutBranch(repo *git.Repository, name string, create bool) error {
  wt, err := repo.Worktree()
  if err != nil {
    return err
  }

  branchRef := plumbing.NewBranchReferenceName(name)
  _, err = repo.Reference(branchRef, false)
  if err != nil {
    remoteBr := strings.Split(name, "/")
    if len(remoteBr) != 2 {
      return fmt.Errorf("invalid branch %s", name)
    }

    branchRef = plumbing.NewRemoteReferenceName(remoteBr[0], remoteBr[1])
  }

  err = wt.Checkout(&git.CheckoutOptions{
    Branch: branchRef,
    Create: create,
  })

  return err
}

func (g goGit) Create(name string, createOpt *CreateOption) (br *types.Branch, err error) {
  opt := createOpt.GoGit
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

func (g goGit) List(_ *ListOption) (brs []*types.Branch, err error) {
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
