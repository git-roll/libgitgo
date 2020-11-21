package types

import (
    "github.com/go-git/go-git/v5/plumbing/object"
    git "github.com/libgit2/git2go/v31"
)

type Commit struct {
    GoGit *object.Commit
    Git2Go *git.Commit
}
