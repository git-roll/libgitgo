package libmerge

import (
	"fmt"
	"github.com/git-roll/libgitgo/pkg/types"
	git "github.com/libgit2/git2go/v31"
	"time"
)

type git2go struct {
	*types.Options
}

func (g git2go) Start(branch, remote string, opt *MergeOptions) (err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	// validate worktree status
	st, err := repo.StatusList(&git.StatusOptions{
		Show:     git.StatusShowIndexAndWorkdir,
	})

	if err != nil {
		return
	}
	if count, err := st.EntryCount(); err != nil || count > 0 {
		return ErrWorktreeIsNotClean
	}

	var refSpec string
	if len(remote) > 0 {
		refSpec = fmt.Sprintf("refs/remotes/%s/%s", remote, branch)
	} else {
		refSpec = fmt.Sprintf("refs/heads/%s", branch)
	}

	commit, err := repo.AnnotatedCommitFromRevspec(refSpec)
	if err != nil {
		return
	}

	analysis, preference, err := repo.MergeAnalysis([]*git.AnnotatedCommit{commit})
	if err != nil {
		return
	}

	if (analysis & git.MergeAnalysisUpToDate) > 0 {
		return nil
	}

	if (analysis & git.MergeAnalysisUnborn) > 0 ||
		(analysis & git.MergeAnalysisFastForward) > 0 && (preference & git.MergePreferenceNoFastForward) == 0 {
		head, err := repo.Head()
		if err != nil {
			return err
		}

		_, err = head.SetTarget(commit.Id(), "")
		return nil
	}

	err = repo.Merge([]*git.AnnotatedCommit{commit}, &git.MergeOptions{
		TreeFlags:       git.MergeTreeFailOnConflict,
	}, &git.CheckoutOpts{})

	if err != nil {
		return
	}

	// check index conflicts
	index, err := repo.Index()
	if err != nil {
		return
	}

	if index.HasConflicts() {
		return ErrConflictAfterMerging
	}

	// create merge commit
	treeOid, err := index.WriteTree()
	if err != nil {
		return
	}

	if err = index.Write(); err != nil {
		return
	}

	head, err := repo.Head()
	if err != nil {
		return
	}

	now := time.Now()
	oid, err := repo.CreateCommitFromIds("", &git.Signature{
		Name:  opt.Author.Name,
		Email: opt.Author.Email,
		When:  now,
	}, &git.Signature{
		Name:  opt.Committer.Name,
		Email: opt.Committer.Email,
		When:  now,
	}, fmt.Sprintf("merge commit %s", commit.Id().String()), treeOid, head.Target(), commit.Id())

	_, err = head.SetTarget(oid, "")
	return
}

func (g git2go) FastForward(branch, remote string) (err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	var refSpec string
	if len(remote) > 0 {
		refSpec = fmt.Sprintf("refs/remotes/%s/%s", remote, branch)
	} else {
		refSpec = fmt.Sprintf("refs/heads/%s", branch)
	}

	commit, err := repo.AnnotatedCommitFromRevspec(refSpec)
	if err != nil {
		return
	}

	analysis, _, err := repo.MergeAnalysis([]*git.AnnotatedCommit{commit})
	if err != nil {
		return
	}

	if (analysis & git.MergeAnalysisUpToDate) > 0 {
		return nil
	}

	if (analysis & git.MergeAnalysisFastForward) == 0 {
		return fmt.Errorf("can't merge if not fast forward")
	}

	head, err := repo.Head()
	if err != nil {
		return
	}

	_, err = head.SetTarget(commit.Id(), "")
	return
}
