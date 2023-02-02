# Reporting :detective:

This directory contains the code for the reporting system.\
For launching the reporting script, you will need a github & twitter token.

You can find these tokens in the following links:
- twitter: https://developer.twitter.com/en/docs/authentication/oauth-2-0
- github: https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token

## Usage

```bash
USAGE
  reporting [flags]

FLAGS
  -backlog=true           generate backlog
  -changelog=true         generate changelog
  -curation=true          generate curation
  -format markdown        output format
  -github-token ...       github token
  -help=false             show help
  -output-path ./output/  output directory path
  -since ...              since date RFC 3339 (ex: 2003-01-19T00:00:00.000Z)
  -tips=true              generate tips
  -twitter-token ...      twitter token
```

------------

## Example output

