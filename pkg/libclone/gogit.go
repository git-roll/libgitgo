package libclone

import (
	"context"
	"github.com/git-roll/libgitgo/pkg/types"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
)

type goGit struct {
	workdir  string
	ctx      context.Context
	progress sideband.Progress
	auth     *types.Auth
}

func (g goGit) Start(url string, branch string, bare bool, cloneOpt *Options) (*types.Repository, error) {
	auth, err := g.auth.GenGoGitAuth(url)
	if err != nil {
		return nil, err
	}

	opt := cloneOpt.GoGit
	var refs plumbing.ReferenceName
	if len(opt.RemoteName) > 0 && len(branch) > 0 {
		refs = plumbing.NewRemoteReferenceName(opt.RemoteName, branch)
	}

	r, err := git.PlainCloneContext(g.ctx, g.workdir, bare, &git.CloneOptions{
		URL:               url,
		Auth:              auth,
		RemoteName:        opt.RemoteName,
		ReferenceName:     refs,
		SingleBranch:      opt.SingleBranch,
		NoCheckout:        opt.NoCheckout,
		Depth:             opt.Depth,
		RecurseSubmodules: git.SubmoduleRescursivity(opt.SubmoduleRescursivity),
		Progress:          g.progress,
		Tags:              git.TagMode(opt.TagMode),
	})

	if err != nil {
		return nil, err
	}

	return &types.Repository{GoGit: r}, nil
}
