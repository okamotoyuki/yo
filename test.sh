#!/usr/bin/env bash

try() {
    expected="$1"
    input="$2"

    rm -f temp/*
    mkdir -p temp
    cp test/main.go temp
    ./yo "$input" > temp/temp.s
    cd temp
    go build
    actual=`./temp`

    if [ "$actual" = "$expected" ]; then
      echo "$input => $actual"
    else
      echo "$expected expected, but got $actual"
      exit 1
    fi

    cd ..
}

try 0 0
try 42 42

echo OK
