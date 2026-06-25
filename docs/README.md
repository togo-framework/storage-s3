# storage-s3 — documentation

togo storage driver

## Overview

Package s3store is an AWS S3 driver for togo storage. It implements
togo.Storage and overrides the default filesystem storage when installed.
Blank-import + set S3_BUCKET (and the standard AWS_* credentials/region).


Env: S3_BUCKET (required), AWS_REGION (default us-east-1), AWS_ACCESS_KEY_ID,

## Install

```bash
togo install togo-framework/storage-s3
```

Set `STORAGE_DRIVER=s3`.

## Configuration

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

| Env var |
|---|
| `S3_BUCKET"` |

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
- Full README: ../README.md
