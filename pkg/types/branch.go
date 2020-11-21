package types

import (
    "fmt"
    "github.com/go-git/go-git/v5/config"
    git2go "github.com/libgit2/git2go/v31"
)

type Branch struct {
    Git2Go *git2go.Branch
    GoGit *config.Branch
}

func (r Branch) Name() string {
    if r.Git2Go != nil {
        name, err := r.Git2Go.Name()
        if err != nil {
            panic(err)
        }

        return name
    }

    if r.GoGit != nil {
        return r.GoGit.Name
    }

    panic("both pointers are nil")
}

func (r Branch) String() string {
    if r.Git2Go != nil {
        if r.GoGit != nil {
            panic("both pointers are filled")
        }

        name, nameErr := r.Git2Go.Name()
        up, upErr := r.Git2Go.Upstream()
        text := ""
        if nameErr != nil {
            text = fmt.Sprintf("Name():%s\n", nameErr)
        } else {
            text = fmt.Sprintf("Name():%s\n", name)
        }

        if upErr != nil {
            text += fmt.Sprintf("Upstream():%s\n", upErr)
        } else {
            text += fmt.Sprintf("Upstream():%s\n", up.Name())
        }

        return text
    }

    if r.GoGit != nil {
        if r.Git2Go != nil {
            panic("both pointers are filled")
        }

        return fmt.Sprintf("Name:%s\nRemote:%s", r.GoGit.Name, r.GoGit.Remote)
    }

    panic("both pointers are nil")
}
