package libignore

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

type git2go struct {
	*types.Options
}

func (g git2go) Check(relativePath string) (err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	ignored, err := repo.IsPathIgnored(relativePath)
	if err != nil {
		return
	}

	if ignored {
		err = ErrPathIgnored
	}

	return
}
