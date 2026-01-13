#!/usr/bin/env bash

set -euo pipefail

shopt -s globstar

if ! [[ "$0" =~ script/genproto.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

source ./backend/script/lib.sh

API_ROOT="./backend/api"

function dirs {
  dirs=()
  while IFS= read -r dir; do
    dirs+=("$dir")
  done < <(find ./backend/api/proto -type f -name "*.proto" -exec dirname {} \; | sort -u)
  echo "${dirs[@]}"
}

function pb_files {
  pb_files=$(find . -type f -name '*.proto')
  echo "${pb_files[@]}"
}

function gen_for_modules() {
  local go_out="./backend/internal/common/genproto"
  if [ -d "$go_out" ]; then
    log_warning "found existing $go_out, cleaning all files under it"
    run rm -rf $go_out
  fi

  run mkdir -p "$go_out"

  for pb_file in $(pb_files); do
    local service_name=$(basename "$pb_file" .proto)
    local dir_name="${service_name}"
    local out_dir="$go_out/$dir_name"

    if [ -d "$out_dir" ]; then
        log_warning "found existing $out_dir, cleaning all files under it"
        run rm -rf "$out_dir"/*
    else
      run mkdir -p "$out_dir"
    fi
    log_info "generating code for $service_name to $out_dir"

    run protoc \
      -I="/usr/local/include/" \
      -I="${API_ROOT}" \
      "--go_out=${go_out}" --go_opt=module=github.com/emergency-material-system/backend/internal/common/genproto \
      --go-grpc_opt=require_unimplemented_servers=false \
      "--go-grpc_out=${go_out}" --go-grpc_opt=module=github.com/emergency-material-system/backend/internal/common/genproto \
      "$pb_file"
  done
  log_success "protoc gen done!"
}

echo "directories containing protos to be built: $(dirs)"
echo "found pb_files: $(pb_files)"
gen_for_modules