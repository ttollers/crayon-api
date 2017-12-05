#!/usr/bin/env bash
CHAPTER_DIR=/Users/tom/Projects/personal/book-lib/harry-potter-1/fr/chapter-1

mkdir temp
cp $CHAPTER_DIR/translated_final.json ./temp/translated_final.json
cp $CHAPTER_DIR/native_final.json ./temp/native_final.json

docker build -t ttollers/crayon-api ./

rm -r temp
