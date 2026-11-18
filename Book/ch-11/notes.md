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

## Embedding Files (`go:embed`)

Bundles support files into the binary, so you still ship a single executable.

### Setup

- Put `//go:embed <path>` directly above a **package-level** variable (no space after `//`).
- Treat embedded variables as immutable.

### Variable types

| Type | Use for | Import |
| --- | --- | --- |
| `string` / `[]byte` | a single file | blank import `_ "embed"` |
| `embed.FS` | files and directories | `"embed"` |

### `embed.FS`

- Acts as a virtual filesystem (works with `fs.WalkDir`, `ReadFile`).
- The directory name is part of the path: embedding `help` means reading `help/info.txt`.

### Patterns

- Multiple names are space-separated.
- Wildcards follow `path.Match` (`*`, `?`).
- Compile error if a pattern matches nothing, or if a `string`/`[]byte` gets multiple files.

### Hidden files (`.` or `_`)

Skipped by default.

- `dir/*`: includes them one level deep.
- `all:dir`: includes them in all subdirectories.

## `go generate`

Runs commands written in `//go:generate <cmd>` comments (e.g. `protoc` to produce `.pb.go` files). `go build` never runs it, so `go generate ./...` must be run manually and the output committed.

## `go generate` vs Makefile

- `go generate`: creates source code. Makefile: validates and compiles.
- Commit generated code.
- Make `generate` a dependency of `build`.
- Skip committing if output has timestamps or generation is slow.

## Build Info in a Binary

- Every `go build` binary embeds module versions, build flags, and VCS revision.
- View it with `go version -m <binary>`.
- `govulncheck -mode binary <binary>` scans it for known vulnerabilities (no exact lines; check the source for those).
- Read it in code with the `debug/buildinfo` package.

## Cross-Compiling

- Binaries are native to one OS and CPU. Set `GOOS` and `GOARCH` to build for another.
- Defaults to the current machine. `darwin` = macOS, `amd64` = 64-bit Intel.
- Example: `GOOS=linux GOARCH=amd64 go build`
- Check the result with `file <binary>`.

## Build Tags and Platform-Specific Code

Two ways to compile a file only for certain targets:

- **Filename**: `name_GOOS_GOARCH.go`, e.g. `x_windows.go`, `x_windows_arm64.go`.
- **Build tag**: `//go:build <expr>` on the line before `package` (no space after `//`). Supports `&&`, `||`, `!`, parentheses.
  - Example: `//go:build arm64 && darwin && !ios`
  - Meta tags: `unix`, `cgo`.
- Filename is easier to find, tags are more precise. Both can be combined.
- Custom tags: `//go:build gopher`, enabled with `-tags gopher`.
- `//go:build ignore` skips a file.
