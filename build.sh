#!/bin/sh

NAME="briantest"
VERSION="0.1.2"
EMAIL="chenliang@idcos.com"

rm -rf go.mod go.sum vendor terraform-provider-*
go mod init terraform-provider-briantest && go mod tidy && go mod vendor

OS="darwin"
ARCH="arm64"
GOOS=${OS} GOARCH=${ARCH} go build -o terraform-provider-${NAME}_${VERSION} && zip terraform-provider-${NAME}_${VERSION}_${OS}_${ARCH}.zip terraform-provider-${NAME}_${VERSION}

OS="darwin"
ARCH="amd64"
GOOS=${OS} GOARCH=${ARCH} go build -o terraform-provider-${NAME}_${VERSION} && zip terraform-provider-${NAME}_${VERSION}_${OS}_${ARCH}.zip terraform-provider-${NAME}_${VERSION}

OS="windows"
ARCH="amd64"
GOOS=${OS} GOARCH=${ARCH} go build -o terraform-provider-${NAME}_${VERSION} && zip terraform-provider-${NAME}_${VERSION}_${OS}_${ARCH}.zip terraform-provider-${NAME}_${VERSION}

OS="linux"
ARCH="amd64"
GOOS=${OS} GOARCH=${ARCH} go build -o terraform-provider-${NAME}_${VERSION} && zip terraform-provider-${NAME}_${VERSION}_${OS}_${ARCH}.zip terraform-provider-${NAME}_${VERSION}

shasum -a 256 *.zip > terraform-provider-briantest_0.1.2_SHA256SUMS && gpg -u ${EMAIL} --detach-sign terraform-provider-briantest_0.1.2_SHA256SUMS
ls -l terraform-provider-${NAME}_${VERSION}_*