# go_protobuf_users

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

# Go development

--- .bash_profile ----

export GOPATH="${HOME}/.go"

export GOROOT="$(brew --prefix golang)/libexec"

export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"

test -d "${GOPATH}" || mkdir "${GOPATH}"

test -d "${GOPATH}/src/github.com" || mkdir -p "${GOPATH}/src/github.com"

--- brew ---

brew install go

go get golang.org/x/tools/cmd/godoc

go get github.com/golang/lint/golint

brew install protobuf

brew install protoc-gen-go

--- GIT ---

git clone https://github.com/jeisberg/go_protobuf_users

--- PROTO ---

cd go_protobuf_users

protoc -I users users/users.proto --go_out=plugins=grpc:users

--- Server ---

go build main.go

./main

--- CLIENT ---

go build client.go

./client