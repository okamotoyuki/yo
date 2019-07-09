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
try 2 "1 + 1"
try 6 "1 + 2 + 3"
try 8 "3+5"
try 6 "8 - 2"
try 0 "6 - 3 - 3"
try 4 "7-3"
try 14 "2 * 7"
try 24 "2 * 3 * 4"
try 5 "20 / 4"
try 1 "9/3/3"
try 13 "15 - 24 / 6 + 2"
try 8 "29-9/3*7"
try 1 "1 * (5 - 3) / 2"
try 10 "12-8/(22-18)"
try 1 "-1 + 2"
try 10 "+5 + +5"
try 5 "-1 + -(-2 + 4) + 8"
try 1 "1 == 1"
try 0 "1 != 1"
try 1 "2 < 6"
try 0 "8 < 4"
try 1 "5 > 3"
try 0 "1 > 9"
try 1 "3 <= 3"
try 0 "7 <= 6"
try 1 "11 >= 9"
try 1 "13 >= 13"

echo OK
