#!/bin/bash -l

#Enable bash history
HISTFILE=/root/.bash_history
set -o history

# Generate creds for the workshop-app role using the Vault CLI
{ read LEASE_ID; } < <(vault read -format=json lob_a/workshop/database/creds/workshop-app | jq -r '.lease_id')

# Extend the lease of the credentials by 2 minutes
vault write sys/leases/renew lease_id="${LEASE_ID}" increment="120"

# Lookup the lease of the credentials
vault write sys/leases/lookup lease_id="${LEASE_ID}"

# Revoke the credentials
vault write sys/leases/revoke lease_id="${LEASE_ID}"

exit 0
