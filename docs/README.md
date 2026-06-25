# storage-s3 — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package s3store is an AWS S3 driver for togo storage. It implements
togo.Storage and overrides the default filesystem storage when installed.
Blank-import + set S3_BUCKET (and the standard AWS_* credentials/region).

	togo install togo-framework/storage-s3

## Install

```bash
togo install togo-framework/storage-s3
```

Set `STORAGE_DRIVER=s3`.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `G` | _see provider docs_ |
| `S3_BUCKET` | _see provider docs_ |

## Usage

```go
st := k.Storage
st.Put(ctx, "path/file.txt", data)
b, _ := st.Get(ctx, "path/file.txt")
url := st.Path("path/file.txt")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/storage-s3
- README: ../README.md
