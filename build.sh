#!/bin/bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0 
find $GITHUB_WORKSPACE/  -name "go.mod" | find $PWD -name "go.mod" | rev |cut -c7- | rev > gomods
cat gomods | while read line 
do
   export MY_PATH=$line && cd $MY_PATH  && go build .
done
