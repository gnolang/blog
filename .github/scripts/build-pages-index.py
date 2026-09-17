#!/usr/bin/env python3
"""Generate posts/index.html for the GitHub Pages deployment.

The Pages site exists to host post assets: 42 image references inside published
posts point at https://gnolang.github.io/blog/<post-dir>/src/..., and those posts
are live on chain, so the paths under posts/ must keep resolving exactly as they
do today. Without an index.html the site root simply returned 404 while every
asset underneath served fine.

This writes an index at the root of the uploaded artifact. It is generated at
deploy time and never committed.
"""

import html
import pathlib
import re
import sys

CHAIN_URL = "https://gno.land/r/gnoland/blog"

FRONT_MATTER = re.compile(r"\A---\n(.*?)\n---\n", re.DOTALL)


def read_post(readme):
    """Return (date, slug, title) for a post, or None if it has no front matter."""
    text = readme.read_text(encoding="utf-8", errors="replace").replace("\r\n", "\n")

    match = FRONT_MATTER.match(text)
    if not match:
        return None

    fields = {}
    for line in match.group(1).split("\n"):
        key, sep, value = line.partition(":")
        if sep:
            fields[key.strip()] = value.strip()

    slug = fields.get("slug")
    if not slug:
        return None

    title = slug
    for line in text[match.end():].split("\n"):
        if line.startswith("# "):
            title = line[2:].strip().strip("*")
            break

    return fields.get("publication_date", ""), slug, title


def main():
    posts_dir = pathlib.Path(sys.argv[1] if len(sys.argv) > 1 else "posts")

    entries = []
    for readme in sorted(posts_dir.glob("*/README.md")):
        post = read_post(readme)
        if post is None:
            print(f"skipping {readme}: no usable front matter", file=sys.stderr)
            continue

        date, slug, title = post
        entries.append((date, slug, title, readme.parent.name))

    entries.sort(reverse=True)

    rows = "\n".join(
        "      <li>"
        f'<time>{html.escape(date[:10])}</time> '
        f'<a href="{CHAIN_URL}:p/{html.escape(slug)}">{html.escape(title)}</a> '
        f'<a class="raw" href="{html.escape(directory)}/README.md">source</a>'
        "</li>"
        for date, slug, title, directory in entries
    )

    index = posts_dir / "index.html"
    index.write_text(PAGE.format(count=len(entries), rows=rows), encoding="utf-8")
    print(f"wrote {index} with {len(entries)} posts")


PAGE = """<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>gno.land blog — sources and assets</title>
    <style>
      body {{ margin: 0 auto; padding: 2rem 1rem; max-width: 46rem;
             font-family: system-ui, sans-serif; line-height: 1.5; }}
      ul {{ list-style: none; padding: 0; }}
      li {{ padding: .35rem 0; border-bottom: 1px solid #eee; }}
      time {{ font-variant-numeric: tabular-nums; color: #666; margin-right: .5rem; }}
      .raw {{ font-size: .8em; color: #666; }}
    </style>
  </head>
  <body>
    <h1>gno.land blog — sources and assets</h1>
    <p>
      This site hosts the raw sources and images for the gno.land blog. The blog itself
      lives on chain: <a href="{chain}">{chain}</a>.
    </p>
    <p>{count} posts.</p>
    <ul>
{rows}
    </ul>
  </body>
</html>
""".replace("{chain}", CHAIN_URL)


if __name__ == "__main__":
    main()
