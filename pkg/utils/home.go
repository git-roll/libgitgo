package utils

import (
    "os"
    "path/filepath"
)

func ResolveHomePath(str string) (string, error) {
    if str[0] != '~' {
        return str, nil
    }

    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }

    return filepath.Join(home, str[1:]), nil
}
