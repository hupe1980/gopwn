#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run $(ls -1 cmd/*.go | grep -v _test.go) completion "$sh" >"completions/gopwn.$sh"
done