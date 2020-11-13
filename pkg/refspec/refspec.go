package refspec

import (
    "fmt"
)

func PushBranch(branch string) string {
    return fmt.Sprintf("refs/heads/%s:refs/heads/%s", branch, branch)
}
