#!/bin/bash
set -euo pipefail

here="$(readlink -m $(dirname $0))"

for d in hello-*/; do
  cd $here
  set -x
  cd $d
  gofmt -w *.go
  set +x
done
