# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to [Semantic
Versioning].

## [Unreleased]

- _No changes yet_

## [1.0.0] - 2020-10-25

- Analyse Go code to find function definitions with too many return values.
- Use the `-max` flag to configure the number of return values allowed (default
  is `2`).
- Use the `-public-max` and `-private-max` flags to .
- Use the `-tests` flag to analyze test files.
- Use the `-excludes` flag to exclude files based on a custom pattern from the
  analysis.
- Use the `-verbose` flag to get more verbose output.
- Use the `-set_exit_status` flag to exit with a non-zero exit code if any
  function exceeds the limit.

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
