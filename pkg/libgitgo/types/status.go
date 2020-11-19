package types

import (
    gogit "github.com/go-git/go-git/v5"
    git2go "github.com/libgit2/git2go/v31"
)

type Status struct {
    Git2Go *git2go.StatusList
    GoGit gogit.Status
}
