[![GitHub Actions][ci-image]][ci-url]
[![Go Report Card][grc-image]][grc-url]

# gofunchecks

A collections of static analysis tool to analyze function in Go.

- [goparamcount](./cmd/goparamcount)
- [goreturncount](./cmd/goreturncount)

## Quick start

Install the tool(s) you wish to use.

```shell
# goparamcount: find functions that have too many parameters.
$ go get github.com/ericcornelissen/gofunchecks/cmd/goparamcount

# goreturncount: find functions that have too many return values.
$ go get github.com/ericcornelissen/gofunchecks/cmd/goreturncount
```

And run it on your projects by using the tool name as a command.

[ci-url]: https://github.com/ericcornelissen/gofunchecks/actions
[ci-image]: https://github.com/ericcornelissen/gofunchecks/workflows/Test%20and%20Lint/badge.svg
[grc-url]: https://goreportcard.com/report/github.com/ericcornelissen/gofunchecks
[grc-image]: https://goreportcard.com/badge/github.com/ericcornelissen/gofunchecks
