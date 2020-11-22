package types

import (
    "github.com/go-git/go-git/v5/plumbing/object"
    git2go "github.com/libgit2/git2go/v31"
)

type Diff struct {
    Git2Go *git2go.Diff
    GoGit object.Changes
}

func (r Diff) String() string {
  if r.Git2Go != nil {
    if r.GoGit != nil {
      panic("both pointers are filled")
    }

    buf, err := r.Git2Go.ToBuf(git2go.DiffFormatPatch)
    if err != nil {
        panic(err)
    }

    return string(buf)
  }

  if r.GoGit != nil {
    if r.Git2Go != nil {
      panic("both pointers are filled")
    }
  }

  panic("both pointers are nil")
}
