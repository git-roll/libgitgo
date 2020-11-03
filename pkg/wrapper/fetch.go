package wrapper

import (
    "fmt"
    "github.com/git-roll/git-cli/pkg/utils"
    remoteGit "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/config"
    "github.com/go-git/go-git/v5/plumbing"
    "github.com/go-git/go-git/v5/plumbing/transport/ssh"
    "os"
    "path/filepath"
)

func FetchOrDie(remoteName, branch string) {
    if len(remoteName) == 0 {
        fmt.Fprintln(os.Stderr, "remote [branch]")
        return
    }

    repo, err := remoteGit.PlainOpen(utils.GetPwdOrDie())
    utils.DieIf(err)

    remote, err := repo.Remote(remoteName)
    utils.DieIf(err)

    var refs []config.RefSpec
    if len(branch) > 0 {
        for _, fetch := range remote.Config().Fetch {
            if fetch.Match(plumbing.NewBranchReferenceName(branch)) {
                refs = append(refs, fetch)
            }
        }

        if len(refs) == 0 {
            fmt.Fprintln(os.Stderr, "can't fetch branch", branch)
            return
        }
    } else {
        refs = remote.Config().Fetch
    }

    home, err := os.UserHomeDir()
    utils.DieIf(err)

    auth, err := ssh.NewPublicKeysFromFile(ssh.DefaultUsername, filepath.Join(home, "Documents/keys/client-test"), "")
    utils.DieIf(err)

    err = repo.Fetch(&remoteGit.FetchOptions{
        RemoteName: remoteName,
        RefSpecs:   refs,
        Auth:       auth,
        Progress:   os.Stdout,
        Tags:       0,
    })

    fmt.Println("Fetching", remoteName)

    if err == remoteGit.NoErrAlreadyUpToDate {
        fmt.Println(err.Error())
        return
    }

    utils.DieIf(err)
    fmt.Println(remoteName, "Fetched")
}
