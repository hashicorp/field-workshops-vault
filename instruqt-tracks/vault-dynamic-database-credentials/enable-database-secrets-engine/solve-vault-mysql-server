#!/bin/bash -l

#Enable bash history
HISTFILE=/root/.bash_history
set -o history

# List Vault Secrets
vault secrets list

# Enable the Database secrets engine
vault secrets enable -path=lob_a/workshop/database database

exit 0
