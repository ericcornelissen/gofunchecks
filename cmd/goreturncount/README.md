# goreturncount

Find functions that have too many return values.

## Motivation

TODO: write motivation

## Installation

```shell
$ go get github.com/ericcornelissen/gofunchecks/cmd/goreturncount
```

## Usage

For basic usage, run the following command from the root of your project:

```shell
$ goreturncount ./...
```

Which will analyze all packages beneath the current directory. If you want to
change the number of return values that is allowed you can use the `-max` flag:

```shell
$ goreturncount -max 3 ./...
```

You can specify the maximum number of return values separately for public and
private functions using the `-public-max` and `-private-max` flags:

```shell
$ goreturncount -public-max 2 -private-max 4 ./...
```
