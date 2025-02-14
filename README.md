# vup

A version manager for semantic version based strings.

## Getting started

### Install

TODO

### Quick start

```shell
vup minor -u v0.1.0 
#v0.2.0
```

```shell
vup major -d v2.0.0
#v1.0.0
```

## Usage

### Combine with `git` tags

```shell
vup minor -u $(git describe --tags --abbrev=0)
```
