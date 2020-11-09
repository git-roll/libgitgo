package clone

import (
    "context"
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
    gogit "github.com/go-git/go-git/v5"
    gitgo "github.com/libgit2/git2go/v31"
)

type Git2GoOption struct {
    gitgo.DownloadTags
}

type GoGitOption struct {
    // usually `origin`
    RemoteName string
    // specified branch only
    SingleBranch bool
    NoCheckout bool
    Depth int
    gogit.SubmoduleRescursivity
    gogit.TagMode
}

func Start(url string, branch string, bare bool, opt1 *Git2GoOption, opt2 *GoGitOption, opt *types.Options) (*types.Repository, error) {
    return with(opt).Start(url, branch, bare, opt1, opt2)
}

type wrapper interface {
    Start(url string, branch string, bare bool, opt1 *Git2GoOption, opt2 *GoGitOption) (*types.Repository, error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        return &goGit{workdir: opt.WorkDir, auth: &opt.Auth, ctx: context.TODO(), progress: opt.Progress}
    case types.PreferGit2Go:
        return &git2go{workdir: opt.WorkDir}
    }

    return nil
}
