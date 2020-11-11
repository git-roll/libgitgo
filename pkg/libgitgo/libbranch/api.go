package libbranch

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

type Git2GoListOption struct {
    Type string
}

func List(listOpt *Git2GoListOption, opt *types.Options) ([]*types.Branch, error) {
    return with(opt).List(listOpt)
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

func Create(name string, createOpt1 *Git2GoCreateOption, createOpt2 *GoGitCreateOption, opt *types.Options) (*types.Branch, error) {
    return with(opt).Create(name, createOpt1, createOpt2)
}

type wrapper interface {
    List(*Git2GoListOption) ([]*types.Branch, error)
    Create(name string, createOpt1 *Git2GoCreateOption, createOpt2 *GoGitCreateOption) (*types.Branch, error)
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
