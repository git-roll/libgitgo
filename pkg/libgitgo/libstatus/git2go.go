package libstatus

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	*types.Options
}

func (g git2go) List() (list *types.Status, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	statusList, err := repo.StatusList(&git.StatusOptions{
		Show:     git.StatusShowIndexAndWorkdir,
	})

	if err != nil {
		return
	}

	list = &types.Status{Git2Go: statusList}
	return
}
