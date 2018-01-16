#!/bin/sh
# Generate a changelog from the entries in the readme
awk '/Change/{flag=1}/General information/{flag=0}flag' README.md > CHANGELOG.md
# Download the default icon
curl -s -O 'http://roboticoverlords.org/images/default.png'
