package wrapper

import (
    "github.com/git-roll/git-cli/pkg/utils"
    git "github.com/libgit2/git2go/v31"
)

func CreateSymbolicOrDie(name, target string) {
    repo, err := git.OpenRepository(utils.GetPwdOrDie())
    utils.DieIf(err)

    _, err = repo.References.CreateSymbolic(name, target, true, "")
    utils.DieIf(err)
}

func CreateRefOrDie(name, target string) {
    repo, err := git.OpenRepository(utils.GetPwdOrDie())
    utils.DieIf(err)

    ref, err := repo.References.Lookup(target)
    utils.DieIf(err)

    _, err = repo.References.Create(name, ref.Target(), true, "")
    utils.DieIf(err)
}
