package libcommit

import (
	"github.com/git-roll/libgitgo/pkg/types"
	git "github.com/libgit2/git2go/v31"
	"golang.org/x/xerrors"
	"time"
)

type git2go struct {
	*types.Options
}

func (g git2go) IsAncestor(ancestor, second string) (positive bool, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	ancestorRef, err := repo.References.Lookup(ancestor)
	if err != nil {
		return
	}

	secondRef, err := repo.References.Lookup(second)
	if err != nil {
		return
	}

	ancestorOid, err := repo.MergeBase(ancestorRef.Target(), secondRef.Target())
	if err != nil {
		return
	}

	positive = ancestorOid.Equal(ancestorRef.Target())
	return
}

func (g git2go) Amend(message string, opt *CommitOptions) (commit *types.Commit, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	head, err := repo.Head()
	if err != nil {
		return
	}

	headCommit, err := repo.LookupCommit(head.Target())
	if err != nil {
		return
	}

	if headCommit.ParentCount() > 1 {
		return nil, xerrors.Errorf("can't amend merge commits")
	}

	headTree, err := headCommit.Tree()
	if err != nil {
		return
	}

	now := time.Now()
	commitId, err := headCommit.Amend("", &git.Signature{
		Name:  opt.Author.Name,
		Email: opt.Author.Email,
		When:  now,
	}, &git.Signature{
		Name:  opt.Committer.Name,
		Email: opt.Committer.Email,
		When:  now,
	}, message, headTree)
	if err != nil {
		return
	}

	ci, err := repo.LookupCommit(commitId)
	if err != nil {
		panic(err)
	}
	commit = &types.Commit{Git2Go: ci}
	return
}

func (g git2go) CommitStaging(message string, opt *CommitOptions) (commit *types.Commit, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	statusList, err := repo.StatusList(&git.StatusOptions{})
	if err != nil {
		return
	}

	statusCount, err := statusList.EntryCount()
	if err != nil {
		return
	}

	if statusCount == 0 {
		return nil, nil
	}

	index, err := repo.Index()
	if err != nil {
		return
	}

	var modified, deleted, added []string
	for i := 0; i < statusCount; i++ {
		entry, err := statusList.ByIndex(i)
		if err != nil {
			return nil, err
		}

		switch entry.Status {
		case git.StatusWtModified, git.StatusWtTypeChange:
			modified = append(modified, entry.HeadToIndex.NewFile.Path)
		case git.StatusWtDeleted:
			deleted = append(deleted, entry.HeadToIndex.OldFile.Path)
		case git.StatusWtRenamed:
			added = append(added, entry.HeadToIndex.NewFile.Path)
			deleted = append(deleted, entry.HeadToIndex.OldFile.Path)
		}
	}

	if len(modified) > 0 {
		err = index.UpdateAll(modified, nil)
		if err != nil {
			return
		}
	}

	if len(deleted) > 0 {
		err = index.RemoveAll(deleted, nil)
		if err != nil {
			return
		}
	}

	if len(added) > 0 {
		err = index.AddAll(added, git.IndexAddDefault, nil)
		if err != nil {
			return
		}
	}

	treeOid, err := index.WriteTree()
	if err != nil {
		return
	}

	if err = index.Write(); err != nil {
		return
	}

	var parent *git.Reference
	if len(opt.Git2Go.Parent) > 0 {
		if parent, err = repo.References.Lookup(opt.Git2Go.Parent); err != nil {
			return
		}
	} else {
		if parent, err = repo.Head(); err != nil {
			return
		}
	}

	now := time.Now()
	oid, err := repo.CreateCommitFromIds(opt.Git2Go.RefName, &git.Signature{
		Name:  opt.Author.Name,
		Email: opt.Author.Email,
		When:  now,
	}, &git.Signature{
		Name:  opt.Committer.Name,
		Email: opt.Committer.Email,
		When:  now,
	}, message, treeOid, parent.Target())
	if err != nil {
		return
	}

	obj, err := repo.LookupCommit(oid)
	if err != nil {
		panic(err)
	}

	commit = &types.Commit{Git2Go: obj}
	return
}

func (g git2go) Get(refName string) (commit *types.Commit, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	ref, err := repo.References.Lookup(refName)
	if err != nil {
		return
	}

	obj, err := repo.LookupCommit(ref.Target())
	if err != nil {
		return
	}

	commit = &types.Commit{Git2Go: obj}
	return
}
