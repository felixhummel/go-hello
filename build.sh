#!/bin/bash
set -euo pipefail

for d in "hello-simple/ hello-stdin/"; do
  (cd $d && go install)
done
