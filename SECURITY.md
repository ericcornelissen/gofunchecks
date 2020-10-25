# Security Policy

## Supported Versions

The tables below tells you which version of each *gofunchecks* tool is
currently being supported with security updates.

### goparamcount

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0.0 | :x:                |

### goreturncount

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |

## Reporting a Vulnerability

The maintainers of *gofunchecks* take security bugs very seriously. We
appreciate your efforts to responsibly disclose your findings. Due to the
non-funded open-source nature of this project, we take a best-efforts approach
when it comes to engaging with (security) reports.

To report a security issue, send an email to [security@ericcornelissen.dev] and
include the words _"SECURITY"_ and _"gofunchecks"_ in the subject line. Please
do not open a regular issue or Pull Request in the public repository.

## Scope

The following are **not** in scope:

- File paths provided as taint input: *gofunchecks* will only ever read `.go`
  files.

If you are unsure, send a report anyway.

## Acknowledgments

We would like to publicly thank the following repoters:

- _None yet_

[security@ericcornelissen.dev]: mailto:security@ericcornelissen.dev?subject=SECURITY%20%28gofunchecks%29
