#!/bin/sh
if [ -f "package.json" ]; then
    if [ ! -d "node_modules" ]; then
        npm install
    fi
fi

exec "$@"
