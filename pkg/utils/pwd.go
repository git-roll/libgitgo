package utils

import (
    "os"
)

func GetPwdOrDie() string {
    pwd, err := os.Getwd()
    DieIf(err)
    return pwd
}
