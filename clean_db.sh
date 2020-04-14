#!/bin/bash
# title         : clean_db.sh
# description   : Clean your current database by dropping and recreating it, and restoring it to the needed state for
#                 integration tests.
# author        : Caidan Williams, Reagan Karnes, Jeff Jones
# date          : 2019-06-25
# version       : 2.0
# usage         : ./clean_db.sh
# notes         : If you `source` this script, your shell's directory will change.
# bash_version  : 5.0.3(1)-release
# ======================================================================================================================

# Exit immediately on command failure
set -e

# Change directory to this scripts dir (this script runs in a sub-shell, so it won't affect the parent shell)
cd "${0%/*}"

# Import some utility functions
source "functions/general.sh"

# Positional parameters for the SQL scripts with safe defaults
# -- target database name and user
db_name="gitprime"
db_user=${USER}
# -- SQL script paths to refresh and restore the database
data_script="${1:-"sql/restore_database.sql"}" # this is most likely to change
drop_script="sql/refresh_database_template.sql"

git lfs pull || echo "Error pulling LFS resources. Have you run `git lfs install` in this directory?"

function restartRedis()
{
    log
    log "#####################################################################"
    log "Restarting reddisson service."
    log "#####################################################################"
    log

    success=0
    redis-cli -p 6999 flushall && success=1
    if [[ success -eq 1 ]]; then
        log
        log "#####################################################################"
        log "Successfully restarted reddisson service."
        log "#####################################################################"
    else
        log
        log "#####################################################################"
        log "Unable to automatically restart reddison service."
        log "Manually restart the services before running test."
        log "#####################################################################"
    fi

}

function setup()
{
#    echo
#    unzip test_dbs.zip || return
#    echo
#
#    log
#    log "#####################################################################"
#    log "Successfully unzipped 'test_dbs.zip'"
#    log "#####################################################################"
#    log
#
#    if test -f ${data_script}; then
#        log "#####################################################################"
#        log "Found ${data_script}. Setting up test database."
#        log "#####################################################################"
#        log
#    else
#        log "#####################################################################"
#        log "Unable to find '${data_script}'. ABORTING!"
#        log "#####################################################################"
#        return
#    fi
#
#
#    # Log file paths
#    log_dir="logs/$(date +%Y-%m-%d)" && mkdir -p ${log_dir}
#    log_file="${log_dir}/clean_database_$(date +%s).log"
#
#    # Start of script execution
#
#    log "Using settings:"
#    log "\tdb_name: ${db_name}"
#    log "\tdb_user: ${db_user}"
#    log "\tdrop_script: ${drop_script}"
#    log "\tdata_script: ${data_script}"
#
#    format_log_header ${drop_script} ${log_file}
#
#    # Drop and recreate the current database
#    psql \
#        --username ${db_user} \
#        --file ${drop_script} \
#        --echo-queries \
#        --set db_name=${db_name} \
#        --set qdb_name=\'${db_name}\' \
#        --set db_user=${db_user} \
#        postgres >> ${log_file}
#
#    react_to_exit_code $? "Failed while trying to clean the database"
#
#    log
#    log "#####################################################################"
#    log "Successfully cleaned database using ${drop_script}"
#    log "#####################################################################"
#
#    echo >> ${log_file} # Append a new line to help visually separate SQL logs
#    format_log_header ${data_script} ${log_file}
#
#    # Restore data and schema to the database
#    psql \
#        --no-psqlrc \
#        --dbname ${db_name} \
#        --username ${db_user} \
#        --file ${data_script} \
#        --echo-queries \
#        --single-transaction \
#        --set ON_ERROR_STOP=on >> ${log_file}
#
#    react_to_exit_code $? "Failed while trying to restore the database"
#
#    log
#    log "#####################################################################"
#    log "Successfully restored database using ${data_script}"
#    log "#####################################################################"

    restartRedis
}

function failure()
{
    log
    log "#####################################################################"
    log "Unable to setup test database due to errors."
    log "#####################################################################"
}

function sudo_setup()
{
    log
    log "#####################################################################"
    log "'sudo' permissions for reddisson have not been set."
    log "Running 'redis_permissions.sh'."
    log "#####################################################################"
    ./redis_permissions.sh
    log
    log "#####################################################################"
    log "Rerun './clean_db.sh'"
    log "#####################################################################"
    exit 0
}

echo "########## !!!!! WARNING !!!!! ##########"
echo "This action will wipe your Flow database."
echo "Are you sure you want to proceed? [y]es "
read -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    log
    log "Execution confirmed. Wiping database."
    log
    setup || failure
    rm -rf sql/
else
    log "Execution aborted. Database untouched."
fi

log
handle_exit 0 "Saved log to ${log_file}"

