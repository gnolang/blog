#!/bin/sh

MAX_WIDTH=800

for dir in */; do
    if [ -d "${dir}src" ]; then
        if [ ! -d "${dir}src/thumbs" ]; then
            mkdir -p "${dir}src/thumbs"
            cd "${dir}src" || continue

             for img in *.png *.jpg; do
                if [ -f "$img" ]; then
                    width=$(gm identify -format "%w" "$img")
                    if [ "$width" -gt "$MAX_WIDTH" ]; then
                        gm convert "$img" -resize "${MAX_WIDTH}x" -quality 95 "thumbs/$img" 2>/dev/null
                        if [ $? -ne 0 ]; then
                            echo "Error processing $img"
                        fi
                    else
                        echo "$img is already less than ${MAX_WIDTH} pixels wide, skipping..."
                    fi
                fi
            done

            cd ../..
            echo "Thumbnails created in ${dir}src/thumbs"
        fi
    fi
done