package libbranch

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
	"golang.org/x/xerrors"
)

type git2go struct {
	*types.Options
}

func (g git2go) Get(name, remote string) (br *types.Branch, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	brName := name
	if len(remote) > 0 {
		brName = remote + "/" + name
	}

	branch, err := repo.LookupBranch(brName, git.BranchAll)
	if err != nil {
		return
	}

	return &types.Branch{Git2Go: branch}, nil
}

func (g git2go) BranchesHaveMergedTo(name, remote string) (brs []*types.Branch, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	targetBr, err := repo.LookupBranch(name, git.BranchAll)
	if err != nil {
		return
	}

	bri, err := repo.NewBranchIterator(git.BranchLocal)
	if err != nil {
		return
	}

	err = bri.ForEach(func(br *git.Branch, _ git.BranchType) error{
		brName, err := br.Name()
		if err != nil {
			return err
		}

		if brName == name {
			return nil
		}

		base, err := repo.MergeBase(br.Target(), targetBr.Target())
		if err != nil {
			return err
		}

		if base == br.Target() {
			brs = append(brs, &types.Branch{Git2Go: br})
		}

		return nil
	})

	return
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

func (g git2go) DeleteAll(brs []string) error {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return err
	}

	for _, br := range brs {
		branch, err := repo.LookupBranch(br, git.BranchLocal)
		if err != nil {
			return err
		}

		if err = branch.Delete(); err != nil {
			return err
		}
	}


	return nil
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

	return g.checkoutBranch(repo, br)
}

func (g git2go) CheckoutNew(name string) (*types.Branch, error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return nil, err
	}

	br, err := g.createBranch(repo, name, &Git2GoCreateOption{})
	if err != nil {
		return nil, err
	}

	err = g.checkoutBranch(repo, br)
	if err != nil {
		return nil, err
	}

	return &types.Branch{Git2Go: br}, nil
}

func (g git2go) Create(name string, createOpt *CreateOption) (*types.Branch, error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return nil, err
	}

	br, err := g.createBranch(repo, name, &createOpt.Git2Go)
	if err != nil {
		return nil, err
	}
	return &types.Branch{Git2Go: br}, nil
}

func (g git2go) checkoutBranch(repo *git.Repository, br *git.Branch) error {
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

	brName, err := br.Name()
	if err != nil {
		return err
	}

	return repo.SetHead(brName)
}

func (g git2go) createBranch(repo *git.Repository, name string, opt *Git2GoCreateOption) (*git.Branch, error) {
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

	return repo.CreateBranch(name, targetCommit, opt.Force)
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
