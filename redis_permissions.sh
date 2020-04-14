#!/bin/bash
# title         : clean_db.sh
# description   : Clean your current database by dropping and recreating it, and restoring it to the needed state for
#                 integration tests.
# author        : Reagan Karnes
# date          : 2020-1-9
# version       : 1.0
# usage         : ./redis_permissions.sh
# ======================================================================================================================

# Exit immediately on command failure
set -e

# Change directory to this scripts dir (this script runs in a sub-shell, so it won't affect the parent shell)
cd "${0%/*}"

# Import some utility functions
source "functions/general.sh"

function write()
{
    echo "%${USER} ALL=NOPASSWD: $(which systemctl) restart redis" > gp_redis
    log
    log "Password may be needed to write to protected directory"
    sudo chown root:root gp_redis
    echo
    sudo mv gp_redis /etc/sudoers.d

    log
    log "#####################################################################"
    log "Writing:"
    log "  '%${USER} ALL=NOPASSWD: $(which systemctl) restart redis'"
    log "to the path: '/etc/sudoers.d/gp_redis'"
    log "To reverse this, run 'sudo rm /etc/sudoers.d/gp_redis'"
    log "#####################################################################"
}

THIS_OS=$(uname -s)
case ${THIS_OS} in
    Darwin*)
        echo "#####################################################################"
        echo "This is not needed for macOS. Aborting."
        echo "#####################################################################"
        exit 0
        ;;
esac

log
log "#####################################################################"
log "This script allows the following specific commands to be run without"
log "requiring you to enter your password. These permissions are limited to"
log "these specific commands and do not reduce your system security or store"
log "your system password for other commands."
log "Commands:"
log "  'sudo systemctl restart redis'"
log "#####################################################################"
log
log "You only need to run this script once."
log "Are you sure you want to proceed? [y]es "
read -n 1 -r

if [[ $REPLY =~ ^[Yy]$ ]]; then
    log
    log "Execution confirmed. Writing to '/etc/sudoers.sh/gp_redis'"
    log
    write || log "Unable to write to directory."
else
    log "Execution aborted."
fi
