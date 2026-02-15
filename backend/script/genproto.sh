#!/usr/bin/env bash

set -euo pipefail

source ./backend/script/lib.sh

PROTO_ROOT="./backend/api/proto"
PROTO_OUT="./backend/internal/common/genproto"

function gen_for_proto() {
  run mkdir -p "$PROTO_OUT"

  for proto_file in $(find "$PROTO_ROOT" -name "*.proto"); do
    local service_name=$(basename "$proto_file" .proto)
    local out_dir="$PROTO_OUT/$service_name"
    
    run mkdir -p "$out_dir"
    
    log_info "generating gRPC code for $service_name to $out_dir"
    
    run protoc \
      --proto_path="$PROTO_ROOT" \
      --go_out="$out_dir" --go_opt=paths=source_relative \
      --go-grpc_out="$out_dir" --go-grpc_opt=paths=source_relative \
      "$proto_file"
  done
  
  log_success "gRPC code generation done!"
}

gen_for_proto
