package remote

import "github.com/git-roll/git-cli/pkg/args"

type Runner interface {
    List() ([]string, error)
    Create(args.Map) error
    Lookup(name string) (remote args.Object, err error)
}

var (
    Git2GoParams []args.ParameterKey
    GoGitParams []args.ParameterKey
)

func Run(lib args.LibKey, workdir string) Runner {
    switch lib {
    case args.LibKeyGoGit:
        return &goGit{worktree: workdir}
    case args.LibKeyGit2Go:
        return &git2go{worktree: workdir}
    }

    return nil
}

const (
    ParameterKeyName      = args.ParameterKey("name")
    ParameterKeyURL       = args.ParameterKey("url")
    ParameterKeyFetchSpec = args.ParameterKey("fetchSpec")
)

func init() {
    Git2GoParams = append(Git2GoParams,
        ParameterKeyName,
        ParameterKeyURL,
        ParameterKeyFetchSpec,
    )

    GoGitParams = Git2GoParams
}
