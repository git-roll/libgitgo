package arg

import (
    "fmt"
    "github.com/spf13/pflag"
)

type LibKey string

const (
    LibKeyGit2Go = LibKey("git2go")
    LibKeyGoGit = LibKey("go-git")
)

type ParameterKey string

type Key string

func ComposeKey(lib LibKey, para ParameterKey) Key {
    if len(lib) == 0 {
        return Key(para)
    }

    return Key(fmt.Sprintf("%s.%s", lib, para))
}

func RegisterFlags(flagSet *pflag.FlagSet, git2go, gogit []ParameterKey) Map {
    m := make(Map, len(git2go) + len(gogit))
    for _, para := range git2go {
        key := ComposeKey(LibKeyGit2Go, para)
        value := ""
        m[key] = &value

        flagSet.StringVar(&value, string(key), "", fmt.Sprintf("%s for library %s", para, LibKeyGit2Go))
    }

    for _, para := range gogit {
        key := ComposeKey(LibKeyGoGit, para)
        value := ""
        m[key] = &value

        flagSet.StringVar(&value, string(key), "", fmt.Sprintf("%s for library %s", para, LibKeyGit2Go))
    }

    return m
}

func RegisterCommonFlags(flagSet *pflag.FlagSet, flags []ParameterKey) Map {
    m := make(Map, len(flags))
    for _, para := range flags {
        key := ComposeKey("", para)
        value := ""
        m[key] = &value

        flagSet.StringVar(&value, string(key), "", "")
    }

    return m
}
