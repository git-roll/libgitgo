package libindex

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	*types.Options
}

func (g git2go) Add(paths []string) (err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	index, err := repo.Index()
	if err != nil {
		return
	}

	if err = index.AddAll(paths, git.IndexAddDefault, nil); err != nil {
		return
	}

	return index.Write()
}
