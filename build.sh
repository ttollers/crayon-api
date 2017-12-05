#!/usr/bin/env bash

set -a

source .env

mkdir temp
cp $CHAPTER_DIR/translated_final.json ./temp/translated_final.json
cp $CHAPTER_DIR/native_final.json ./temp/native_final.json

docker build -t ttollers/crayon-api ./

rm -r temp
