# storage-s3

An **AWS S3** driver for [togo](https://to-go.dev) storage. Implements `togo.Storage`
and overrides the default filesystem storage when installed — your app keeps using
`k.Storage.Put/Get/Delete/Path`, now backed by S3.

## Install

```bash
togo install togo-framework/storage-s3
```

## Configure (`.env`)

```ini
S3_BUCKET=my-bucket
AWS_REGION=us-east-1
AWS_ACCESS_KEY_ID=…
AWS_SECRET_ACCESS_KEY=…
```

Any default AWS credential source works (env vars, shared config, IAM role). If
`S3_BUCKET` is unset the plugin no-ops and the default storage stays active.

## How it works

Blank-importing the package registers an S3-backed `togo.Storage` provider at
`PriorityService+10`, so it overrides the built-in filesystem store. `Path()`
returns the public object URL.

MIT © togo-framework
