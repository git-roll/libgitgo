package libpush

import (
    "bytes"
    "github.com/git-roll/libgitgo/pkg/refspec"
    "github.com/git-roll/libgitgo/pkg/types"
    git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	*types.Options
}

func (g git2go) Start(branches []string, remoteName string, force bool) (remoteOut string, err error) {
	repo, err := g.OpenGit2GoRepo()
	if err != nil {
		return
	}

	remote, err := repo.Remotes.Lookup(remoteName)
	if err != nil {
		return
	}

	var refSpecs []string
	if len(branches) == 0 {
		refSpecs = []string{refspec.PushBranch("*")}
	} else {
		for _, br := range branches {
			if len(br) > 0 {
				refSpecs = append(refSpecs, refspec.PushBranch(br))
				continue
			}

			if refSpecs, err = remote.PushRefspecs(); err != nil {
				return
			}

			if len(refSpecs) > 0 {
				break
			}

			head, err := repo.Head()
			if err != nil {
				return "", err
			}

			br, err = head.Branch().Name()
			if err != nil {
				return "", err
			}

			refSpecs = []string{refspec.PushBranch(br)}
		}
	}

	out := &bytes.Buffer{}

	err = remote.Push(refSpecs, &git.PushOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			SidebandProgressCallback:     func(str string) git.ErrorCode{
				out.WriteString(str)
				return git.ErrOk
			},
			CompletionCallback:           func(git.RemoteCompletion) git.ErrorCode{
				return git.ErrOk
			},
			CredentialsCallback:          g.GenGit2GoAuth,
			TransferProgressCallback:     func(stats git.TransferProgress) git.ErrorCode {
				return git.ErrOk
			},
			UpdateTipsCallback:           func(refname string, a *git.Oid, b *git.Oid) git.ErrorCode{
				return git.ErrOk
			},
			CertificateCheckCallback:     func(cert *git.Certificate, valid bool, hostname string) git.ErrorCode{
				return git.ErrOk
			},
			PackProgressCallback:         func(stage int32, current, total uint32) git.ErrorCode {
				return git.ErrOk
			},
			PushTransferProgressCallback: func(current, total uint32, bytes uint) git.ErrorCode{
				return git.ErrOk
			},
			PushUpdateReferenceCallback:  func(refname, status string) git.ErrorCode {
				return git.ErrOk
			},
		},
	})

	return out.String(), nil
}
