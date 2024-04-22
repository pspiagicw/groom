# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased(v0.0.2)

### Added

- Added the `directory` option to change directory before executing a task.
- Help Printing.
- Added a `--simple` flag to task listing.
- Added parent directory recursion. `groom` now finds `groom.toml` in the parent directories.
- Created a companion `neovim` plugin. See [here](https://github.com/pspiagicw/groom.nvim)

### Changed
- Moved to [`demp`](https://github.com/pspiagicw/demp) for dollar templating instead of hideos custom algorithm.
- Moved to [`shellwords`](https://github.com/buildkite/shellwords) for splitting shell commands instead of custom algorithm.

## v0.0.1

### Added

- Added pretty listing of tasks.

### Changed

- Moved location of main.go.

