#!/bin/sh
awk '/Change/{flag=1}/General information/{flag=0}flag' README.md > ChangeLog.md
curl -s -O 'http://roboticoverlords.org/images/default.png'
