#!/bin/sh
#
# Create release tarballs/zip for 64-bit linux, BSD and Plan9 + 64-bit ARM + raspberry pi 2/3
#

# Download the default icon
curl -s -O 'https://roboticoverlords.org/images/default.png'

# Generate a changelog from the entries in the readme
awk '/Change/{flag=1}/General information/{flag=0}flag' README.md > CHANGELOG.md

# Name and version
name=gendesk
version=$(grep version main.go | head -1 | sed 's/\"//g' | tr ' ' '\n' | tail -1)
echo "$name $version"

# Create a source tarball
mkdir "$name-$version"
cp -r $name.1 *.go go.mod go.sum LICENSE README.md CHANGELOG.md default.png vendor "$name-$version/"
gzip "$name-$version/$name.1"
mkdir -p dist
tar Jcf "dist/$name-$version.tar.xz" "$name-$version/"
rm -r "$name-$version"

# Build executables
echo 'Compiling...'
echo '* Linux x86_64'
GOARCH=amd64 GOOS=linux go build -mod=vendor -o $name.linux
echo '* macOS x86_64'
GOARCH=amd64 GOOS=darwin go build -mod=vendor -o $name.macos
echo '* FreeBSD x64_64'
GOARCH=amd64 GOOS=freebsd go build -mod=vendor -o $name.freebsd
echo '* NetBSD x86_64'
GOARCH=amd64 GOOS=netbsd go build -mod=vendor -o $name.netbsd
echo '* Linux ARM64'
GOARCH=arm64 GOOS=linux go build -mod=vendor -o $name.linux_arm64
echo '* Raspberry Pi 2/3'
GOARCH=arm GOARM=6 GOOS=linux go build -mod=vendor -o $name.rpi
echo '* Linux x86_64 static and compressed'
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -mod=vendor -v -trimpath -ldflags "-s" -a -o $name.linux_static && upx $name.linux_static

# Compress the Linux releases with xz
for p in linux linux_arm64 rpi linux_static; do
  echo "Compressing $name-$version.$p.tar.xz"
  mkdir "$name-$version-$p"
  cp $name.1 "$name-$version-$p/"
  gzip "$name-$version-$p/$name.1"
  cp $name.$p "$name-$version-$p/$name"
  cp LICENSE README.md CHANGELOG.md default.png "$name-$version-$p/"
  mkdir -p dist
  tar Jcf "dist/$name-$version-$p.tar.xz" "$name-$version-$p/"
  rm -r "$name-$version-$p"
  rm $name.$p
done

# Compress the other tarballs with gz
for p in macos freebsd netbsd; do
  echo "Compressing $name-$version.$p.tar.gz"
  mkdir "$name-$version-$p"
  cp $name.1 "$name-$version-$p/"
  gzip "$name-$version-$p/$name.1"
  cp $name.$p "$name-$version-$p/$name"
  cp LICENSE README.md CHANGELOG.md default.png "$name-$version-$p/"
  mkdir -p dist
  tar zcf "dist/$name-$version-$p.tar.gz" "$name-$version-$p/"
  rm -r "$name-$version-$p"
  rm $name.$p
done

echo 'DONE!'
