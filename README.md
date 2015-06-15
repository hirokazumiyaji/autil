# autil

## Description

AWS Utility Command

## Usage

```
autil <command> <profile name>

command
  config: print profile infomation
  ec2: print ec2 host

profile name
  profile name of aws credentials(~/.aws/credentials)
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/hirokazumiyaji/autil
```

## Contribution

1. Fork ([https://github.com/hirokazumiyaji/autil/fork](https://github.com/hirokazumiyaji/autil/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Hirokazu Miyaji](https://github.com/hirokazumiyaji)
