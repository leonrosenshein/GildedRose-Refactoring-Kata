#!/bin/bash

set -euo pipefail

if [ ! -d "venv" ]; then
    python -m venv venv
fi
pushd go
  go build texttest_fixture.go
popd
venv/bin/pip install texttest
venv/bin/texttest -d . -con "$@"
