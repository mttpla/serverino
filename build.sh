#!/usr/bin/env bash

PACKAGE="github.com/mttpla/serverino/pkg/utils"
COMMIT_HASH="$(git rev-parse HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')
GIT_DESCRIBE="$(git describe --always)"

LDFLAGS=(
  "-X '${PACKAGE}.FlagCommitHash=${COMMIT_HASH}'"
  "-X '${PACKAGE}.FlagBuildTimestamp=${BUILD_TIMESTAMP}'"
  "-X '${PACKAGE}.FlagGitDescribe=${GIT_DESCRIBE}'"
)

gofmt -s -w .
go build -ldflags="${LDFLAGS[*]}" -o serverino.out ./cmd/serverino.go