# rerun [![CircleCI](https://circleci.com/gh/shin1x1/rerun.svg?style=svg)](https://circleci.com/gh/shin1x1/rerun)

> A tool that automatically reruns a exited commands

`rerun` is a tool written in Go that automatically restarts the exited command and keeps running. It's like a simplified supervisor on a development PC.

```
![screencast](docs/screencast.svg)
```

## Installation

Download binary from below link.

<https://github.com/shin1x1/rerun/releases>

You can install using `go get`.

```
go get -u github.com/shin1x1/rerun
```

## Usage

Specifies the command you want to run after `rerun`. Reruns the command after 10 seconds when the specified command exited. You can change the number of seconds in the interval with `-s `.

```bash
# Reruns the command after 10 seconds
$ rerun ./sample.sh

# Reruns the command after a second
$ rerun -s 1 ./sample.sh
```

If the command contains `-`, add `--` before the command.

```bash
$ rerun kubectl get events -w
```
