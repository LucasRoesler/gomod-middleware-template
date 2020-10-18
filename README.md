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
