#!/bin/sh
set -eu

root_dir=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
output_dir=${1:-"$root_dir/build/capi"}
mkdir -p "$output_dir"

cd "$root_dir"
CGO_ENABLED=1 go build -buildmode=c-archive -o "$output_dir/libparking_capi.a" ./cmd/parking-capi
CGO_ENABLED=1 go build -buildmode=c-shared -o "$output_dir/libparking_capi.so" ./cmd/parking-capi

cc -I"$output_dir" examples/c-link/main.c "$output_dir/libparking_capi.a" -o "$output_dir/c-link-static" -lpthread
cc -I"$output_dir" examples/c-link/main.c -L"$output_dir" -lparking_capi -o "$output_dir/c-link-dynamic"

"$output_dir/c-link-static"
case "$(uname -s)" in
Darwin) DYLD_LIBRARY_PATH="$output_dir" "$output_dir/c-link-dynamic" ;;
*) LD_LIBRARY_PATH="$output_dir" "$output_dir/c-link-dynamic" ;;
esac
