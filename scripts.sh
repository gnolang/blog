#!/bin/sh
set -e

PASSWORD=${PASSWORD:-undefined}

maketx_call() {
    NET=${NET:-test3}
    REMOTE=${REMOTE:-$NET.gno.land:36657}
    KEY=${KEY:-moula}
    GAS_FEE=${GAS_FEE:-1000000ugnot}
    GAS_WANTED=${GAS_WANTED:-2000000}
    SEND=${SEND:-}

    echo "gnokey maketx call ${KEY} $@"
    echo "${PASSWORD}" | gnokey maketx call ${KEY} --chainid "${NET}" --remote "${REMOTE}" --gas-fee "${GAS_FEE}" --gas-wanted "${GAS_WANTED}" --send "${SEND}" --broadcast true --insecure-password-stdin true "$@"
}

echo "+ gnokey list | grep moula"
gnokey list | grep moula\ \(

set -e
#SEND=100ugnot maketx_call --pkgpath "gno.land/r/demo/banktest" --func "Deposit" --args "ugnot" --args "1"

SLUG="intro"
TITLE="Intro to Gnoland - The Smart Contract Platform to Improve Our Understanding of the World"
POSTFILE="blog/1-intro.md"
TAGS="gnoland,gnosh,gnot,permissionless,consensus,proof-of-contribution,dao,governance,ibc,democracy,freedom"
#maketx_call --pkgpath "gno.land/r/gnoland/blog" --func "ModAddPost"  --args "$SLUG" --args "$TITLE" --args#file "$POSTFILE" --args "$TAGS"
#maketx_call --pkgpath "gno.land/r/gnoland/blog" --func "ModEditPost" --args "$SLUG" --args "$TITLE" --args#file "$POSTFILE" --args "$TAGS"

SLUG="tech-ama1"
TITLE="Gno.land Community Technical AMA #1 - Recap"
POSTFILE="blog/2-tech-ama1.md"
TAGS="gnoland,gnosh,gnot,permissionless,consensus,proof-of-contribution,dao,governance,ibc,democracy,freedom"
#maketx_call --pkgpath "gno.land/r/gnoland/blog" --func "ModAddPost"  --args "$SLUG" --args "$TITLE" --args#file "$POSTFILE" --args "$TAGS"
#maketx_call --pkgpath "gno.land/r/gnoland/blog" --func "ModEditPost" --args "$SLUG" --args "$TITLE" --args#file "$POSTFILE" --args "$TAGS"


SLUG="gor-launch"
TITLE="Game of Realms Is On: Win Rewards for Contributing to Gno.land"
POSTFILE="blog/3-gor-launch.md"
TAGS="gnoland,game-of-realms,launch"
#maketx_call --pkgpath "gno.land/r/gnoland/blog" --func "ModAddPost"  --args "$SLUG" --args "$TITLE" --args#file "$POSTFILE" --args "$TAGS"
maketx_call --pkgpath "gno.land/r/gnoland/blog" --func "ModEditPost" --args "$SLUG" --args "$TITLE" --args#file "$POSTFILE" --args "$TAGS"
