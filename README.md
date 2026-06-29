# Sidecar sqlite3 maintenance container

Sidecar container for running maintenance on sqlite3 databases as an initContainer.

## Running

Vacuum and analyze will be run on Sqlite3 files added with `--db` and found in directories specified with `--dbdir`

Example:

```bash
docker run -it --rm -v /path/to/sqlite/dbs:/data \
  ghcr.io/simons-containers/sqlite3-maint \
  --dbdir /data
```

## Building

| Arg | Description |
|---|---|
| `VERSION` | Current version tag

Build container using build-args from versions.yaml:

```bash
docker build -t \
  distroless-traefik:$(yq -r .traefik versions.yaml) \
  $(yq -r 'to_entries[] | "--build-arg " + (.key | upcase) + "_VERSION=" + .value' versions.yaml) \
  -f Containerfile .
```

## License

Repository contents are licensed under the **MIT License**.
