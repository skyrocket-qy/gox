#!/bin/bash
set -e
echo "mode: set" > coverage.out
i=0
for dir in \
    ./dsa/graph/find-bridge \
    ./dsa/graph/a_star \
    ./dsa/ds/radix-tree \
    ./gormx/lib/scope \
    ./qrcode \
    ./phone \
    ./auth \
    ./structx \
    ./redisx \
    ./redisx/slidewindow \
    ./common \
    ./refletx \
    ./logx \
    ./probfilter/skiplist \
    ./probfilter/xorfilter \
    ./probfilter/cuckoofilter \
    ./probfilter/minhashlsh \
    ./probfilter/quotientfilter \
    ./stringx \
    ./excel \
    ./periodcheck \
    ./middleware/connectw \
    ./middleware/ginw \
    ./body \
    ./httpx
do
    go test -coverprofile=profile.${i}.out $dir
    if [ -f profile.${i}.out ]; then
        tail -n +2 profile.${i}.out >> coverage.out
    fi
    i=$((i+1))
done

go tool cover -func=coverage.out
rm profile.*.out
