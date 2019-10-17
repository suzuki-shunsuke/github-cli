# github-cli

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
