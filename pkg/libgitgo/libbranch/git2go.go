package libbranch

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
	"golang.org/x/xerrors"
)

type git2go struct {
	*types.Options
}

func (g git2go) Current() (br *types.Branch, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	head, err := repo.Head()
	if err != nil {
		return
	}

	if !head.IsBranch() {
		return nil, xerrors.Errorf("HEAD is not a branch")
	}

	br = &types.Branch{Git2Go: head.Branch()}
	return
}

func (g git2go) Delete(name string) error {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return err
	}

	br, err := repo.LookupBranch(name, git.BranchLocal)
	if err != nil {
		return err
	}

	return br.Delete()
}

func (g git2go) Checkout(name string) error {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return err
	}

	br, err := repo.LookupBranch(name, git.BranchAll)
	if err != nil {
		return err
	}

	commit, err := repo.LookupCommit(br.Target())
	if err != nil {
		return err
	}

	tree, err := commit.Tree()
	if err != nil {
		return err
	}

	err = repo.CheckoutTree(tree, &git.CheckoutOpts{
		Strategy:         git.CheckoutSafe | git.CheckoutRecreateMissing,
	})

	if err != nil {
		return err
	}

	ref, err := repo.References.Dwim(name)
	if err != nil {
		return err
	}

	return repo.SetHead(ref.Name())
}

func (g git2go) Create(name string, createOpt *CreateOption) (*types.Branch, error) {
	opt := createOpt.Git2Go
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

func (g git2go) List(opt *ListOption) (brs []*types.Branch, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return nil, err
	}

	it, err := repo.NewBranchIterator(opt.Git2Go.Type)
	if err != nil {
		return nil, err
	}

	err = it.ForEach(func(br *git.Branch, _ git.BranchType) error {
		brs = append(brs, &types.Branch{Git2Go: br})
		return nil
	})
	return
}
