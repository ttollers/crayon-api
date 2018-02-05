#!/usr/bin/env bash

set -a

source .env

mkdir temp
cp $CHAPTER_DIR/harry-potter-1/fr/chapter-1/translated_final.json ./temp/translated_final.json
cp $CHAPTER_DIR/harry-potter-1/fr/chapter-1/native_final.json ./temp/native_final.json

docker build -t ttollers/crayon-api ./

rm -r temp
