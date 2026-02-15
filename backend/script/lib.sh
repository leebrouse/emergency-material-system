#!/usr/bin/env bash

function log_info() {
    echo -e "\033[0;34m[INFO]\033[0m $1"
}

function log_warning() {
    echo -e "\033[0;33m[WARN]\033[0m $1"
}

function log_success() {
    echo -e "\033[0;32m[SUCCESS]\033[0m $1"
}

function run() {
    log_info "executing: $*"
    "$@"
}
