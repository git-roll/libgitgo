package utils

import (
    "fmt"
    "os"
)

func DieIf(err error) {
    if err == nil {
        return
    }

    fmt.Fprintln(os.Stderr, err.Error())
    os.Exit(2)
}
