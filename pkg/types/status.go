package types

import "C"
import (
	gogit "github.com/go-git/go-git/v5"
	git2go "github.com/libgit2/git2go/v31"
)

type Git2GoStatus git2go.Status

const (
	StatusCurrent         = git2go.StatusCurrent
	StatusIndexNew        = git2go.StatusIndexNew
	StatusIndexModified   = git2go.StatusIndexModified
	StatusIndexDeleted    = git2go.StatusIndexDeleted
	StatusIndexRenamed    = git2go.StatusIndexRenamed
	StatusIndexTypeChange = git2go.StatusIndexTypeChange
	StatusWtNew           = git2go.StatusWtNew
	StatusWtModified      = git2go.StatusWtModified
	StatusWtDeleted       = git2go.StatusWtDeleted
	StatusWtTypeChange    = git2go.StatusWtTypeChange
	StatusWtRenamed       = git2go.StatusWtRenamed
	StatusIgnored         = git2go.StatusIgnored
	StatusConflicted      = git2go.StatusConflicted
)

type Status struct {
	Git2Go *git2go.StatusList
	GoGit  gogit.Status
}
