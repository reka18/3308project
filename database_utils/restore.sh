#!/bin/bash

# Exit immediately on command failure
set -e

# Change directory to this scripts dir (this script runs in a subshell, so it won't affect the parent shell)
cd "${0%/*}"

# Import some utility functions
source "functions/functions.sh"

python functions/replace_db_user.py ${USER}

# Positional parameters for the SQL scripts with safe defaults
# -- target database name and user
db_name=${1:-socialmediasite}
db_user=${2:-${USER}}
# -- SQL script paths to refresh and restore the database
data_script=${3:-"sql/restore_database.sql"}

# Log file paths
log_dir="logs/$(date +%Y-%m-%d)" && mkdir -p ${log_dir}
log_file="${log_dir}/clean_database_$(date +%s).log"

# Start of script execution
log "Using settings:"
log "\tdb_name: ${db_name}"
log "\tdb_user: ${db_user}"
log "\tdata_script: ${data_script}"

echo >> ${log_file} # Append a new line to help visually separate SQL logs
format_log_header ${data_script} ${log_file}

# Restore data and schema to the database
psql \
    --no-psqlrc \
    --dbname ${db_name} \
    --username ${db_user} \
    --file ${data_script} \
    --echo-queries \
    --single-transaction \
    --set ON_ERROR_STOP=on >> ${log_file}

react_to_exit_code $? "Failed while trying to restore the database"

log
log "#####################################################################"
log "Successfully restored database using ${data_script}"
log "#####################################################################"

log
handle_exit 0 "Saved log to ${log_file}"
