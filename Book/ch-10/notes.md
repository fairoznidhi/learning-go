# Notes

## Repositories, Modules, Packages

- **Repository**: where source lives in version control.
- **Module**: source distributed and versioned as one unit. One module per repo is recommended.
- **Package**: a directory of source code inside a module.
- **Module path**: globally unique name, usually the repo URL (`github.com/user/repo`). Case-sensitive, avoid uppercase.

## go.mod

Created with `go mod init MODULE_PATH`. Makes the current directory the module root.

- `module`: the module path.
- `go`: minimum Go version. Code must be compatible with it.
- `toolchain` / `GOTOOLCHAIN`: control which Go version builds the code (`auto`, `local`, or a version like `go1.20.4`). The env var wins over the directive.
- `require`: dependencies and minimum versions. Second block with `// indirect` lists dependencies of dependencies (documentation only).
- `replace`, `exclude`, `retract`: see below.

Since Go 1.22, `for` loops create a new variable per iteration, if the `go` directive is 1.22+ (applied per module).

## Importing and Exporting

- Uppercase first letter = exported. Lowercase or `_` = package-private.
- Exported identifiers are your API. Document them and keep them backward compatible.
- Import path = module path + path to the package.
- Unused imports are a compile error.
- The package name comes from the `package` clause, not the directory. All files in a directory must match.
- Package name should match the directory name, except for `main`, invalid identifiers (e.g. `do-format`), and version directories (`v2`).
- Relative import paths don't work with modules.

## Naming Packages

- Be descriptive. Avoid `util`.
- Package is a noun, function is a verb: `names.Extract`, `names.Format`.
- Don't repeat the package name in identifiers (`names.ExtractNames`).

## Overriding a Package Name

- Alias on name collision: `crand "crypto/rand"` next to `"math/rand"`.
- `.` import puts identifiers in the current namespace. Discouraged.
- `_` import runs `init` only (blank import).
- Package names can be shadowed by local declarations.

## Go Doc Comments

- Comment directly above the item, no blank line.
- Use `//`, and start with the symbol name (`// Add adds ...`).
- Blank `//` line splits paragraphs.
- Extra indent = preformatted block. `# ` = header.
- `[pkg.Symbol]` links to a symbol. A raw URL becomes a link.
- Text links: `[Text]` in the body, then `// [Text]: URL` at the end.
- Comment before `package` = package doc. Long ones go in `doc.go`.
- Tools: `go doc pkg`, `go doc pkg.Identifier`, `pkgsite` (HTML preview).

## internal Package

Exports in `internal` are visible only to the parent of `internal` and its subtree (siblings and their children). Anything else gets a compile error. Use it to share code in a module without making it public API.

## Circular Dependencies

If A imports B, B cannot import A (directly or indirectly). Fix by merging the packages, or moving the shared parts to one of them or a new package.

## Organizing a Module

- Small module: one package.
- **Application**: `main` at root, logic in `internal`.
- **Library**: root package name matches the repo name (no hyphens).
- Library with binaries: `cmd/<name>/main.go` for each binary.
- Organize by feature slice (customer, inventory), not by layer (models, db).
- Use `internal` to avoid Hyrum's law: anything exported will be depended on.
- `golang-standards/project-layout` is not endorsed by the Go team.

## Gracefully Renaming and Moving API

- Function/method: add a new one that calls the old one.
- Constant: add a new constant with the same value.
- Type: alias with `type Bar = Foo`. Same fields and methods, assignable without conversion. New methods/fields go on the original type.
- Aliases can't access unexported members from another package.
- Package-level variables and struct fields can't have alternate names.

## init Function

- Takes no params, returns nothing, runs the first time the package is referenced.
- Multiple `init` per package or file are allowed. Avoid them.
- Blank imports for registration (`_ "github.com/lib/pq"`) are an obsolete pattern. Register explicitly in your own code.
- Use only to set up effectively immutable package-level variables. Document any I/O.

## Third-Party Modules

- `go get ./...`: scans imports and adds missing modules to `go.mod`.
- `go get pkg`: adds a specific module (marked `// indirect` even if used directly).
- `go mod tidy`: syncs `go.mod` and `go.sum` with the source, fixes `indirect` comments, removes unused.
- No tag on a module gives a **pseudoversion** (`v0.0.0-date-hash`).
- `go.sum`: hashes for each module version and its `go.mod`. Commit both `go.mod` and `go.sum` for repeatable builds.
- Module cache: `go clean -modcache` clears it.
- Go builds from source into a single binary.

## Versions

- `go list -m -versions <module>`: list available versions.
- `go get pkg@v1.0.0`: pick a version (up or down).
- Old entries may linger in `go.sum`. Harmless.

### Semantic Versioning

`vMAJOR.MINOR.PATCH`

- Patch: bug fix.
- Minor: backward-compatible feature (patch resets to 0).
- Major: breaking change (minor and patch reset to 0).

### Minimal Version Selection

Go picks the **lowest** version that satisfies every `go.mod` in the dependency tree, not the latest. Inspect with `go mod graph`. Import compatibility rule: minor and patch versions must be backward compatible, otherwise it's a bug to report. Go imports one version per module path.

### Upgrading

| Command | Effect |
| --- | --- |
| `go get -u=patch pkg` | latest patch of the current minor |
| `go get -u pkg` | latest minor/patch of the current major |
| `go get pkg@vX.Y.Z` | exact version |

Pre-release versions (`v1.4.0-beta1`) are never selected automatically. Specify them explicitly.

### Incompatible Versions (Semantic Import Versioning)

- Major version increases, and for v2+ the module path ends in `/vN`.
- Import: `".../simpletax/v2"`.
- v1 and v2 can coexist in one program since they have different import paths.
- `go mod tidy` removes the unused old version.

## Vendoring

`go mod vendor` copies dependencies into `vendor/`, which is used instead of the module cache. Rerun after dependency changes. Falling out of favor. Useful for CI with no persistent cache. Makes the repo larger.

## pkg.go.dev

Indexes open-source Go modules: godocs, license, README, dependencies, and dependents.

## Publishing a Module

- Push to a version control system and tag it. No central upload needed.
- Commit `go.mod` and `go.sum`.
- Add a `LICENSE` file. Prefer permissive licenses (MIT, BSD, Apache). Never write your own.
- Tag versions with SemVer. Pre-release: `v1.4.0-rc2`.
- **Breaking change (v2)**: store it in a `v2/` subdirectory **or** on a `v2` branch (or `v1` for old code). Then change the `go.mod` module path to end in `/v2`, update internal imports, and tag `v2.0.0`.

## Overriding Dependencies

- `replace old => new v1.0.0`: swap in a fork. Left-side version is optional. A local path on the right has no version.
- Avoid local `replace`: it breaks for others.
- `exclude module vX.Y.Z`: stop **you** from using a version of another module.
- `retract vX.Y.Z` or `retract [v1.7.0, v1.8.5]`: tell **others** not to use versions of **your** module. Needs a new release. Existing builds keep working.

## Workspaces

Work on multiple local modules together without `replace`.

```
go work init ./workspace_app
go work use ./workspace_lib
```

- Creates `go.work` listing the modules (`use (...)`).
- Local copies are used instead of the published ones.
- Don't commit `go.work`.
- `GOWORK=off` disables it to test against the published module.
- Release in dependency order: commit, tag, then `go get` the new version in dependents.

## Module Proxy and Checksum DB

- `go get` fetches from Google's proxy (a cache of public modules), not from the repo directly.
- The checksum database records the hash of each module version. Downloads are verified against it, so tampering or a re-used tag is caught.
- `GOPROXY=direct`: skip the proxy. Or run your own (Artifactory, Sonatype, Athens).
- `GOPRIVATE=*.example.com,company.com/repo`: download these directly, keeping private repo names off public services.
