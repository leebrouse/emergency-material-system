#!/usr/bin/env bash

set -euo pipefail

shopt -s globstar

if ! [[ "$0" =~ script/genopenapi.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

source ./backend/script/lib.sh

API_ROOT="./backend/api"
OPENAPI_OUT="./backend/internal/common/genopenapi"

function openapi_files {
  openapi_files=$(find ./backend/api/openapi -type f -name '*.yaml')
  echo "${openapi_files[@]}"
}

function gen_for_openapi() {
  if [ -d "$OPENAPI_OUT" ]; then
    log_warning "found existing $OPENAPI_OUT, cleaning all files under it"
    run rm -rf $OPENAPI_OUT
  fi

  run mkdir -p "$OPENAPI_OUT"

  for openapi_file in $(openapi_files); do
    local service_name=$(basename "$openapi_file" .yaml)
    local out_dir="$OPENAPI_OUT/$service_name"

    run mkdir -p "$out_dir"

    log_info "generating REST API code for $service_name to $out_dir"

    # Generate types
    run oapi-codegen \
      -package "$service_name" \
      -generate types \
      -o "$out_dir/types.gen.go" \
      "$openapi_file"

    # Generate client (optional)
    run oapi-codegen \
      -package "$service_name" \
      -generate client \
      -o "$out_dir/client.gen.go" \
      "$openapi_file"

    # Generate server interface
    run oapi-codegen \
      -package "$service_name" \
      -generate gin-server \
      -o "$out_dir/server.gen.go" \
      "$openapi_file"

  done
  log_success "OpenAPI code generation done!"
}

echo "found openapi_files: $(openapi_files)"
gen_for_openapi