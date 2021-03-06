package main

const Name string = "qucli"

// To set this from outside, use go build -ldflags "-X main.Version \"$(VERSION)\""
var Version string

// GitCommit describes latest commit hash.
// This value is extracted by git command when building.
// To set this from outside, use go build -ldflags "-X main.GitCommit \"$(COMMIT)\""
var GitCommit string
