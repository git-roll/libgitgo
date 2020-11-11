package libbranch

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

type ListOption struct {
    Git2Go Git2GoListOption
}

type Git2GoListOption struct {
    Type string
}

func List(listOpt *ListOption, opt *types.Options) ([]*types.Branch, error) {
    return with(opt).List(listOpt)
}

type CreateOption struct {
    Git2Go Git2GoCreateOption
    GoGit GoGitCreateOption
}

type Git2GoCreateOption struct {
    Target string
    Force bool
}

type GoGitCreateOption struct {
    Remote string
    Merge string
    Rebase string
}

func Create(name string, createOpt *CreateOption, opt *types.Options) (*types.Branch, error) {
    return with(opt).Create(name, createOpt)
}

type wrapper interface {
    List(*ListOption) ([]*types.Branch, error)
    Create(name string, createOpt *CreateOption) (*types.Branch, error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        return &goGit{opt}
    case types.PreferGit2Go:
        return &git2go{opt}
    }

    return nil
}
