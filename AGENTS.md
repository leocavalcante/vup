# AGENTS.md

This file provides instructions for AI agents working with this repository.

## Project Overview

`vup` is a command-line tool for managing semantic version strings. It can increment or decrement major, minor, and patch versions. It also supports the 'v' prefix in version strings.

## Build and Test Commands

### Build

To build the project, run:

```shell
make vup
```

### Test

To run the tests, run:

```shell
make test
```

### Lint

To lint the code, run:

```shell
make lint
```

## Code Style

This project follows standard Go programming language conventions. Please ensure your code is formatted with `gofmt` before submitting.

## Architecture

The project is structured as follows:

- `main.go`: The entry point of the application.
- `cmd/`: Contains the command-line interface logic, using the Cobra library. Each command is in its own file.
- `internal/`: Contains the core logic of the application, separated into a `vup` package.

## Contribution Guidelines

- New features should be accompanied by tests.
- Please run `make lint` and `make test` before submitting your changes.
- Commit messages should be clear and descriptive.
