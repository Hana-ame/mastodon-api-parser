#!/bin/bash

shopt -s extglob
mkdir text/$1
# mv `\ls -1 . | grep -v "(index.js|index.ts|json|move_txtfiles.sh|node_modules|package.json|package-lock.json|readme.md|text|tsconfig.json)"` ./text/$1

mv *.txt text/$1