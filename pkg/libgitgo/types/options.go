package types

import "io"

type Options struct {
	Progress io.Writer
	WorkDir string
	PreferredLib
	Auth
}

type Auth struct {
	User       string
	Password   string
	SSHId      string
	Passphrase string
}
