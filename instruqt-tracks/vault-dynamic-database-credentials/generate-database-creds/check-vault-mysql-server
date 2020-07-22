#!/bin/bash -l

set -e

grep -q 'curl.*--header.*"X-Vault-Token:.*root".*"http://localhost:8200/v1/lob_a/workshop/database/creds/workshop-app-long".*|.*jq' /root/.bash_history || fail-message "You haven't generated creds for the workshop-app-long role using the Vault API yet."

grep -q "vault.*read.*lob_a/workshop/database/creds/workshop-app-long" /root/.bash_history || fail-message "You haven't generated creds for the workshop-app-long role using the Vault CLI yet."

grep -q "vault.*read.*lob_a/workshop/database/creds/workshop-app" /root/.bash_history || fail-message "You haven't generated creds for the workshop-app role using the Vault CLI yet."

grep -q "mysql.*-u.*-p" /root/.bash_history || fail-message "You haven't connected to the MySQL server with the 'mysql' command yet."

exit 0
