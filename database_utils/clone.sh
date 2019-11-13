#!/bin/bash

# Exit immediately on command failure
set -e

# Change directory to this scripts dir (this script runs in a subshell, so it won't affect the parent shell)
cd "${0%/*}"

# Import some utility functions
source "functions/functions.sh"

# Store date and time here to allow matching log and dump file names
today_date=$(date +%Y-%m-%d)
epoch_time=$(date +%s)

# Positional parameters for the SQL sql with safe defaults
db_name=${1:-socialmediasite}
db_user=${2:-${USER}}
file_name=${3:-"restore_database.sql"}

# SQL sql to refresh and restore the database
dump_dir="dumps" && mkdir -p ${dump_dir}
dump_file="${dump_dir}/${file_name}"

# Log file locations
log_dir="logs/${today_date}" && mkdir -p ${log_dir}
log_file="${log_dir}/clone_database_${epoch_time}.log"

# Start of script execution

log "Using settings:"
log "\tdb_name: ${db_name}"
log "\tdb_user: ${db_user}"
log "\tfile_name: ${file_name}"

# Clone the data and schema of the database
pg_dump \
    --username ${db_user} \
    --dbname ${db_name} \
    --file ${dump_file} \
    --verbose 1>&2 2> ${log_file}

react_to_exit_code $? "Failed while trying to clone the database"

log
log "########################################################################"
log "Successfully cloned database to ${dump_file}"
log "########################################################################"

log
handle_exit 0 "Saved log to ${log_file}"
