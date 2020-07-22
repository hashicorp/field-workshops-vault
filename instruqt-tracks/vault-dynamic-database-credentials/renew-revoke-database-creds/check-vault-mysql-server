#!/bin/bash -l

set -e

grep -q "vault.*read.*lob_a/workshop/database/creds/workshop-app" /root/.bash_history || fail-message "You haven't generated creds for the workshop-app role in this challenge using the Vault CLI yet."

grep -q "vault.*write.*sys/leases/renew.*lease_id=.*increment=" /root/.bash_history || fail-message "You haven't renewed the lease for your credentials yet."

grep -q "vault.*write.*sys/leases/lookup.*lease_id=.*" /root/.bash_history || fail-message "You haven't looked up the lease for your credentials yet."

grep -q "vault.*write.*sys/leases/revoke.*lease_id=.*" /root/.bash_history || fail-message "You haven't revoked the lease for your credentials yet."

exit 0
