# `gnolang/blog`

_Disclaimer: This repo is highly experimental and subject to breaking changes._

Tools and content to manage blogposts on the [`r/gnoland/blog` realm](https://github.com/gnolang/gno/tree/master/examples/gno.land/r/gnoland/blog),
powered by the [`p/demo/blog` library](https://github.com/gnolang/gno/tree/master/examples/gno.land/p/demo/blog).

(Current) live version: https://test3.gno.land/r/gnoland/blog.

## Posts ([`./posts`](./posts))

The blog posts are written in markdown format, and include a frontmatter prefix to define metadata.

## Contributing a blog post

See [CONTRIBUTING.md](./CONTRIBUTING.md).

## CLI ([`./gnoblog-cli`](./gnoblog-cli))

_See [#1](https://github.com/gnolang/blog/issues/1)_

This is a standalone client for managing blog posts located in [`./posts`](./posts).
It functions similarly to [`gnokey`](https://github.com/gnolang/gno/tree/master/gno.land/cmd/gnokey), but is solely dedicated to blog content.

## Reporting scripts ([`./reporting`](./reporting))

You'll find scripts in this folder for generating reports based on contributions and activity.
