#!/bin/bash

LAST_COMMIT_HASH=`git log --pretty=oneline | head -1 | tr " " "\n" | head -1`
DIR_PROJECT=`pwd`

cd ./argparser
go get github.com/terryhay/dolly/utils@$LAST_COMMIT_HASH
cd $DIR_PROJECT

cd ./generator
go get github.com/terryhay/dolly/utils@$LAST_COMMIT_HASH
go get github.com/terryhay/dolly/argparser@$LAST_COMMIT_HASH
cd $DIR_PROJECT

cd ./examples
go get github.com/terryhay/dolly/utils@$LAST_COMMIT_HASH
go get github.com/terryhay/dolly/argparser@$LAST_COMMIT_HASH
go get github.com/terryhay/dolly/man_style_help@$LAST_COMMIT_HASH
cd $DIR_PROJECT