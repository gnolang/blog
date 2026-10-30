# `gnoblog-cli`

`gnoblog-cli` is a configurable tool that allows for easy deployment of blog posts to the
`r/gnoland/blog` realm.

All blog posts posted with `gnoblog-cli` need to be in the correct format:
- Filename: `README.md`,
- Frontmatter must be of a specific format,
- There must be an H1 in below the frontmatter, which will be treated as the post title.

For reference, see existing posts, like [Peace!](../../posts/2022-05-02_peace)

`gnoblog-cli` utilizes keys from the local Gno keybase to deploy to
the blog realm, whose path is configurable via the `--pkgpath` flag.

`gnoblog-cli` can has the ability to:
- Post a single blog post `README.md` file to the chain, in which case
it takes the direct path to the file as an argument,
- Batch post multiple `README.md` files by recursively looking through directories,
finding `README.md`s, and packing them into transactions to be broadcast to the chain.

## Quick start (recommended)

You only need your key, a target website, and the posts directory:

```sh
go run ./cmd/gnoblog-cli post posts/ --key moul --target gno.land
```

With `--target`, the CLI fetches the website's `gnoconnect` `<meta>` tags to
discover the RPC endpoint and chain ID automatically, so you don't have to pass
`--remote` and `--chainid` by hand.

Gas and batching are handled for you:
- **Gas is auto-estimated** per transaction via on-chain simulation (a `1.3x`
  safety multiplier is applied; tune with `--gas-adjustment`). The fee is derived
  from `--gas-price` (defaults to the mainnet floor `1ugnot/1000gas`).
- **Posts are auto-paged** into multiple transactions that stay under the chain's
  `MaxTxBytes` limit, so uploading the whole archive at once never hits the tx-size
  wall. Already-published posts are skipped.

Add `--dry-run` to preview the transaction plan (batches, gas, fees) without
broadcasting anything:

```sh
go run ./cmd/gnoblog-cli post posts/ --key moul --target gno.land --dry-run
```

## Manual configuration

You can still set everything explicitly instead of using `--target`:

```sh
go run ./cmd/gnoblog-cli post posts/ \
  --key moul \
  --chainid gnoland-1 \
  --remote https://rpc.gno.land \
  --pkgpath gno.land/r/gnoland/blog
```

Gas can also be pinned instead of estimated: pass `--gas-wanted` (a value `> 0`
disables auto-estimation) and/or `--gas-fee` (a flat fee that overrides
`--gas-price`).

## Useful flags

| Flag | Default | Description |
| --- | --- | --- |
| `--key` | — | name of the keypair to sign with (required) |
| `--target` | — | gno.land web endpoint to auto-discover `--remote` and `--chainid` (e.g. `gno.land`) |
| `--remote` | `localhost:26657` | RPC node URL (auto-set by `--target`) |
| `--chainid` | `dev` | chain ID (auto-set by `--target`) |
| `--pkgpath` | `gno.land/r/gnoland/blog` | blog realm path |
| `--gas-wanted` | `0` | gas per tx; `0` = auto-estimate via simulation |
| `--gas-adjustment` | `1.3` | multiplier applied to the simulated gas estimate |
| `--gas-price` | `1ugnot/1000gas` | gas price used to derive the fee when `--gas-fee` is unset |
| `--gas-fee` | — | flat gas fee that overrides `--gas-price` |
| `--dry-run` | `false` | simulate and print the tx plan without broadcasting |
| `--edit` | `false` | edit existing on-chain posts instead of adding new ones |
