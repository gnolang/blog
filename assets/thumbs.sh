#!/bin/sh -xe

for dir in */; do (
    cd $dir
    mkdir -p thumbs

    for img in *.png *.jpg; do
        convert $img -resize 600x ./thumbs/$img
    done
); done
