#!/bin/bash

# Print to console with datetime
function Info() {
    DATE_STR=$(date +"%Y-%m-%dT%H:%M:%S %z")

    echo -e "[${DATE_STR}] $*"
}

# Put headers in the log file to help differentiate each sql script output
function format_log_header() {
    local script_name=${1}
    local underline=$(printf "=%.0s" $(eval echo "{1..${#script_name}}"))
    local log_file=${2}

    echo -e "${script_name}\n${underline}\n" >> ${log_file}
}

# Exit script if exit code is not zero
function react_to_exit_code() {
    exit_code=$1

    shift 1

    log_message="$*"

    if [[ ${exit_code} != 0 ]];
    then
        handle_exit 1000 "$log_message"
    fi
}

# Log exit reason and exit script with exit code
function handle_exit() {
    EXIT_CODE=$1

    shift

    if [[ ! -z "$@" ]];
    then
        log "Exiting: $*"
    fi

    exit "${EXIT_CODE}"
}

function cleanup_on_pip_fail() {
    rm -rf $1
    log "Pip install failure. Deleting clone."
    exit 1
}