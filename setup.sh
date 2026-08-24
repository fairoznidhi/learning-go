#!/usr/bin/env bash
set -e

if command -v dlv >/dev/null 2>&1; then
	echo "delve already installed: $(dlv version | head -n1)"
else
	echo "delve not found, installing..."
	go install github.com/go-delve/delve/cmd/dlv@latest
	echo "delve installed. Make sure \$(go env GOPATH)/bin is on your PATH."
fi
