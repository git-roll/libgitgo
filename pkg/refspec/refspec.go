package refspec

import (
    "fmt"
    "github.com/go-git/go-git/v5/config"
)

func PushBranch(remote, branch string) config.RefSpec {
    return config.RefSpec(fmt.Sprintf("refs/heads/%s:refs/remote/%s/%s", branch, remote, branch))
}
