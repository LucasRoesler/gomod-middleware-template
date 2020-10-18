# Openfaas Go Module middleware

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
$ faas-cli new --lange gomod-middleware myfunc


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
