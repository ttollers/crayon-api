#!/usr/bin/env bash

set -a

source .env

mkdir temp
cp $CHAPTER_DIR/harry-potter-1/fr/chapter-1/en.json ./temp/en.json
cp $CHAPTER_DIR/harry-potter-1/fr/chapter-1/fr.json ./temp/fr.json

docker build -t ttollers/crayon-api ./

rm -r temp
