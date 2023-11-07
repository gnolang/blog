#!/bin/sh -x

for dir in */; do
    if [ -d "${dir}src" ]; then
        mkdir -p "${dir}src/thumbs"
        cd "${dir}src" || continue

        for img in *.png *.jpg; do
            sips --resampleWidth 600 "$img" --out "thumbs/$img"
        done

        cd ../..
    fi
done