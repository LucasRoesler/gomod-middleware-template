ARG GOTAG=1.15-alpine3.12
ARG WATCHDOGVERSION=0.8.0
ARG APLINETAG=3.12

FROM --platform=${TARGETPLATFORM:-linux/amd64} openfaas/of-watchdog:${WATCHDOGVERSION} as watchdog
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:${GOTAG} as build

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

RUN apk --no-cache add git

ENV CGO_ENABLED=0

RUN mkdir -p /go/src/function
WORKDIR /go/src/function
COPY ./function .

# Run a gofmt and exclude all vendored code.
RUN test -z "$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./function/vendor/*"))" || { echo "Run \"gofmt -s -w\" on your Golang code"; exit 1; }

ARG GO111MODULE="on"
ARG GOPROXY=""
ARG GOFLAGS=""

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go test ./... -cover
RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build --ldflags "-s -w" -a -installsuffix cgo -o function_handler .

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:${APLINETAG}
# Add non root user and certs
RUN apk --no-cache add ca-certificates \
    && addgroup -S app && adduser -S -g app app \
    && mkdir -p /home/app \
    && chown app /home/app

WORKDIR /home/app

COPY --from=build --chown=app /go/src/function/function_handler    handler
COPY --from=watchdog --chown=app /fwatchdog .

USER app

ENV fprocess="./handler"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8082"

CMD ["./fwatchdog"]
