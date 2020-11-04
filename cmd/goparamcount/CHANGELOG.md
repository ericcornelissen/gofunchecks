# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to [Semantic
Versioning].

## [Unreleased]

- Fix recursive analysis on Windows.

## [1.0.1] - 2020-11-02

- Add `-legal` flag.
- Add `-version` flag.
- Fix message in case function has one parameter.

## [1.0.0] - 2020-10-06

- Add `-tests` flag.
- Add one-letter aliases for common flags.
- Fix mistake in usage message (double "goparamcount:" prefix).

## [0.1.5] - 2020-10-05

- Add `-excludes` flag.
- Update default number of parameters allowed to `3` (from `2`).

## [0.1.4] - 2020-10-05

- Add `-verbose` flag.

## [0.1.3] - 2020-10-04

- Add `-public-max` and `-private-max` flags.

## [0.1.2] - 2020-09-22

- Improve program output.
- Remove `-help` flag.

## [0.1.1] - 2020-09-20

- Fix bug where multiple parameters of the same type where counted as one.

## [0.1.0] - 2020-09-20

- Analyse Go code to find function definitions with too many parameters.
- Configure the number of parameters allowed ~~(default is `2`)~~.
- Use the `-set_exit_status` flag to exit with a non-zero exit code if any
  function exceeds the limit.

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
