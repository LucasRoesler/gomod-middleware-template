# Openfaas Go Module middleware

[![GitHub Super-Linter](https://github.com/LucasRoesler/gomod-middleware-template/workflows/Lint%20and%20Test/badge.svg)](https://github.com/marketplace/actions/super-linter)

This repository contains a function template for OpenFaaS that is designed for Go modules.

The template is also designed so that your local Go environment will "just work"™️. It is structured
so that the generated function module is a fully functioning and compilable Go project. This means you
can use any standard Go tool from that folder and it will work as expected.

## Getting the template
```sh
$ faas-cli template pull https://github.com/LucasRoesler/gomod-middleware-template
$ faas-cli new --list

Languages available as templates:
- gomod-middleware
```


## Starting a new function

```sh
$ faas-cli template pull https://github.com/LucasRoesler/gomod-middleware-template
$ faas-cli new --lang gomod-middleware myfunc
Folder: myfunc created.
  ___                   _____           ____
 / _ \ _ __   ___ _ __ |  ___|_ _  __ _/ ___|
| | | | '_ \ / _ \ '_ \| |_ / _` |/ _` \___ \
| |_| | |_) |  __/ | | |  _| (_| | (_| |___) |
 \___/| .__/ \___|_| |_|_|  \__,_|\__,_|____/
      |_|


Function created in folder: myfunc
Stack file written: myfunc.yml

Notes:
You have created a new function which uses Golang 1.15 and modules.

The `main.go` is a generated file, all of your code should
go in the `handler` folder.

See more: https://docs.openfaas.com/cli/templates/
For detailed examples:
https://github.com/LucasRoesler/gomod-middleware-template
```

Then, update your stack YAML to include

```yaml
configuration:
  templates:
    - name: gomod-middleware
      source: https://github.com/LucasRoesler/gomod-middleware-template
```

Then edit the function logic in `myfunc/handler`.  The `main.go` is part of the template and should generally never be edited.


## Build Args
The template allows you to control version of Go, Alpine, and the Watchdog that are used in the build and final image.

Add this to your function YAML to control the versions
```yaml
build_args:
  GOTAG: 1.15-alpine3.12
  WATCHDOGVERSION: 0.8.0
  APLINETAG: 3.12
  CGO_ENABLED: 0
```


## Multi-arch support
You can also control the build and final architecture using the following build args, the defaults
are show.


```yaml
TARGETPLATFORM: linux/amd64
BUILDPLATFORM: linux/amd64
TARGETOS: ""
TARGETARCH: ""
```

The value of `TARGETOS` sets the value of `GOOS` during `go build`. Similarly, `TARGETARCH` sets the value of `GOARCH`.
