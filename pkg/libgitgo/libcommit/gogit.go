package libcommit

import (
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
  "github.com/go-git/go-git/v5"
  "github.com/go-git/go-git/v5/plumbing"
  "github.com/go-git/go-git/v5/plumbing/object"
  "golang.org/x/xerrors"
  "time"
)

type goGit struct {
  *types.Options
}

func (g goGit) IsAncestor(ancestor, target string) (positive bool, err error) {
  repo, err := g.OpenGoGitRepo()
  if err != nil {
    return
  }

  ancestorCommit, err := repo.CommitObject(plumbing.NewHash(ancestor))
  if err != nil {
    return
  }

  targetCommit, err := repo.CommitObject(plumbing.NewHash(target))
  if err != nil {
    return
  }

  return ancestorCommit.IsAncestor(targetCommit)
}

func (g goGit) Amend(_ string, _ *CommitOptions) (*types.Commit, error) {
  return nil, xerrors.Errorf("go-git doesn't support commit amending")
}

func (g goGit) CommitStaging(message string, opt *CommitOptions) (commit *types.Commit, err error) {
  repo, err := g.OpenGoGitRepo()
  if err != nil {
    return
  }

  wt, err := repo.Worktree()
  if err != nil {
    return
  }

  now := time.Now()
  hash, err := wt.Commit(message, &git.CommitOptions{
    All:       opt.All,
    Author:    &object.Signature{
      Name:  opt.Author.Name,
      Email: opt.Author.Email,
      When:  now,
    },
    Committer: &object.Signature{
      Name:  opt.Committer.Name,
      Email: opt.Committer.Email,
      When:  now,
    },
  })

  if err != nil {
    return
  }

  obj, err := repo.CommitObject(hash)
  if err != nil {
    return
  }

  commit = &types.Commit{GoGit: obj}
  return
}

func (g goGit) Get(refName string) (commit *types.Commit, err error) {
  repo, err := g.Options.OpenGoGitRepo()
  if err != nil {
    return
  }

  ref, err := repo.Reference(plumbing.ReferenceName(refName), true)
  if err != nil {
    return
  }

  obj, err := repo.CommitObject(ref.Hash())
  if err != nil {
    return
  }

  commit = &types.Commit{GoGit: obj}
  return
}
