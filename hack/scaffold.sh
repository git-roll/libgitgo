#!/usr/bin/env zsh

set -e

if [ -z $1 ]; then
  echo "a single-word package name is required. such as branch, clone, pull balabala"
  exit 2
fi

echo -n "🐳 Creating directory for the lib..."
mkdir -p "pkg/libgitgo/$1"
echo "Done"

ObjectExp="err"

if [ $# -gt 1 ]; then
  echo -n "🐳 Generating types..."
  ObjectName=${(C)1}
  ObjectExp="*types.${ObjectName},"
  cat > "pkg/libgitgo/types/$1.go" <<EOF
package types

import (
    "fmt"
    gogit "github.com/go-git/go-git/v5"
    git2go "github.com/libgit2/git2go/v31"
)

type ${ObjectName} struct {
    Git2Go *git2go.${ObjectName}
    GoGit *gogit.${ObjectName}
}

func (r ${ObjectName}) String() string {
  if r.Git2Go != nil {
    if r.GoGit != nil {
      panic("both pointers are filled")
    }
  }

  if r.GoGit != nil {
    if r.Git2Go != nil {
      panic("both pointers are filled")
    }
  }

  panic("both pointers are nil")
}
EOF
  echo "Done"
fi

echo -n "🐳 Generating interfaces..."
cat > "pkg/libgitgo/$1/api.go" <<EOF
package $1

import (
    "github.com/git-roll/libgitgo/pkg/libgitgo/types"
)

func Open(opt *types.Options) (${ObjectExp} error) {
    return with(opt).Open()
}

type wrapper interface {
    Open() (${ObjectExp} error)
}

func with(opt *types.Options) wrapper {
    switch opt.PreferredLib {
    case types.PreferGoGit:
        return &goGit{workdir: opt.WorkDir}
    case types.PreferGit2Go:
        return &git2go{workdir: opt.WorkDir}
    }

    return nil
}
EOF
echo "Done"

echo -n "🐳 Generating implementation for git2go..."
cat > "pkg/libgitgo/$1/git2go.go" <<EOF
package $1

import (
	"github.com/git-roll/libgitgo/pkg/libgitgo/types"
	git "github.com/libgit2/git2go/v31"
)

type git2go struct {
	workdir string
}

func (g git2go) Open() (${ObjectExp} error) {
  panic("implement me")
}
EOF
echo "Done"

echo -n "🐳 Generating implementation for go-git..."
cat > "pkg/libgitgo/$1/gogit.go" <<EOF
package $1

import (
  "github.com/git-roll/libgitgo/pkg/libgitgo/types"
  "github.com/go-git/go-git/v5"
)

type goGit struct {
  workdir string
}

func (g goGit) Open() (${ObjectExp} error) {
  panic("implement me")
}
EOF
echo "Done"

set +e
