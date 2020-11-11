package libbranch

import (
	"fmt"
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	*types.Options
}

func (g git2go) Create(name string, opt *Git2GoCreateOption, _ *GoGitCreateOption) (*types.Branch, error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return nil, err
	}

	commitOid, err := git.NewOid(opt.Target)
	if err != nil {
		target, err := repo.References.Lookup(opt.Target)
		if err != nil {
			return nil, err
		}

		targetRef, err := target.Resolve()
		if err != nil {
			return nil, err
		}

		commitOid = targetRef.Target()
	}

	targetCommit, err := repo.LookupCommit(commitOid)
	if err != nil {
		return nil, err
	}

	br, err := repo.CreateBranch(name, targetCommit, opt.Force)
	if err != nil {
		return nil, err
	}

	return &types.Branch{Git2Go: br}, nil
}

func (g git2go) List(opt *Git2GoListOption) (brs []*types.Branch, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return nil, err
	}

	var brType git.BranchType
	switch opt.Type {
	case "BranchLocal":
		brType = git.BranchLocal
	case "BranchRemote":
		brType = git.BranchRemote
	case "BranchAll", "":
		brType = git.BranchAll
	default:
		return nil, fmt.Errorf(`BranchType could be one of "BranchLocal", "BranchRemote", or "BranchAll"`)
	}

	it, err := repo.NewBranchIterator(brType)
	if err != nil {
		return nil, err
	}

	err = it.ForEach(func(br *git.Branch, _ git.BranchType) error {
		brs = append(brs, &types.Branch{Git2Go: br})
		return nil
	})
	return
}
