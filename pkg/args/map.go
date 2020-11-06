package args

import (
    "fmt"
    "github.com/fatih/structs"
)

type Map map[Key]*string

type LibWrapper struct {
    m Map
    lib LibKey
}

func (w LibWrapper) MustGet(para ParameterKey) (v string, err error) {
    key := ComposeKey(w.lib, para)
    v = *w.m[key]
    if len(v) == 0 {
        err = RequisiteFailed{key}
        return
    }

    return
}

func (w LibWrapper) MayGet(para ParameterKey) string {
    return *w.m[ComposeKey(w.lib, para)]
}

func (arg Map) Git2GoWrapper() *LibWrapper {
    return &LibWrapper{
        m: arg,
        lib: LibKeyGit2Go,
    }
}

func (arg Map) GoGitWrapper() *LibWrapper {
    return &LibWrapper{
        m: arg,
        lib: LibKeyGoGit,
    }
}

type Object map[string]interface{}

func (o Object) String() string {
    return fmt.Sprintf("%#v", map[string]interface{}(o))
}

func Normalize(any interface{}) Object {
    return structs.Map(any)
}
