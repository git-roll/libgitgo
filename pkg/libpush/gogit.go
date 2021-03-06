package libpush

import (
	"bytes"
	"fmt"
	"github.com/git-roll/libgitgo/pkg/refspec"
	"github.com/git-roll/libgitgo/pkg/types"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type goGit struct {
	*types.Options
}

func (g goGit) Start(branches []string, remoteName string, force bool) (remoteOut string, err error) {
	repo, err := g.Options.OpenGoGitRepo()
	if err != nil {
		return
	}

	remote, err := repo.Remote(remoteName)
	if err != nil {
		return
	}

	var url string
	if len(remote.Config().URLs) > 0 {
		url = remote.Config().URLs[0]
	}

	auth, err := g.Options.Auth.GenGoGitAuth(url)
	if err != nil {
		return
	}

	var refSpecs []config.RefSpec
	if len(branches) == 0 {
      refSpecs = []config.RefSpec{config.RefSpec(refspec.PushBranch("*"))}
	} else {
		for _, br := range branches {
			if len(br) > 0 {
				refSpecs = append(refSpecs, config.RefSpec(refspec.PushBranch(br)))
				continue
			}

			head, err := repo.Head()
			if err != nil {
				return "", err
			}

			if !head.Name().IsBranch() {
				return "", fmt.Errorf("the current head is not a branch: %s", head.Name().Short())
			}

			refSpecs = append(refSpecs, config.RefSpec(refspec.PushBranch(head.Name().Short())))
		}
	}

	out := &bytes.Buffer{}
	opt := &git.PushOptions{
		RemoteName: remoteName,
		RefSpecs:   refSpecs,
		Auth:       auth,
		Progress:   out,
		Force:      force,
	}

	if g.Options.Context != nil {
		err = repo.PushContext(g.Options.Context, opt)
	} else {
		err = repo.Push(opt)
	}

	return out.String(), err
}
