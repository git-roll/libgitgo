package types

import (
    "github.com/go-git/go-git/v5/plumbing"
    git "github.com/libgit2/git2go/v31"
)

type Reference struct {
    GoGit *plumbing.Reference
    Git2Go *git.Reference
}
