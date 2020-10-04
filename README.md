# goparamcount

Find functions that have too many parameters.

## Motivation

Functions with many parameters can be difficult to understand and the order of
parameters may be confusing. Also, having many parameters can be an indication
that the function is doing too many things. And, if you really need a lot of
data it may be a better idea to wrap it in a `struct` anyway.

## Installation

```shell
$ go get github.com/ericcornelissen/goparamcount
```

## Usage

For basic usage, run the following command from the root of your project:

```shell
$ goparamcount ./...
```

Which will analyze all packages beneath the current directory. If you want to
change the number of parameters that is allowed you can use the `-max` flag:

```shell
$ goparamcount -max 3 ./...
```

You can specify the maximum number of parameters separately for public and
private functions using the `-public-max` and `-private-max` flags:

```shell
$ goparamcount -public-max 2 -private-max 4 ./...
```
