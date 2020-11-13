package libcommit

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
	"time"
)

type git2go struct {
	*types.Options
}

func (g git2go) CommitStaging(message string, opt *CommitOptions) (commit *types.Commit, err error) {
	repo, err := g.Options.OpenGit2GoRepo()
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
		Name:  opt.Author.Name,
		Email: opt.Author.Email,
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
	repo, err := g.Options.OpenGit2GoRepo()
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
