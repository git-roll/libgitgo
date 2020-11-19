package libbranch

import "C"
import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
)

type ListOption struct {
	Git2Go Git2GoListOption
}

type Git2GoBranchType git.Branch

const (
	BranchAll    = git.BranchAll
	BranchLocal  = git.BranchLocal
	BranchRemote = git.BranchRemote
)

type Git2GoListOption struct {
	Type git.BranchType
}

func List(listOpt *ListOption, opt *types.Options) ([]*types.Branch, error) {
	return with(opt).List(listOpt)
}

type CreateOption struct {
	Git2Go Git2GoCreateOption
	GoGit  GoGitCreateOption
}

type Git2GoCreateOption struct {
	Target string
	Force  bool
}

type GoGitCreateOption struct {
	Remote string
	Merge  string
	Rebase string
}

func Create(name string, createOpt *CreateOption, opt *types.Options) (*types.Branch, error) {
	return with(opt).Create(name, createOpt)
}

func Checkout(name string, opt *types.Options) error {
	return with(opt).Checkout(name)
}

func CheckoutNew(name string, opt *types.Options) (*types.Branch, error) {
	return with(opt).CheckoutNew(name)
}

func Current(opt *types.Options) (*types.Branch, error) {
	return with(opt).Current()
}

func Delete(name string, opt *types.Options) error {
	return with(opt).DeleteAll([]string{name})
}

func DeleteAll(branches []string, opt *types.Options) error {
	return with(opt).DeleteAll(branches)
}

func BranchesHaveMergedTo(name, remote string, opt *types.Options) ([]*types.Branch, error) {
	return with(opt).BranchesHaveMergedTo(name, remote)
}

func Get(name, remote string, opt *types.Options) (*types.Branch, error) {
	return with(opt).Current()
}

type wrapper interface {
	List(*ListOption) ([]*types.Branch, error)
	Create(name string, createOpt *CreateOption) (*types.Branch, error)
	Checkout(name string) error
	CheckoutNew(name string) (*types.Branch, error)
	Current() (*types.Branch, error)
	DeleteAll(brs []string) error
	BranchesHaveMergedTo(name, remote string) ([]*types.Branch, error)
	Get(name, remote string) (*types.Branch, error)
}

func with(opt *types.Options) wrapper {
	switch opt.PreferredLib {
	case types.PreferGoGit:
		return &goGit{opt}
	case types.PreferGit2Go:
		fallthrough
	default:
		return &git2go{opt}
	}
}
