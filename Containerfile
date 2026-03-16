FROM golang:1.22-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /src
ADD main.go /src/main.go
RUN go mod init sqlite3-maint
RUN go get github.com/mattn/go-sqlite3
RUN CGO_ENABLED=1 \
  go build \
  -trimpath \
  -ldflags="-s -w -linkmode external -extldflags '-static'" \
  -o sqlite3-maint

FROM scratch

COPY --from=builder /src/sqlite3-maint /sqlite3-maint

ENTRYPOINT ["/sqlite3-maint"]
CMD ["--help"]
