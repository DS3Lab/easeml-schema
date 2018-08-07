#!/bin/sh

set -e

if [ ! -f "src/build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Set up the environment to use the workspace.
GOPATH="$PWD"
export GOPATH

# Launch the arguments with the configured environment.
exec "$@"
