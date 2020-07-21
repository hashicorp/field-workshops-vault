#!/bin/bash -l

set -e

grep -q "vault.*secrets.*list" /root/.bash_history || fail-message "You haven't listed the Vault secrets engines yet."

grep -q "vault.*secrets.*enable.*-path=lob_a/workshop/database database" /root/.bash_history || fail-message "You haven't enabled the Database secrets engine yet on the path lob_a/workshop/database."

exit 0
