package arg

type Map map[Key]*string

func OneLineMap(k ParameterKey, v *string) Map  {
    return Map{Key(k): v}
}

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

func (w LibWrapper) Get(para ParameterKey) string {
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

func (arg Map) Get(para ParameterKey) string {
    return *arg[Key(para)]
}
