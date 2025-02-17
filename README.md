# vup

A version manager for semantic version based strings.

## Getting started

### Install

TODO

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
| `vup major` | Updates the major part of the string |
| `vup minor` | Updates the minor part of the string |
| `vup patch` | Updates the patch part of the string |

### Prefix

You can use the `v` letter as version prefix, like _v_1.0.1, it will be handled properly.

### Downgrades

By default `vup` will upgrade (ie. increase) the version number, but you can add the `-d` flag to make downgrades:

```shell
vup major -d v1.2.3
# v0.2.3
```

### Examples

#### Combine with `git` tags

```shell
vup minor -u $(git describe --tags --abbrev=0)
```
