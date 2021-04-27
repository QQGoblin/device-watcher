#! /bin/bash

# 获取提交记录ID
GIT_COMMIT=$(git rev-parse HEAD)
GIT_SHA=$(git rev-parse --short HEAD)

# 获取tag信息
GIT_TAG=$(git describe --tags --abbrev=0 --exact-match 2>/dev/null )



VERSION_METADATA=dev-$(date "+%Y%m%d")-${GIT_SHA}
# Clear the "unreleased" string in BuildMetadata
if [[ -n $GIT_TAG ]]
then
  VERSION_METADATA=${GIT_TAG}
fi

LDFLAGS="-X github.com/QQGoblin/device-watcher/pkg/version.Version=${VERSION_METADATA}
         -X github.com/QQGoblin/device-watcher/pkg/version.GitCommit=${GIT_COMMIT}"


GO111MODULE=on
GOPROXY=https://goproxy.cn
BUILD_GOOS=${GOOS:-linux}
BUILD_GOARCH=${GOARCH:-amd64}
GOBINARY=${GOBINARY:-go}
go build -ldflags "$LDFLAGS" -v -o output/dw cmd/main.go

