#!/bin/sh
awk '/Change/{flag=1}/General information/{flag=0}' README.md > ChangeLog.md
curl -s -O 'http://roboticoverlords.org/images/default.png'
