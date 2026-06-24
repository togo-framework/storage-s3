<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/storage-s3</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/storage-s3"><img src="https://pkg.go.dev/badge/github.com/togo-framework/storage-s3.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/storage-s3
```

<!-- /togo-header -->

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

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
