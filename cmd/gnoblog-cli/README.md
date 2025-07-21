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
- Batch post multiple `README.md` files by recursive looking through directories,
finding `README.md`s, and pack all of them to a single transaction to be 
broadcast to the chain.


  



