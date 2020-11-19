package types

import (
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	git2go "github.com/libgit2/git2go/v31"
)

type Repository struct {
	Git2Go *git2go.Repository
	GoGit  *gogit.Repository
}

type HEAD struct {
	Git2Go *git2go.Reference
	GoGit  *plumbing.Reference
}
