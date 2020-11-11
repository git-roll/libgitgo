package libclone

import (
	"context"
	"fmt"
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	sshCli "golang.org/x/crypto/ssh"
	"os"
	"path/filepath"
)

type goGit struct {
	workdir  string
	ctx      context.Context
	progress sideband.Progress
	auth     *types.Auth
}

func (g goGit) Start(url string, branch string, bare bool, cloneOpt *Option) (*types.Repository, error) {
	ep, err := transport.NewEndpoint(url)
	if err != nil {
		return nil, err
	}

	var auth transport.AuthMethod
	switch ep.Protocol {
	case "ssh":
		if len(ep.User) > 0 && len(ep.Password) > 0 {
			auth = &ssh.Password{
				User:                  ep.User,
				Password:              ep.Password,
				HostKeyCallbackHelper: ssh.HostKeyCallbackHelper{HostKeyCallback: sshCli.InsecureIgnoreHostKey()},
			}
		} else if len(g.auth.User) > 0 {
			auth = &ssh.Password{
				User:                  g.auth.User,
				Password:              g.auth.Password,
				HostKeyCallbackHelper: ssh.HostKeyCallbackHelper{HostKeyCallback: sshCli.InsecureIgnoreHostKey()},
			}
		} else if len(g.auth.SSHId) > 0 {
			if g.auth.SSHId[0] == '~' {
				home, err := os.UserHomeDir()
				if err != nil {
					return nil, err
				}

				g.auth.SSHId = filepath.Join(home, g.auth.SSHId[1:])
			}

			auth, err = ssh.NewPublicKeysFromFile(ssh.DefaultUsername, g.auth.SSHId, g.auth.Passphrase)
			if err != nil {
				return nil, err
			}
		}
	case "file", "http", "https":
		if len(ep.User) > 0 && len(ep.Password) > 0 {
			auth = &http.BasicAuth{
				Username: ep.User,
				Password: ep.Password,
			}
		}
	default:
		return nil, fmt.Errorf("unsupported protocol %s", ep.Protocol)
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
		RecurseSubmodules: opt.SubmoduleRescursivity,
		Progress:          g.progress,
		Tags:              opt.TagMode,
	})

	if err != nil {
		return nil, err
	}

	return &types.Repository{GoGit: r}, nil
}
