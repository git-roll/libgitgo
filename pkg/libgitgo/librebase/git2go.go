package librebase

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
	"strings"
	"time"
)

type git2go struct {
	*types.Options
}

func (g git2go) Start(branch, upstream, onto string, opt *RebaseOptions) (err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	brRef, err := repo.References.Lookup(branch)
	if err != nil {
		return
	}

	brCommit, err := repo.AnnotatedCommitFromRef(brRef)
	if err != nil {
		return
	}

	upstreamRef, err := repo.References.Lookup(upstream)
	if err != nil {
		return
	}

	upstreamCommit, err := repo.AnnotatedCommitFromRef(upstreamRef)
	if err != nil {
		return
	}

	rebaseOpt, err := git.DefaultRebaseOptions()
	if err != nil {
		return
	}

	var ontoCommit *git.AnnotatedCommit
	if len(onto) > 0 {
		ontoRef, err := repo.References.Lookup(onto)
		if err != nil {
			return err
		}

		ontoCommit, err = repo.AnnotatedCommitFromRef(ontoRef)
		if err != nil {
			return err
		}
	}


	rebase, err := repo.InitRebase(brCommit, upstreamCommit, ontoCommit, &rebaseOpt)
	if err != nil {
		return
	}

	now := time.Now()
	author := &git.Signature{
		Name:  opt.Author.Name,
		Email: opt.Author.Email,
		When:  now,
	}

	committer := &git.Signature{
		Name:  opt.Committer.Name,
		Email: opt.Committer.Email,
		When:  now,
	}

	opCount := int(rebase.OperationCount())
	for op := 0; op < opCount; op++ {
		operation, err := rebase.Next()
		if err != nil {
			return err
		}

		commit, err := repo.LookupCommit(operation.Id)
		if err != nil {
			return err
		}

		if err = rebase.Commit(operation.Id, author, committer, commit.Message()); err != nil {
			return err
		}
	}

	return
}

func (g git2go) Abort() (err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	rebaseOpt, err := git.DefaultRebaseOptions()
	if err != nil {
		return
	}

	rebase, err := repo.OpenRebase(&rebaseOpt)
	if err != nil {
		return
	}

	return rebase.Abort()
}

func (g git2go) CompactPrivateCommits(upstream, messagePrefix string, opt *RebaseOptions) (err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	brRef, err := repo.Head()
	if err != nil {
		return
	}

	brCommit, err := repo.AnnotatedCommitFromRef(brRef)
	if err != nil {
		return
	}

	upstreamRef, err := repo.References.Lookup(upstream)
	if err != nil {
		return
	}

	upstreamCommit, err := repo.AnnotatedCommitFromRef(upstreamRef)
	if err != nil {
		return
	}

	rebaseOpt, err := git.DefaultRebaseOptions()
	if err != nil {
		return
	}

	rebase, err := repo.InitRebase(brCommit, upstreamCommit, nil, &rebaseOpt)
	if err != nil {
		return
	}

	now := time.Now()
	author := &git.Signature{
		Name:  opt.Author.Name,
		Email: opt.Author.Email,
		When:  now,
	}

	committer := &git.Signature{
		Name:  opt.Committer.Name,
		Email: opt.Committer.Email,
		When:  now,
	}

	var lastCommit *git.Commit
	message := ""

	for i := uint(0); i < rebase.OperationCount(); i++ {
		op := rebase.OperationAt(i)
		commit, err := repo.LookupCommit(op.Id)
		if err != nil {
			panic(err)
		}

		if lastCommit != nil &&
			(commit.ParentCount() > 1 ||
				len(messagePrefix) > 0 && !strings.HasPrefix(commit.Message(), messagePrefix)) {
			if err = applyCommit(repo, lastCommit, message, author, committer); err != nil {
				return err
			}

			message = commit.Message()
		}

		lastCommit = commit
	}

	if lastCommit != nil {
		if err = applyCommit(repo, lastCommit, message, author, committer); err != nil {
			return
		}
	}

	return rebase.Finish()
}

func applyCommit(repo *git.Repository, commit *git.Commit, message string, author, committer *git.Signature) (err error) {
	head, err := repo.Head()
	if err != nil {
		return
	}

	headCommit, err := repo.LookupCommit(head.Target())
	if err != nil {
		return
	}

	targetTree, err := commit.Tree()
	if err != nil {
		return
	}

	err = repo.CheckoutTree(targetTree, &git.CheckoutOpts{
		Strategy:         git.CheckoutSafe | git.CheckoutRecreateMissing ,
		ProgressCallback: func(path string, completed, total uint) git.ErrorCode{
			return git.ErrOk
		},
	})

	if err != nil {
		return
	}

	if commit.ParentCount() > 1 {
		panic(commit.Id().String())
	}

	if commit.Parent(0).Id() == head.Target() {
		if err = repo.SetHeadDetached(commit.Id()); err != nil {
			return
		}
		return
	}

	if len(message) == 0 {
		message = commit.Message()
	}

	newCommitID, err := repo.CreateCommit("", author, committer, message, targetTree, headCommit)
	if err != nil {
		return
	}

	return repo.SetHeadDetached(newCommitID)
}
