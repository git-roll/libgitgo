package refspec

import (
    "fmt"
)

func PushBranch(branch string) string {
    return fmt.Sprintf("refs/heads/%s:refs/heads/%s", branch, branch)
}

func FetchBranch(branch, remote string) string {
    if len(branch) == 0 {
        return fmt.Sprintf("+refs/heads/*:refs/remotes/%s/*", remote)
    }

    return fmt.Sprintf("+refs/heads/%s:refs/remotes/%s/%s", branch, remote, branch)
}
