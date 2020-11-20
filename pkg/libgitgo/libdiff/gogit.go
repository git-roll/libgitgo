package libdiff

import (
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

type goGit struct {
  *types.Options
}

func (g goGit) HeadToWorkDir() (*types.Diff, error) {
  panic("go-git doesn't support diff calculation against worktree")
}
