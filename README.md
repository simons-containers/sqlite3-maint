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

Repository contents (e.g., `Containerfile`, build scripts, and configuration) are licensed under the **MIT License**.

Software included in built container images (such as **musl**, **tzdata**, and **Mozilla CA Certificates**) are provided under their respective upstream licenses and is not covered by the MIT license for this repository.

## Acknowledgements

This project depends on a number of upstream components and data sources:

- **musl** – Lightweight C standard library implementation for Linux providing the standard C runtime and POSIX interfaces with a focus on simplicity, correctness, and static linking.  
  https://musl.libc.org/

- **tzdata** – The IANA Time Zone Database, which provides the canonical global timezone definitions used for correct time handling.  
  https://www.iana.org/time-zones

- **Mozilla CA Certificates** – The curated set of trusted root Certificate Authorities maintained by Mozilla and used by many systems for TLS verification.  
  https://wiki.mozilla.org/CA
