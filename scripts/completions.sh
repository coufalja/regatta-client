#!/bin/sh
# scripts/completions.sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh; do
	go run main.go completion "$sh" >"completions/regatta-client.$sh"
done
