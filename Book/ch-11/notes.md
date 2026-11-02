# Notes

## Go Directory Layout

```
~/go/                ← GOPATH: Go's root directory (default)
│
├── bin/               ← GOBIN: installed executables go here (if configured)
│                        PATH: where the terminal looks for executables
│
└── pkg/
    └── mod/           ← downloaded dependencies (cache)
```

```
go env GOPATH
go env GOBIN
echo $PATH
```

## go run

`go run main.go` builds the binary in a temporary directory, executes it from there, and deletes it when the program exits.

## go get vs go install

- `go get` adds or updates dependencies in a Go project.
- `go install` builds and installs an executable tool on the computer. The binary goes to `GOBIN`.
- Why it's used: it has become the method of choice for distributing third-party developer tools. One command (`go install pkg@version`) downloads, builds, and installs, with no manual build or prebuilt binaries needed.

## go vet

Built in. Catches likely bugs in code that compiles (wrong `Printf` verbs, unreachable code, copied locks). Run `go vet ./...`.

## Linters

Third-party tools that go further than `go vet` (unused code, deprecated APIs, style).

- `staticcheck`: bugs, performance, simplifications. Low noise.
- `revive`: style and naming. Replaces the deprecated `golint`.
- `golangci-lint`: runs many linters (including the two above) together. Common in CI.

## govulncheck

Official tool that scans code and dependencies for known vulnerabilities, reporting only those actually called.