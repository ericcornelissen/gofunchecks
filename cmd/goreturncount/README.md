# goreturncount

Find functions that have too many return values.

## Motivation

Functions with many return values can be difficult to understand and the order
of the values may be confusing. Also, having many return values can be an
indication that the function is doing too many things. If you really need to
return a lot of data it may be a better idea to wrap it in a `struct` anyway.

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
