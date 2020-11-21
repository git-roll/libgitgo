package libfs

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
    "os"
)

func ReadFile(name string, opt *types.Options) ([]byte, error) {
    return with(opt).ReadFile(name)
}

func WriteFile(name, content string, opt *types.Options) error {
    return with(opt).WriteFile(name, content)
}

func Stat(name string, opt *types.Options) (os.FileInfo, error) {
    return with(opt).Stat(name)
}

type wrapper interface {
    ReadFile(string) ([]byte, error)
    WriteFile(name, content string) error
    Stat(string) (os.FileInfo, error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGit2Go:
        panic("go2git doesn't support fs operation")
    case types.PreferGoGit:
        fallthrough
    default:
        return &goGit{opt}
    }
}
