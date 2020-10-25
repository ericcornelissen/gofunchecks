# Contributing Guidelines

The *gofunchecks* project welcomes contributions and corrections of all forms.
This includes improvements to the documentation or code base, new tests, bug
fixes, and implementations of new features. We recommend opening an issue before
making any significant changes so you can be sure your work won't be rejected.
But for changes such as fixing a typo you can open a Pull Request directly.

Before you continue, please do make sure to read through the relevant sections
of this document.

In this document you can read about:

- [Bug Reports](#bug-reports)
  - [Security reports](#security-reports)
- [Feature Requests](#feature-requests)
- [Workflow](#workflow)
  - [Prerequisites](#prerequisites)

---

## Bug Reports

If you have problems with any of the *gofunchecks* tools or think you've found a
bug, please report it to the developers; we cannot promise to do anything but we
might well want to fix it.

Before reporting a bug, make sure you've actually found a real bug. Carefully
read the documentation and see if it really says you can do what you're trying
to do. If it's not clear whether you should be able to do something or not,
report that too; it's a bug in the documentation! Also, make sure the bug has
not already been reported.

When preparing to report a bug, try to isolate it to a small working example
that reproduces the problem. Then, create a bug report including this example
and its results as well as any error or warning messages. Please don't
paraphrase these messages: it's best to copy and paste them into your report.
Finally, be sure to explain what you expected to happen; this will help us
decide whether it is a bug or a problem with the documentation.

Once you have a precise problem you can report it online as a [Bug Report].

### Security Reports

For more information on how to report security-related issues see the
[SECURITY.md] document.

---

## Feature Requests

If you require a new feature from one of the *gofunchecks* tools you can request
for it to be added. If you require an entirely new *gofunchecks* tool you can
request it as well.

Before submitting a feature request, make sure you've checked if what you want
to achieve isn't already possible. Carefully read the documentation and try to
get the tool to do what you want. If it is possible, but not (clearly)
documented, report that too; it's a gap in the documentation (or unintended
behaviour). Also, make sure the feature has not already been requested.

When preparing to submit a feature request, take a moment to consider if your
situation is generally applicable. Try to make the feature request generic so
that it is not only useful to your situation but other situations as well. Be
sure to explain in detail why the feature is useful and, if possible, how it
should work.

Once you have a precise request you can report it online as a [Feature Request].

---

## Workflow

If you decide to make a contribution, please do use the following workflow:

- Fork the repository.
- Create a new branch from the latest `main`.
- Make your changes on the new branch.
- Run `go fmt ./...` and commit to the new branch.
- Push your changes and open a Pull Request.

### Prerequisites

The prerequisites for contributing to this project are:

- Go; version `1.14` or higher
- Git

Furthermore we recommend installing the following tools to analyze your changes
before opening a Pull Request:

- [gocyclo]
- [golint]
- [misspell]

[Bug Report]: https://github.com/ericcornelissen/gofunchecks/issues/new?labels=bug&template=bug_report.md
[Feature Request]: https://github.com/ericcornelissen/gofunchecks/issues/new?labels=enhancement&template=feature_request.md
[gocyclo]: https://github.com/fzipp/gocyclo#readme
[golint]: https://github.com/golang/lint#readme
[misspell]: https://github.com/client9/misspell#readme
[SECURITY.md]: ./SECURITY.md
