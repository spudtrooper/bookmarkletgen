#!/bin/sh
#
# Runs the index generator.
#
set -e

rm -rf output
mkdir -p output
args=(
    --base_source_url
    https://github.com/spudtrooper/bookmarklets/blob/main/js
    --js_dir
    ../bookmarklets/js
    --outfile_html
    output/bookmarklets.html
    --outfile_md
    output/bookmarklets.md
    --footer_html
    "[<a href=\"https://github.com/spudtrooper/bookmarklets\">Source</a>]"
)
go run main.go "${args[@]}" 2> /dev/null

echo "OK"
