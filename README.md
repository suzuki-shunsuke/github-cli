# github-cli

[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/github-cli/status.svg)](https://cloud.drone.io/suzuki-shunsuke/github-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/github-cli)](https://goreportcard.com/report/github.com/suzuki-shunsuke/github-cli)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/github-cli.svg)](https://github.com/suzuki-shunsuke/github-cli)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/github-cli/master/LICENSE)

Command line tool to call GitHub v3 API.

https://developer.github.com/v3/

We created this tool to delete repositories's subscription, so basically we don't support other unneeded APIs.

## Install

Download the binary from [release page](https://github.com/suzuki-shunsuke/github-cli/releases).

## Usage

```console
$ github help
```

## Tips: Delete subscriptions of repositories

```console
$ export GITHUB_TOKEN=<your github token>
$ github repos list-org <organization name> | xargs github watch delete
```

## License

[MIT](LICENSE)
