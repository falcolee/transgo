#!/bin/bash

appName="transgo"
builtAt="$(date +'%F %T %z')"
goVersion=$(go version | sed 's/go version //')
gitAuthor="skylzl001@163.com"

if [ "$1" == "release" ]; then
  gitTag=$(git describe --abbrev=0 --tags)
else
  gitTag=build-next
fi

echo "build version: $gitTag"

ldflags="\
-w -s \
-X 'github.com/falcolee/transgo/common.BuiltAt=$builtAt' \
-X 'github.com/falcolee/transgo/common.GoVersion=$goVersion' \
-X 'github.com/falcolee/transgo/common.GitAuthor=$gitAuthor' \
-X 'github.com/falcolee/transgo/common.GitTag=$gitTag' \
"

if [ "$1" == "release" ]; then
  gox -osarch="linux/386 linux/amd64 darwin/amd64 darwin/arm64" -output="transgo-{{.OS}}-{{.Arch}}" -ldflags="$ldflags"
else
  go build -o "transgo" -ldflags="$ldflags"
fi

mkdir -p "bin"
mv transgo-* bin
cd bin || exit
# compress file (release)
if [ "$1" == "release" ]; then
    mkdir -p compress
    rm compress/*.tar.gz
    rm compress/*.zip
    for i in `find . -type f -name "$appName-linux-*"`
    do
      tar -czvf compress/"$i".tar.gz "$i"
    done
    for i in `find . -type f -name "$appName-darwin-*"`
    do
      tar -czvf compress/"$i".tar.gz "$i"
    done
    for i in `find . -type f -name "$appName-windows-*"`
    do
      zip compress/$(echo $i | sed 's/\.[^.]*$//').zip "$i"
    done
fi