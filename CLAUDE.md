# Pocket Network Development Guide

This file was generated by `/init` of [claude-code](https://github.com/anthropics/claude-code) as of #1092.

## Common Commands

- **Build**: `make ignite_pocketd_build`
- **Lint**: `make go_lint`
- **Format**: `make go_imports`
- **Test All**: `make test_all`
- **Single Test**: `go test -v ./path/to/package -run TestName`
- **E2E Test**: `make test_e2e`
- **Regen Protos**: `make proto_regen`
- **Dev Setup**: `make go_develop && make go_develop_and_test`

## Code Style

- **Imports**: Stdlib first, third-party second, local last (alphabetically within groups)
- **Naming**: CamelCase (exported), camelCase (private), error vars prefixed with "Err"
- **Error Handling**: Return errors with context using Wrapf, early returns
- **Indentation**: Tabs for Go code, consistent with gofmt
- **Comments**: Godoc for exported items, implementation notes with DEV_NOTE
- **Types**: Strong use of generics, interface-based design, explicit type conversions
