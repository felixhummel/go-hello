#!/bin/bash
set -euo pipefail

go install
cat <<'EOF' | hello-stdin
foo
bar
baz
EOF
