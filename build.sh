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
  gox -osarch="linux/amd64 windows/amd64 darwin/amd64" -output="transgo-{{.OS}}_{{.Arch}}" -ldflags="$ldflags"
else
  go build -o "transgo" -ldflags="$ldflags"
fi

mkdir -p "bin"
mv transgo* bin
cd bin || exit
# compress file (release)
if [ "$1" == "release" ]; then
    mkdir compress
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