package libref

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	*types.Options
}

func (g git2go) List() (refs []*types.Reference, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	i, err := repo.NewReferenceIterator()
	if err != nil {
		return
	}

	for ref, err := i.Next(); err == nil; ref, err = i.Next()  {
		refs = append(refs, &types.Reference{Git2Go: ref})
	}

	if git.IsErrorCode(err, git.ErrIterOver) {
		err = nil
	}

	return
}
