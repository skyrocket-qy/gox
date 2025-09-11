#!/bin/bash
set -e
echo "mode: count" > coverage.out
i=0
for dir in $(go list -f '{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}' ./...)
do
    go test -v -covermode=count -coverprofile=profile.${i}.out $dir
    if [ -f profile.${i}.out ]; then
        tail -n +2 profile.${i}.out >> coverage.out
    fi
    i=$((i+1))
done

go tool cover -func=coverage.out
rm profile.*.out
