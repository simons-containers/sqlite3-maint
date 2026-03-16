FROM golang:1.22-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /src
ADD main.go /src/main.go
RUN go mod init sqlite3-maint
RUN go mod tidy
RUN CGO_ENABLED=1 \
  go build \
  -trimpath \
  -ldflags="-s -w -linkmode external -extldflags '-static'" \
  -o sqlite3-maint

FROM scratch
ARG VERSION

COPY --from=builder /src/sqlite3-maint /sqlite3-maint

ENTRYPOINT ["/sqlite3-maint"]
CMD ["--help"]

LABEL org.opencontainers.image.title="sqlite3-maint"
LABEL org.opencontainers.image.description="sidecar sqlite3 maintenance container"
LABEL org.opencontainers.image.source="https://github.com/simons-containers/sqlite3-maint"
LABEL org.opencontainers.image.version="${VERSION}"