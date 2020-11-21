package types

import (
    "fmt"
    gogit "github.com/go-git/go-git/v5"
    git2go "github.com/libgit2/git2go/v31"
)

type Remote struct {
    Git2Go *git2go.Remote
    GoGit *gogit.Remote
}

func (r Remote) Name() string {
    if r.Git2Go != nil {
        return r.Git2Go.Name()
    }

    if r.GoGit != nil {
        return r.GoGit.Config().Name
    }

    panic("both pointers are nil")
}

func (r Remote) URL() string {
    if r.Git2Go != nil {
        return r.Git2Go.Url()
    }

    if r.GoGit != nil {
        return r.GoGit.Config().URLs[0]
    }

    panic("both pointers are nil")
}

func (r Remote) String() string {
    if r.Git2Go != nil {
        if r.GoGit != nil {
            panic("both pointers are filled")
        }

        name := r.Git2Go.Name()
        url := r.Git2Go.Url()
        fetchSpecs, fetchErr := r.Git2Go.FetchRefspecs()
        pushSpecs, pushErr := r.Git2Go.PushRefspecs()
        pushUrl := r.Git2Go.PushUrl()
        text := fmt.Sprintf("Name():%s\nUrl():%s\nPushUrl():%s\n", name, url, pushUrl)
        if fetchErr != nil {
            text += fmt.Sprintf("FetchRefspecs():%s\n", fetchErr)
        } else {
            text += fmt.Sprintf("FetchRefspecs():%#v\n", fetchSpecs)
        }

        if pushErr != nil {
            text += fmt.Sprintf("PushRefspecs():%s\n", pushErr)
        } else {
            text += fmt.Sprintf("PushRefspecs():%#v", pushSpecs)
        }

        return text
    }

    if r.GoGit != nil {
        if r.Git2Go != nil {
            panic("both pointers are filled")
        }

        return r.GoGit.String()
    }

    panic("both pointers are nil")
}
