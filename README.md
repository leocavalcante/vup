# vup

A version manager for semantic version based strings.

## Getting started

### Install

#### Homebrew

```shell
brew install leocavalcante/tap/vup
```

### Quick start

```shell
vup minor v0.1.0 
# v0.2.0
```

```shell
vup major 1.2.3
# 2.2.3
```

## Usage

Given a semantic version string...
```
major.minor.patch
```

| Command | Description |
| --- | --- |
| `vup major` | Updates the **major** part of the string |
| `vup minor` | Updates the **minor** part of the string |
| `vup patch` | Updates the **patch** part of the string |
| `vup rc` | Updates or promotes the **release candidate** part of the string |

### Prefix

You can use the `v` letter as version prefix, like **v**1.0.1, it will be handled properly.

### Downgrades

By default `vup` will upgrade (ie. increase) the version number, but you can add the `-d` flag to make downgrades:

```shell
vup major -d v1.2.3
# v0.2.3
```

### Release candidates

Add `--rc` to a core version command to advance that component and start a
release candidate series at `rc1`. Lower core components are reset when
advancing a major or minor version.

```shell
vup minor --rc 1.2.3
# 1.3.0-rc1
```

Use the `rc` command to increment an existing release candidate:

```shell
vup rc 1.3.0-rc1
# 1.3.0-rc2
```

The `rc` command requires an existing `-rcN` suffix. Promote a candidate to its
final release with `--promote`:

```shell
vup rc --promote 1.3.0-rc2
# 1.3.0
```

### Examples

#### Combine with `git` tags

```shell
vup minor -u $(git describe --tags --abbrev=0)
```

#### GoReleaser + Makefile (`make minor`)

```makefile
.PHONY: minor
minor:
	@git tag $$(vup minor $$(git describe --tags --abbrev=0))
	@go tool goreleaser --clean
```
