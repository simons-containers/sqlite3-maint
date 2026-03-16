# Sidecar sqlite3 maintenance container

Sidecar container for running maintenance on sqlite3 databases as an initContainer.

## Running

Vacuum and analyze will be run on Sqlite3 files added with `--db` and found in directories specified with `--dbdir`

Example:

```bash
docker run -it --rm -v data:/data \
  ghcr.io/simons-containers/sqlite3-maint \
  --dbdir /var/lib/emby/data
```

## Building

| Arg | Description |
|---|---|
| `VERSION` | Current version tag

Build container:

```bash
docker build \
  -t sqlite3-maint:${VERSION} \
  --build-arg VERSION=${VERSION} \
  -f Containerfile .
```

## License

Repository contents are licensed under the **MIT License**.
