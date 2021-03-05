# flaver

![build](https://github.com/tamada/flaver/workflows/build/badge.svg)

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?logo=spdx)](https://github.com/tamada/flaver/blob/main/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-blue.svg)](https://github.com/tamada/flaver/releases/tag/v1.0.0)

[![Docker](https://img.shields.io/badge/Docker-ghcr.io%2Ftamada%2Fflave%3A1.0.0-green?logo=docker)](https://github.com/users/tamada/packages/container/package/flaver)
[![tamada/brew/flaver](https://img.shields.io/badge/Homebrew-tamada%2Fbrew%2Fflaver-green?logo=homebrew)](https://github.com/tamada/homebrew-brew)

[![Discussion](https://img.shields.io/badge/GitHub-Discussion-orange?logo=GitHub)](https://github.com/tamada/flaver/discussions)

Find LAtest VERsions and their release dates of specified products from GitHub etc.

## :speaking_head: Overview

## :runner: Usage

```shell
Usage:
  flaver [PRODUCTs...] [flags]

Examples:
flaver tamada/flaver

Flags:
  -a, --all       finds all versions of the products
  -h, --help      help for flaver
  -v, --version   version for flaver
ARGUMENTS
  PRODUCTs        specified product names such as github.com/tamada/flaver, or tamada/flaver.
                  If no product names were specified, flaver runs on interactive mode.
```

### :beer: Homebrew

```shell
$ brew tap tamada/brew
$ brew install flaver
```
